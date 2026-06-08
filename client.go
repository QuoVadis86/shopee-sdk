package shopee

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// HTTPClient is the interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var defaultClient HTTPClient = &http.Client{Timeout: 30 * time.Second}

// Client handles HTTP communication with the Shopee Open Platform.
type Client struct {
	PartnerID   int64
	PartnerKey  string
	AccessToken string
	ShopID      int64
	BaseURL     string
	HTTPClient  HTTPClient
}

// ClientOption configures a Client.
type ClientOption func(*Client)

// WithRegion sets the region for the API base URL.
func WithRegion(region Region) ClientOption {
	return func(c *Client) {
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

// NewClient creates a new Shopee API client.
func NewClient(partnerID int64, partnerKey, accessToken string, shopID int64, opts ...ClientOption) *Client {
	c := &Client{
		PartnerID:   partnerID,
		PartnerKey:  partnerKey,
		AccessToken: accessToken,
		ShopID:      shopID,
		BaseURL:     BaseURLs[RegionGlobal],
		HTTPClient:  defaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) baseQuery(timestamp int64) url.Values {
	q := url.Values{}
	q.Set("partner_id", strconv.FormatInt(c.PartnerID, 10))
	q.Set("timestamp", strconv.FormatInt(timestamp, 10))
	q.Set("access_token", c.AccessToken)
	if c.ShopID > 0 {
		q.Set("shop_id", strconv.FormatInt(c.ShopID, 10))
	}
	return q
}

func (c *Client) signAndBuildURL(apiPath string, timestamp int64) string {
	q := c.baseQuery(timestamp)
	sign := GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, timestamp, c.AccessToken, c.ShopID)
	q.Set("sign", sign)
	return fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())
}

// DoGet performs a GET request.
// queryParams are additional query parameters beyond the standard auth params.
func (c *Client) DoGet(apiPath string, queryParams map[string]string, result any) error {
	ts := time.Now().Unix()
	q := c.baseQuery(ts)
	for k, v := range queryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	sign := GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, ts, c.AccessToken, c.ShopID)
	q.Set("sign", sign)
	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshal: %w (body: %s)", err, truncate(string(body), 500))
	}
	return nil
}

// DoPost performs a POST request.
// bodyPayload is the JSON-serializable request body.
func (c *Client) DoPost(apiPath string, bodyPayload any, result any) error {
	ts := time.Now().Unix()

	var bodyBytes []byte
	if bodyPayload != nil {
		var err error
		bodyBytes, err = json.Marshal(bodyPayload)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
	}

	q := c.baseQuery(ts)
	sign := GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, ts, c.AccessToken, c.ShopID)
	q.Set("sign", sign)

	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())
	bodyReader := strings.NewReader(string(bodyBytes))

	req, err := http.NewRequest(http.MethodPost, reqURL, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if err := json.Unmarshal(respBody, result); err != nil {
		return fmt.Errorf("unmarshal: %w (body: %s)", err, truncate(string(respBody), 500))
	}
	return nil
}

// DoGetWithTimestamp performs a GET request with a custom timestamp (for testing).
func (c *Client) DoGetWithTimestamp(apiPath string, queryParams map[string]string, timestamp int64, result any) error {
	q := c.baseQuery(timestamp)
	for k, v := range queryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	sign := GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, timestamp, c.AccessToken, c.ShopID)
	q.Set("sign", sign)
	reqURL := fmt.Sprintf("%s%s?%s", c.BaseURL, apiPath, q.Encode())

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshal: %w (body: %s)", err, truncate(string(body), 500))
	}
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
