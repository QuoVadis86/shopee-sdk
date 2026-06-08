package shopee

import "strconv"

type BRService struct {
	client *Client
}

func NewBRService(client *Client) *BRService {
	return &BRService{client: client}
}

type BRShopEnrollmentStatus struct {
	ShopID           int64  `json:"shop_id"`
	IsEnrolled       bool   `json:"is_enrolled"`
	EnrollmentStatus string `json:"enrollment_status,omitempty"`
}

type GetBRShopEnrollmentStatusResponse struct {
	BaseResponse
	Response struct {
		ShopList []BRShopEnrollmentStatus `json:"shop_list"`
	} `json:"response"`
}

func (s *BRService) QueryShopEnrollmentStatus(shopIDs []int64) (*GetBRShopEnrollmentStatusResponse, error) {
	ids := make([]string, len(shopIDs))
	for i, id := range shopIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"shop_id_list": stringsJoin(ids, ",")}
	result := &GetBRShopEnrollmentStatusResponse{}
	if err := s.client.DoGet(PathBRQueryShopEnrollmentStatus, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BRInvoiceError struct {
	ShopID      int64  `json:"shop_id"`
	ErrorCode   string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type GetBRShopInvoiceErrorResponse struct {
	BaseResponse
	Response struct {
		ShopList []BRInvoiceError `json:"shop_list"`
	} `json:"response"`
}

func (s *BRService) QueryShopInvoiceError(shopIDs []int64) (*GetBRShopInvoiceErrorResponse, error) {
	ids := make([]string, len(shopIDs))
	for i, id := range shopIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"shop_id_list": stringsJoin(ids, ",")}
	result := &GetBRShopInvoiceErrorResponse{}
	if err := s.client.DoGet(PathBRQueryShopInvoiceError, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BRShopBlockStatus struct {
	ShopID      int64  `json:"shop_id"`
	IsBlocked   bool   `json:"is_blocked"`
	BlockReason string `json:"block_reason,omitempty"`
}

type GetBRShopBlockStatusResponse struct {
	BaseResponse
	Response struct {
		ShopList []BRShopBlockStatus `json:"shop_list"`
	} `json:"response"`
}

func (s *BRService) QueryShopBlockStatus(shopIDs []int64) (*GetBRShopBlockStatusResponse, error) {
	ids := make([]string, len(shopIDs))
	for i, id := range shopIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"shop_id_list": stringsJoin(ids, ",")}
	result := &GetBRShopBlockStatusResponse{}
	if err := s.client.DoGet(PathBRQueryShopBlockStatus, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BRSKUBlockStatus struct {
	ItemID      int64  `json:"item_id"`
	ModelID     int64  `json:"model_id,omitempty"`
	IsBlocked   bool   `json:"is_blocked"`
	BlockReason string `json:"block_reason,omitempty"`
}

type GetBRSKUBlockStatusResponse struct {
	BaseResponse
	Response struct {
		SKUList []BRSKUBlockStatus `json:"sku_list"`
	} `json:"response"`
}

func (s *BRService) QuerySKUBlockStatus(itemIDs []int64) (*GetBRSKUBlockStatusResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetBRSKUBlockStatusResponse{}
	if err := s.client.DoGet(PathBRQuerySKUBlockStatus, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
