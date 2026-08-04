package shopee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

// HTTPClient is the interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var defaultClient HTTPClient = &http.Client{Timeout: 120 * time.Second}

// Client handles HTTP communication with the Shopee Open Platform.
type Client struct {
	PartnerID   int64
	PartnerKey  string
	AccessToken string
	ShopID      int64
	MerchantID  int64
	BaseURL     string
	Region      Region
	HTTPClient  HTTPClient
	mu          sync.RWMutex
}

// ClientOption configures a Client.
type ClientOption func(*Client)

// WithRegion sets the region for the API base URL.
func WithRegion(region Region) ClientOption {
	return func(c *Client) {
		c.Region = region
		if u, ok := BaseURLs[region]; ok {
			c.BaseURL = u
		}
	}
}

// WithBaseURL overrides the base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = strings.TrimRight(baseURL, "/")
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc HTTPClient) ClientOption {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

// WithMerchantID sets the merchant ID for Merchant API calls.
func WithMerchantID(merchantID int64) ClientOption {
	return func(c *Client) {
		c.MerchantID = merchantID
	}
}

// NewClient creates a new Shopee API client.
func NewClient(partnerID int64, partnerKey, accessToken string, shopID int64, opts ...ClientOption) *Client {
	c := &Client{
		PartnerID:   partnerID,
		PartnerKey:  partnerKey,
		AccessToken: accessToken,
		ShopID:      shopID,
		BaseURL:     BaseURLs[RegionGlobal],
		Region:      RegionGlobal,
		HTTPClient:  defaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// SetAccessToken updates the access token in a thread-safe manner.
func (c *Client) SetAccessToken(token string) {
	c.mu.Lock()
	c.AccessToken = token
	c.mu.Unlock()
}

// GetAccessToken returns the current access token in a thread-safe manner.
func (c *Client) GetAccessToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.AccessToken
}

// authSnapshot returns a consistent snapshot of auth fields for signing.
func (c *Client) authSnapshot() (accessToken string, shopID, merchantID int64) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.AccessToken, c.ShopID, c.MerchantID
}

// isMerchantAPI reports whether the path belongs to a Merchant-domain API
// (merchant / global_product / first_mile). These are only used by
// cross-border sellers and sign with merchant_id instead of shop_id.
func IsMerchantAPI(apiPath string) bool {
	return strings.HasPrefix(apiPath, "/api/v2/merchant/") ||
		strings.HasPrefix(apiPath, "/api/v2/global_product/") ||
		strings.HasPrefix(apiPath, "/api/v2/first_mile/")
}

// baseQuery builds the common query parameters. Merchant-domain APIs carry
// merchant_id; all other seller APIs carry shop_id. Never both (Shopee
// rejects a request whose query mixes the two: "Wrong sign").
func (c *Client) baseQuery(apiPath string, timestamp int64) url.Values {
	accessToken, shopID, merchantID := c.authSnapshot()
	q := url.Values{}
	q.Set("partner_id", strconv.FormatInt(c.PartnerID, 10))
	q.Set("timestamp", strconv.FormatInt(timestamp, 10))
	if accessToken != "" {
		q.Set("access_token", accessToken)
	}
	if IsMerchantAPI(apiPath) {
		if merchantID > 0 {
			q.Set("merchant_id", strconv.FormatInt(merchantID, 10))
		}
	} else if shopID > 0 {
		q.Set("shop_id", strconv.FormatInt(shopID, 10))
	}
	return q
}

func (c *Client) generateSign(apiPath string, timestamp int64) string {
	accessToken, shopID, merchantID := c.authSnapshot()
	if IsMerchantAPI(apiPath) {
		// Merchant API: sign = partner_id + path + timestamp + access_token + merchant_id
		return GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, timestamp, accessToken, 0, merchantID)
	}
	return GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, timestamp, accessToken, shopID, merchantID)
}

// DoGet performs a GET request with context support.
func (c *Client) DoGet(ctx context.Context, apiPath string, queryParams map[string]string, result any) error {
	ts := time.Now().Unix()
	q := c.baseQuery(apiPath, ts)
	for k, v := range queryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	sign := c.generateSign(apiPath, ts)
	q.Set("sign", sign)
	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	return c.doRequest(req, result)
}

// DoGetMulti is like DoGet but accepts url.Values, allowing multiple values
// for the same query parameter key (e.g. item_status=NORMAL&item_status=UNLIST).
func (c *Client) DoGetMulti(ctx context.Context, apiPath string, queryParams url.Values, result any) error {
	ts := time.Now().Unix()
	q := c.baseQuery(apiPath, ts)
	for k, vs := range queryParams {
		for _, v := range vs {
			if v != "" {
				q.Add(k, v)
			}
		}
	}
	sign := c.generateSign(apiPath, ts)
	q.Set("sign", sign)
	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	return c.doRequest(req, result)
}

// DoPost performs a POST request with context support.
func (c *Client) DoPost(ctx context.Context, apiPath string, bodyPayload any, result any) error {
	ts := time.Now().Unix()

	var bodyBytes []byte
	if bodyPayload != nil {
		var err error
		bodyBytes, err = json.Marshal(bodyPayload)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
	}

	q := c.baseQuery(apiPath, ts)
	sign := c.generateSign(apiPath, ts)
	q.Set("sign", sign)

	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())
	bodyReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	if bodyPayload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return c.doRequest(req, result)
}

// DoPostMulti sends a multipart/form-data POST request (used by Shopee media
// upload endpoints like media/upload_image, which reject JSON bodies with
// "Content-Type must be multipart/form-data"). fields holds form fields;
// files maps field name -> list of file byte slices (multiple files per field).
// sniffImageMeta returns a proper filename and Content-Type for the uploaded
// bytes by inspecting magic bytes. Shopee rejects file parts whose filename
// lacks a JPG/JPEG/PNG extension or whose part Content-Type is not an image
// type ("Some images are not in JPG, JPEG, PNG format").
func sniffImageMeta(data []byte) (string, string) {
	switch {
	case len(data) >= 8 && string(data[:8]) == "\x89PNG\r\n\x1a\n":
		return "image.png", "image/png"
	case len(data) >= 3 && data[0] == 0xFF && data[1] == 0xD8 && data[2] == 0xFF:
		return "image.jpg", "image/jpeg"
	case len(data) >= 6 && (string(data[:6]) == "GIF87a" || string(data[:6]) == "GIF89a"):
		return "image.gif", "image/gif"
	default:
		return "image.bin", "application/octet-stream"
	}
}

func (c *Client) DoPostMulti(ctx context.Context, apiPath string, fields map[string]string, files map[string][][]byte, result any) error {
	ts := time.Now().Unix()
	q := c.baseQuery(apiPath, ts)
	sign := c.generateSign(apiPath, ts)
	q.Set("sign", sign)

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)
	for k, v := range fields {
		if err := mw.WriteField(k, v); err != nil {
			return fmt.Errorf("write multipart field %s: %w", k, err)
		}
	}
	for k, dataList := range files {
		for _, data := range dataList {
			fname, ctype := sniffImageMeta(data)
			h := textproto.MIMEHeader{}
			h.Set("Content-Disposition",
				fmt.Sprintf(`form-data; name="%s"; filename="%s"`, k, fname))
			h.Set("Content-Type", ctype)
			fw, err := mw.CreatePart(h)
			if err != nil {
				return fmt.Errorf("create form file %s: %w", k, err)
			}
			if _, err := fw.Write(data); err != nil {
				return fmt.Errorf("write form file %s: %w", k, err)
			}
		}
	}
	if err := mw.Close(); err != nil {
		return fmt.Errorf("close multipart: %w", err)
	}

	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, &buf)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	return c.doRequest(req, result)
}

// doRequest executes the HTTP request and unmarshals the response.
func (c *Client) doRequest(req *http.Request, result any) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("DEBUG: Shopee non-2xx response: status=%d url=%s body=%s",
			resp.StatusCode, req.URL.String(), truncate(string(body), 2000))
		var apiErr struct {
			Error     string `json:"error"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		}
		if json.Unmarshal(body, &apiErr) == nil && apiErr.Error != "" {
			return &APIError{ErrorCode: apiErr.Error, Message: apiErr.Message, RequestID: apiErr.RequestID}
		}
		return fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, resp.Status, truncate(string(body), 500))
	}

	// Debug: check for error in success response
	var errCheck struct {
		Error     string `json:"error"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	}
	if json.Unmarshal(body, &errCheck) == nil && errCheck.Error != "" {
		log.Printf("DEBUG: Shopee returned error in 2xx body: url=%s error=%q message=%q request_id=%q body=%s",
			req.URL.String(), errCheck.Error, errCheck.Message, errCheck.RequestID, truncate(string(body), 1000))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshal: %w (body: %s)", err, truncate(string(body), 500))
	}
	return nil
}

// DoGetWithTimestamp performs a GET request with a custom timestamp (for testing).
func (c *Client) DoGetWithTimestamp(ctx context.Context, apiPath string, queryParams map[string]string, timestamp int64, result any) error {
	q := c.baseQuery(apiPath, timestamp)
	for k, v := range queryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	sign := c.generateSign(apiPath, timestamp)
	q.Set("sign", sign)
	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	return c.doRequest(req, result)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
