package shopee

import (
	"context"
	"strconv"
)

type PromotionService struct {
	client *Client
}

func NewPromotionService(client *Client) *PromotionService {
	return &PromotionService{client: client}
}

type DiscountInfo struct {
	DiscountID  int64  `json:"discount_id"`
	Name        string `json:"discount_name"`
	Description string `json:"description,omitempty"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Status      string `json:"status,omitempty"`
}

type AddDiscountParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	StartTime   int64   `json:"start_time"`
	EndTime     int64   `json:"end_time"`
	Type        string  `json:"type"`
	Value       float64 `json:"value"`
	ItemList    []struct {
		ItemID  int64 `json:"item_id"`
		ModelID int64 `json:"model_id,omitempty"`
	} `json:"item_list"`
}

type AddDiscountResponse struct {
	BaseResponse
	Response struct {
		DiscountID int64 `json:"discount_id"`
	} `json:"response,omitempty"`
}

func (s *PromotionService) AddDiscount(params *AddDiscountParams) (*AddDiscountResponse, error) {
	result := &AddDiscountResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddDiscount, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddDiscountItemParams struct {
	DiscountID int64 `json:"discount_id"`
	ItemList   []struct {
		ItemID  int64 `json:"item_id"`
		ModelID int64 `json:"model_id,omitempty"`
	} `json:"item_list"`
}

func (s *PromotionService) AddDiscountItem(params *AddDiscountItemParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddDiscountItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteDiscount(discountID int64) (*BaseResponse, error) {
	payload := map[string]any{"discount_id": discountID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteDiscount, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteDiscountItem(params *AddDiscountItemParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteDiscountItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type DiscountModelDetail struct {
	ModelID          int64   `json:"model_id"`
	ModelName        string  `json:"model_name,omitempty"`
	NormalStock      int     `json:"model_normal_stock"`
	PromotionStock   int     `json:"model_promotion_stock"`
	OriginalPrice    float64 `json:"model_original_price"`
	PromotionPrice   float64 `json:"model_promotion_price"`
}

type DiscountItemDetail struct {
	ItemID             int64                `json:"item_id"`
	ItemName           string               `json:"item_name,omitempty"`
	NormalStock        int                  `json:"normal_stock"`
	ItemPromotionStock int                  `json:"item_promotion_stock"`
	ItemOriginalPrice  float64              `json:"item_original_price"`
	ItemPromotionPrice float64              `json:"item_promotion_price"`
	ModelList          []DiscountModelDetail `json:"model_list"`
	PurchaseLimit      int                  `json:"purchase_limit"`
}

type DiscountDetail struct {
	DiscountID   int64                `json:"discount_id"`
	DiscountName string               `json:"discount_name"`
	Status       string               `json:"status,omitempty"`
	StartTime    int64                `json:"start_time"`
	EndTime      int64                `json:"end_time"`
	ItemList     []DiscountItemDetail `json:"item_list"`
	More         bool                 `json:"more"`
}

type GetDiscountResponse struct {
	BaseResponse
	Response *DiscountDetail `json:"response,omitempty"`
}

func (s *PromotionService) GetDiscount(discountID int64, pageNo, pageSize int) (*GetDiscountResponse, error) {
	q := map[string]string{
		"discount_id": strconv.FormatInt(discountID, 10),
		"page_no":     strconv.Itoa(pageNo),
		"page_size":   strconv.Itoa(pageSize),
	}
	result := &GetDiscountResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetDiscount, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetDiscountListResponse struct {
	BaseResponse
	Response struct {
		DiscountList []DiscountInfo `json:"discount_list"`
		More         bool           `json:"more"`
	} `json:"response"`
}

func (s *PromotionService) GetDiscountList(pageSize, pageNo int, discountStatus string) (*GetDiscountListResponse, error) {
	q := map[string]string{
		"page_size":       strconv.Itoa(pageSize),
		"page_no":         strconv.Itoa(pageNo),
		"discount_status": discountStatus,
	}
	result := &GetDiscountListResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetDiscountList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateDiscount(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateDiscount, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateDiscountItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateDiscountItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetSIPDiscounts(discountID int64) (*BaseResponse, error) {
	q := map[string]string{"discount_id": strconv.FormatInt(discountID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetSIPDiscounts, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) SetSIPDiscount(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionSetSIPDiscount, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteSIPDiscount(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteSIPDiscount, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BundleDealInfo struct {
	BundleDealID int64  `json:"bundle_deal_id"`
	Name         string `json:"name"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Status       string `json:"status,omitempty"`
}

type AddBundleDealResponse struct {
	BaseResponse
	Response struct {
		BundleDealID int64 `json:"bundle_deal_id"`
	} `json:"response,omitempty"`
}

func (s *PromotionService) AddBundleDeal(params any) (*AddBundleDealResponse, error) {
	result := &AddBundleDealResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddBundleDeal, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddBundleDealItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddBundleDealItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetBundleDealListResponse struct {
	BaseResponse
	Response struct {
		BundleDealList []BundleDealInfo `json:"bundle_deal_list"`
		Total          int              `json:"total"`
	} `json:"response"`
}

func (s *PromotionService) GetBundleDealList(pageSize, pageNumber int, status string) (*GetBundleDealListResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
	}
	if status != "" {
		q["status"] = status
	}
	result := &GetBundleDealListResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetBundleDealList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetBundleDeal(bundleDealID int64) (*BaseResponse, error) {
	q := map[string]string{"bundle_deal_id": strconv.FormatInt(bundleDealID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetBundleDeal, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetBundleDealItem(bundleDealID int64) (*BaseResponse, error) {
	q := map[string]string{"bundle_deal_id": strconv.FormatInt(bundleDealID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetBundleDealItem, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateBundleDeal(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateBundleDeal, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateBundleDealItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateBundleDealItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteBundleDeal(bundleDealID int64) (*BaseResponse, error) {
	payload := map[string]any{"bundle_deal_id": bundleDealID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteBundleDeal, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteBundleDealItem(bundleDealID int64) (*BaseResponse, error) {
	payload := map[string]any{"bundle_deal_id": bundleDealID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteBundleDealItem, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddOnDealInfo struct {
	AddOnDealID int64  `json:"add_on_deal_id"`
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Status      string `json:"status,omitempty"`
}

type AddAddOnDealResponse struct {
	BaseResponse
	Response struct {
		AddOnDealID int64 `json:"add_on_deal_id"`
	} `json:"response,omitempty"`
}

func (s *PromotionService) AddAddOnDeal(params any) (*AddAddOnDealResponse, error) {
	result := &AddAddOnDealResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddAddOnDeal, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddAddOnDealMainItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddAddOnDealMainItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddAddOnDealSubItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddAddOnDealSubItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteAddOnDeal(addOnDealID int64) (*BaseResponse, error) {
	payload := map[string]any{"add_on_deal_id": addOnDealID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteAddOnDeal, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteAddOnDealMainItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteAddOnDealMain, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteAddOnDealSubItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteAddOnDealSub, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetAddOnDealListResponse struct {
	BaseResponse
	Response struct {
		AddOnDealList []AddOnDealInfo `json:"add_on_deal_list"`
		Total         int             `json:"total"`
	} `json:"response"`
}

func (s *PromotionService) GetAddOnDealList(pageSize, pageNumber int, promotionStatus string) (*GetAddOnDealListResponse, error) {
	q := map[string]string{
		"page_size":        strconv.Itoa(pageSize),
		"page_no":      strconv.Itoa(pageNumber),
		"promotion_status": promotionStatus,
	}
	result := &GetAddOnDealListResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetAddOnDealList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetAddOnDeal(addOnDealID int64) (*BaseResponse, error) {
	q := map[string]string{"add_on_deal_id": strconv.FormatInt(addOnDealID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetAddOnDeal, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetAddOnDealMainItem(addOnDealID int64) (*BaseResponse, error) {
	q := map[string]string{"add_on_deal_id": strconv.FormatInt(addOnDealID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetAddOnDealMainItem, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetAddOnDealSubItem(addOnDealID int64) (*BaseResponse, error) {
	q := map[string]string{"add_on_deal_id": strconv.FormatInt(addOnDealID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetAddOnDealSubItem, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateAddOnDeal(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateAddOnDeal, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateAddOnDealMainItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateAddOnDealMain, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateAddOnDealSubItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateAddOnDealSub, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) EndAddOnDeal(addOnDealID int64) (*BaseResponse, error) {
	payload := map[string]any{"add_on_deal_id": addOnDealID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionEndAddOnDeal, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddVoucherParams struct {
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Value      float64 `json:"value"`
	StartTime  int64   `json:"start_time"`
	EndTime    int64   `json:"end_time"`
	UsageLimit int     `json:"usage_limit,omitempty"`
}

type AddVoucherResponse struct {
	BaseResponse
	Response struct {
		VoucherID int64 `json:"voucher_id"`
	} `json:"response,omitempty"`
}

func (s *PromotionService) AddVoucher(params *AddVoucherParams) (*AddVoucherResponse, error) {
	result := &AddVoucherResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddVoucher, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteVoucher(voucherID int64) (*BaseResponse, error) {
	payload := map[string]any{"voucher_id": voucherID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteVoucher, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateVoucher(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateVoucher, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type VoucherInfo struct {
	VoucherID  int64  `json:"voucher_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      float64 `json:"value"`
	StartTime  int64   `json:"start_time"`
	EndTime    int64   `json:"end_time"`
	UsageLimit int     `json:"usage_limit,omitempty"`
	Status     string  `json:"status,omitempty"`
}

type GetVoucherResponse struct {
	BaseResponse
	Response *VoucherInfo `json:"response,omitempty"`
}

func (s *PromotionService) GetVoucher(voucherID int64) (*GetVoucherResponse, error) {
	q := map[string]string{"voucher_id": strconv.FormatInt(voucherID, 10)}
	result := &GetVoucherResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetVoucher, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetVoucherListResponse struct {
	BaseResponse
	Response struct {
		VoucherList []VoucherInfo `json:"voucher_list"`
		Total       int           `json:"total"`
	} `json:"response"`
}

func (s *PromotionService) GetVoucherList(pageSize, pageNumber int, status string) (*GetVoucherListResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
		"status":      status,
	}
	result := &GetVoucherListResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetVoucherList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetTimeSlotID(voucherID int64) (*BaseResponse, error) {
	q := map[string]string{"voucher_id": strconv.FormatInt(voucherID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetTimeSlotID, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type FlashSaleInfo struct {
	FlashSaleID int64  `json:"flash_sale_id"`
	Name        string `json:"name"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	Status      string `json:"status,omitempty"`
}

type CreateFlashSaleResponse struct {
	BaseResponse
	Response struct {
		FlashSaleID int64 `json:"flash_sale_id"`
	} `json:"response,omitempty"`
}

func (s *PromotionService) CreateShopFlashSale(params any) (*CreateFlashSaleResponse, error) {
	result := &CreateFlashSaleResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionCreateFlashSale, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetItemCriteria(flashSaleID int64) (*BaseResponse, error) {
	q := map[string]string{"flash_sale_id": strconv.FormatInt(flashSaleID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetItemCriteria, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddShopFlashSaleItems(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddFlashSaleItems, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetFlashSaleListResponse struct {
	BaseResponse
	Response struct {
		FlashSaleList []FlashSaleInfo `json:"flash_sale_list"`
		Total         int             `json:"total"`
	} `json:"response"`
}

func (s *PromotionService) GetShopFlashSaleList(status string, pageSize, pageNumber int) (*GetFlashSaleListResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
	}
	if status != "" {
		q["status"] = status
	}
	result := &GetFlashSaleListResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetFlashSaleList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetShopFlashSale(flashSaleID int64) (*BaseResponse, error) {
	q := map[string]string{"flash_sale_id": strconv.FormatInt(flashSaleID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetFlashSale, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetShopFlashSaleItems(flashSaleID int64) (*BaseResponse, error) {
	q := map[string]string{"flash_sale_id": strconv.FormatInt(flashSaleID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetFlashSaleItems, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateShopFlashSale(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateFlashSale, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateShopFlashSaleItems(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateFlashSaleItems, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteShopFlashSale(flashSaleID int64) (*BaseResponse, error) {
	payload := map[string]any{"flash_sale_id": flashSaleID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteFlashSale, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteShopFlashSaleItems(flashSaleID int64) (*BaseResponse, error) {
	payload := map[string]any{"flash_sale_id": flashSaleID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteFlashSaleItems, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddFollowPrize(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddFollowPrize, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteFollowPrize(followPrizeID int64) (*BaseResponse, error) {
	payload := map[string]any{"follow_prize_id": followPrizeID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteFollowPrize, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateFollowPrize(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateFollowPrize, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetFollowPrizeDetail(followPrizeID int64) (*BaseResponse, error) {
	q := map[string]string{"follow_prize_id": strconv.FormatInt(followPrizeID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetFollowPrizeDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetFollowPrizeList(status string, pageSize, pageNumber int) (*BaseResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
		"status":      status,
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetFollowPrizeList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetTopPicksList() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetTopPicksList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddTopPicks(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddTopPicks, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateTopPicks(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateTopPicks, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteTopPicks(topPicksID int64) (*BaseResponse, error) {
	payload := map[string]any{"top_picks_id": topPicksID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteTopPicks, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShopCategoryInfo struct {
	CategoryID int64  `json:"category_id"`
	Name       string `json:"name"`
	SortOrder  int    `json:"sort_order"`
}

type AddShopCategoryResponse struct {
	BaseResponse
	Response struct {
		CategoryID int64 `json:"category_id"`
	} `json:"response,omitempty"`
}

func (s *PromotionService) AddShopCategory(params any) (*AddShopCategoryResponse, error) {
	result := &AddShopCategoryResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddShopCategory, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetShopCategoryListResponse struct {
	BaseResponse
	Response struct {
		CategoryList []ShopCategoryInfo `json:"category_list"`
	} `json:"response"`
}

func (s *PromotionService) GetShopCategoryList(pageNo, pageSize int) (*GetShopCategoryListResponse, error) {
	q := map[string]string{
		"page_no":   strconv.Itoa(pageNo),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &GetShopCategoryListResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetShopCategoryList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteShopCategory(categoryID int64) (*BaseResponse, error) {
	payload := map[string]any{"category_id": categoryID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteShopCategory, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) UpdateShopCategory(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionUpdateShopCategory, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) AddShopCategoryItemList(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionAddShopCategoryItemList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) GetShopCategoryItemList(categoryID int64, pageSize, pageNumber int) (*BaseResponse, error) {
	q := map[string]string{
		"category_id": strconv.FormatInt(categoryID, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPromotionGetShopCategoryItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) DeleteShopCategoryItemList(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPromotionDeleteShopCategoryItemList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PromotionService) EndBundleDeal(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathBundleDealEndBundleDeal, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *PromotionService) EndDiscount(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathDiscountEndDiscount, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *PromotionService) EndFollowPrize(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathFollowPrizeEndFollowPrize, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *PromotionService) GetItemList(shopcategoryid int64) (*BaseResponse, error) {
    q := map[string]string{
		"shop_category_id": strconv.FormatInt(shopcategoryid, 10),
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathShopCategoryGetItemList, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *PromotionService) EndVoucher(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathVoucherEndVoucher, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}
