package shopee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

var defaultClient HTTPClient = &http.Client{Timeout: 30 * time.Second}

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

func (c *Client) baseQuery(timestamp int64) url.Values {
	accessToken, shopID, merchantID := c.authSnapshot()
	q := url.Values{}
	q.Set("partner_id", strconv.FormatInt(c.PartnerID, 10))
	q.Set("timestamp", strconv.FormatInt(timestamp, 10))
	if accessToken != "" {
		q.Set("access_token", accessToken)
	}
	if shopID > 0 {
		q.Set("shop_id", strconv.FormatInt(shopID, 10))
	}
	if merchantID > 0 {
		q.Set("merchant_id", strconv.FormatInt(merchantID, 10))
	}
	return q
}

func (c *Client) generateSign(apiPath string, timestamp int64) string {
	accessToken, shopID, merchantID := c.authSnapshot()
	return GenerateSignature(c.PartnerKey, c.PartnerID, apiPath, timestamp, accessToken, shopID, merchantID)
}

// DoGet performs a GET request with context support.
func (c *Client) DoGet(ctx context.Context, apiPath string, queryParams map[string]string, result any) error {
	ts := time.Now().Unix()
	q := c.baseQuery(ts)
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

	q := c.baseQuery(ts)
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
		return fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, resp.Status, truncate(string(body), 500))
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("unmarshal: %w (body: %s)", err, truncate(string(body), 500))
	}
	return nil
}

// DoGetWithTimestamp performs a GET request with a custom timestamp (for testing).
func (c *Client) DoGetWithTimestamp(ctx context.Context, apiPath string, queryParams map[string]string, timestamp int64, result any) error {
	q := c.baseQuery(timestamp)
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
