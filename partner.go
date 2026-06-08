package shopee

import "strconv"

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
		ShopList  []PartnerShopInfo `json:"shop_list"`
		Total     int               `json:"total"`
		PageNum   int               `json:"page_num"`
		PageSize  int               `json:"page_size"`
	} `json:"response"`
}

func (s *PartnerService) GetShopsByPartner(pageSize, pageNumber int) (*GetShopsByPartnerResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetShopsByPartnerResponse{}
	if err := s.client.DoGet(PathPartnerGetShopsByPartner, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetMerchantsByPartner() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(PathPartnerGetMerchantsByPartner, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetAccessTokenParams struct {
	Code      string `json:"code"`
	PartnerID int64  `json:"partner_id"`
	ShopID    int64  `json:"shop_id,omitempty"`
}

type GetAccessTokenResponse struct {
	BaseResponse
	Response *struct {
		AccessToken  string `json:"access_token"`
		ExpireIn     int64  `json:"expire_in"`
		ShopID       int64  `json:"shop_id"`
		RefreshToken string `json:"refresh_token,omitempty"`
	} `json:"response,omitempty"`
}

func (s *PartnerService) GetAccessToken(params *GetAccessTokenParams) (*GetAccessTokenResponse, error) {
	result := &GetAccessTokenResponse{}
	if err := s.client.DoPost(PathPartnerGetAccessToken, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetTokenByResendCode(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathPartnerGetTokenByResendCode, params, result); err != nil {
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
	if err := s.client.DoGet(PathPartnerGetShopeeIPRanges, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) SetAppPushConfig(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathPartnerSetAppPushConfig, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetAppPushConfig() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(PathPartnerGetAppPushConfig, map[string]string{}, result); err != nil {
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
	if err := s.client.DoGet(PathPartnerGetLostPushMessage, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PartnerService) GetBoundWHSInfo() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(PathPartnerGetBoundWHSInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
