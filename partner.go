package shopee

import (
	"context"
	"net/url"
	"strconv"
	"time"
)

type PartnerService struct {
	client *Client
}

func NewPartnerService(client *Client) *PartnerService {
	return &PartnerService{client: client}
}

type PartnerShopInfo struct {
	ShopID   int64  `json:"shop_id"`
	ShopName string `json:"shop_name"`
	Region   string `json:"region"`
	Status   string `json:"status"`
}

type GetShopsByPartnerResponse struct {
	BaseResponse
	Response struct {
		ShopList []PartnerShopInfo `json:"shop_list"`
		Total    int               `json:"total"`
		PageNum  int               `json:"page_num"`
		PageSize int               `json:"page_size"`
	} `json:"response"`
}

func (s *PartnerService) GetShopsByPartner(pageSize, pageNumber int) (*GetShopsByPartnerResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetShopsByPartnerResponse{}
	if err := s.client.DoGet(context.Background(), PathPartnerGetShopsByPartner, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetMerchantsByPartner() (*GetMerchantsByPartnerResponse, error) {
	result := &GetMerchantsByPartnerResponse{}
	if err := s.client.DoGet(context.Background(), PathPartnerGetMerchantsByPartner, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

// PartnerMerchantInfo represents a merchant returned by GetMerchantsByPartner.
type PartnerMerchantInfo struct {
	MerchantID int64  `json:"merchant_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}

// GetMerchantsByPartnerResponse is the response for GetMerchantsByPartner.
type GetMerchantsByPartnerResponse struct {
	BaseResponse
	Response struct {
		MerchantList []PartnerMerchantInfo `json:"merchant_list"`
	} `json:"response"`
}

type GetAccessTokenParams struct {
	Code      string `json:"code"`
	PartnerID int64  `json:"partner_id"`
	ShopID    int64  `json:"shop_id,omitempty"`
}

type GetAccessTokenResponse struct {
	BaseResponse
	AccessToken    string  `json:"access_token"`
	RefreshToken   string  `json:"refresh_token"`
	ExpireIn       int64   `json:"expire_in"`
	MerchantIDList []int64 `json:"merchant_id_list"`
	ShopIDList     []int64 `json:"shop_id_list"`
	UserIDList     []int64 `json:"user_id_list"`
}

func (s *PartnerService) GetAccessToken(params *GetAccessTokenParams) (*GetAccessTokenResponse, error) {
	result := &GetAccessTokenResponse{}
	if err := s.client.DoPost(context.Background(), PathPartnerGetAccessToken, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type RefreshAccessTokenParams struct {
	RefreshToken string `json:"refresh_token"`
	PartnerID    int64  `json:"partner_id"`
	ShopID       int64  `json:"shop_id,omitempty"`
	MerchantID   int64  `json:"merchant_id,omitempty"`
}

type RefreshAccessTokenResponse struct {
	BaseResponse
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int64  `json:"expire_in"`
	ShopID       int64  `json:"shop_id"`
	MerchantID   int64  `json:"merchant_id"`
}

type refreshAccessTokenEnvelope struct {
	BaseResponse
	AccessToken  string                      `json:"access_token"`
	RefreshToken string                      `json:"refresh_token"`
	ExpireIn     int64                       `json:"expire_in"`
	ShopID       int64                       `json:"shop_id"`
	MerchantID   int64                       `json:"merchant_id"`
	Response     *RefreshAccessTokenResponse `json:"response"`
}

// RefreshAccessToken exchanges a refresh token for a new access token.
// The auth endpoint uses a public signature, so the client's current access
// token and shop ID are intentionally excluded from the request query.
func (s *PartnerService) RefreshAccessToken(params *RefreshAccessTokenParams) (*RefreshAccessTokenResponse, error) {
	// Create a dedicated public-signature client to avoid shallow copy issues
	// with pointer fields on Client.
	publicClient := &Client{
		PartnerID:  s.client.PartnerID,
		PartnerKey: s.client.PartnerKey,
		// AccessToken and ShopID intentionally empty for public signature
		MerchantID: 0,
		BaseURL:    s.client.BaseURL,
		Region:     s.client.Region,
		HTTPClient: s.client.HTTPClient,
	}

	result := &refreshAccessTokenEnvelope{}
	if err := publicClient.DoPost(context.Background(), PathPartnerRefreshAccessToken, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	if result.Response != nil {
		return result.Response, nil
	}
	return &RefreshAccessTokenResponse{
		BaseResponse: result.BaseResponse,
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		ExpireIn:     result.ExpireIn,
		ShopID:       result.ShopID,
		MerchantID:   result.MerchantID,
	}, nil
}

func (s *PartnerService) GetTokenByResendCode(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPartnerGetTokenByResendCode, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type IPRangeInfo struct {
	IPRange string `json:"ip_range"`
}

type GetShopeeIPRangesResponse struct {
	BaseResponse
	Response struct {
		IPRangeList []IPRangeInfo `json:"ip_range_list"`
	} `json:"response"`
}

func (s *PartnerService) GetShopeeIPRanges() (*GetShopeeIPRangesResponse, error) {
	result := &GetShopeeIPRangesResponse{}
	if err := s.client.DoGet(context.Background(), PathPartnerGetShopeeIPRanges, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) SetAppPushConfig(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPartnerSetAppPushConfig, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetAppPushConfig() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPartnerGetAppPushConfig, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetLostPushMessage(messageID string) (*BaseResponse, error) {
	q := map[string]string{"message_id": messageID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPartnerGetLostPushMessage, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetBoundWHSInfo() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPartnerGetBoundWHSInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

// BuildAuthURL constructs the OAuth authorization URL that sellers visit to
// authorize the partner application. The redirect parameter is the callback URL
// where Shopee will send the authorization code after the seller approves.
func (s *PartnerService) BuildAuthURL(redirect string) (string, error) {
	ts := time.Now().Unix()
	authBase, ok := AuthURLs[s.client.Region]
	if !ok {
		authBase = AuthURLs[RegionGlobal]
	}
	apiPath := "/api/v2/shop/auth_partner"
	sign := GenerateSignature(s.client.PartnerKey, s.client.PartnerID, apiPath, ts, "", 0, 0)

	u, err := url.Parse(authBase)
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("partner_id", strconv.FormatInt(s.client.PartnerID, 10))
	q.Set("timestamp", strconv.FormatInt(ts, 10))
	q.Set("redirect", redirect)
	q.Set("sign", sign)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
