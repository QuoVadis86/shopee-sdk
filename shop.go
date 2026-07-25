package shopee

import (
	"context"
	"strconv"
)

type ShopService struct {
	client *Client
}

func NewShopService(client *Client) *ShopService {
	return &ShopService{client: client}
}

type ShopInfo struct {
	ShopID      int64  `json:"shop_id"`
	ShopName    string `json:"shop_name"`
	Region      string `json:"region"`
	Status      string `json:"status"`
	CreatedTime int64  `json:"create_time"`
}

type SIPAffiShop struct {
	AffiShopID int64  `json:"affi_shop_id"`
	Region     string `json:"region"`
}

type LinkedDirectShop struct {
	DirectShopID     int64  `json:"direct_shop_id"`
	DirectShopRegion string `json:"direct_shop_region"`
}

type OutletShopInfo struct {
	OutletShopID int64 `json:"outlet_shop_id"`
}

type GetShopInfoResponse struct {
	BaseResponse

	ShopName            string              `json:"shop_name"`
	Region              string              `json:"region"`
	Status              string              `json:"status"`
	AuthTime            int64               `json:"auth_time"`
	ExpireTime          int64               `json:"expire_time"`
	IsCB                bool                `json:"is_cb"`
	MerchantID          int64               `json:"merchant_id"`
	IsSIP               bool                `json:"is_sip"`
	SIPAffiShops        []SIPAffiShop       `json:"sip_affi_shops,omitempty"`
	IsMainShop          bool                `json:"is_main_shop"`
	IsDirectShop        bool                `json:"is_direct_shop"`
	LinkedMainShopID    int64               `json:"linked_main_shop_id,omitempty"`
	LinkedDirectShops   []LinkedDirectShop  `json:"linked_direct_shop_list,omitempty"`
	IsUpgradedCBSC          bool                `json:"is_upgraded_cbsc"`
	ShopFulfillmentFlag     string              `json:"shop_fulfillment_flag"`
	IsOneAWB                bool                `json:"is_one_awb,omitempty"`
	MartOutletStructureType string              `json:"mart_outlet_structure_type,omitempty"`
	IsMartShop          bool                `json:"is_mart_shop,omitempty"`
	IsOutletShop        bool                `json:"is_outlet_shop,omitempty"`
	MartShopID          int64               `json:"mart_shop_id,omitempty"`
	OutletShopInfoList  []OutletShopInfo    `json:"outlet_shop_info_list,omitempty"`
}

func (s *ShopService) GetShopInfo() (*GetShopInfoResponse, error) {
	result := &GetShopInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShopProfile struct {
	ShopName      string `json:"shop_name,omitempty"`
	ShopLogo      string `json:"shop_logo,omitempty"`
	Description   string `json:"description,omitempty"`
	InvoiceIssuer string `json:"invoice_issuer,omitempty"`
}

type GetProfileResponse struct {
	BaseResponse
	Response *ShopProfile `json:"response,omitempty"`
}

func (s *ShopService) GetProfile() (*GetProfileResponse, error) {
	result := &GetProfileResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetProfile, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) UpdateProfile(params *ShopProfile) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathShopUpdateProfile, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type WarehouseDetail struct {
	WarehouseID       int64  `json:"warehouse_id"`
	WarehouseName     string `json:"warehouse_name"`
	WarehouseType     int    `json:"warehouse_type,omitempty"`
	LocationID        string `json:"location_id,omitempty"`
	AddressID         int64  `json:"address_id,omitempty"`
	Region            string `json:"region,omitempty"`
	State             string `json:"state,omitempty"`
	City              string `json:"city,omitempty"`
	Address           string `json:"address,omitempty"`
	ZipCode           string `json:"zipcode,omitempty"`
	District          string `json:"district,omitempty"`
	Town              string `json:"town,omitempty"`
	StateCode         string `json:"state_code,omitempty"`
	HolidayModeState  int    `json:"holiday_mode_state,omitempty"`
}

type GetWarehouseDetailResponse struct {
	BaseResponse
	Response []WarehouseDetail `json:"response,omitempty"`
}

func (s *ShopService) GetWarehouseDetail(warehouseType ...int) (*GetWarehouseDetailResponse, error) {
	q := map[string]string{}
	if len(warehouseType) > 0 && warehouseType[0] > 0 {
		q["warehouse_type"] = strconv.Itoa(warehouseType[0])
	}
	result := &GetWarehouseDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetWarehouseDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type NotificationData struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
	URL        string `json:"url,omitempty"`
}

type GetShopNotificationResponse struct {
	BaseResponse
	Cursor int64             `json:"cursor"`
	Data   *NotificationData `json:"data,omitempty"`
}

func (s *ShopService) GetShopNotification(cursor int64, pageSize ...int) (*GetShopNotificationResponse, error) {
	q := map[string]string{}
	if cursor > 0 {
		q["cursor"] = strconv.FormatInt(cursor, 10)
	}
	if len(pageSize) > 0 && pageSize[0] > 0 {
		q["page_size"] = strconv.Itoa(pageSize[0])
	}
	result := &GetShopNotificationResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetNotification, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AuthorisedBrand struct {
	BrandID   int64  `json:"brand_id"`
	BrandName string `json:"brand_name"`
}

type GetAuthorisedResellerBrandResponse struct {
	BaseResponse
	Response *struct {
		IsAuthorisedReseller bool               `json:"is_authorised_reseller"`
		TotalCount           int                `json:"total_count"`
		More                 bool               `json:"more"`
		AuthorisedBrandList  []AuthorisedBrand  `json:"authorised_brand_list"`
	} `json:"response,omitempty"`
}

func (s *ShopService) GetAuthorisedResellerBrand(pageNo, pageSize int) (*GetAuthorisedResellerBrandResponse, error) {
	q := map[string]string{
		"page_no":   strconv.Itoa(pageNo),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &GetAuthorisedResellerBrandResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetAuthResellerBrand, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BillingAddress struct {
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Address      string `json:"address,omitempty"`
	ZipCode      string `json:"zipcode,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
}

type BROnboardingInfo struct {
	TaxIDType        int              `json:"tax_id_type,omitempty"`
	TaxID            string           `json:"tax_id,omitempty"`
	OnboardingStatus int              `json:"onboarding_status,omitempty"`
	OnboardingPassed bool             `json:"onboarding_passed,omitempty"`
	Name             string           `json:"name,omitempty"`
	CPFID            string           `json:"cpf_id,omitempty"`
	Birthday         int64            `json:"birthday,omitempty"`
	BirthdayStr      string           `json:"birthday_str,omitempty"`
	LegalEntityName  string           `json:"legal_entity_name,omitempty"`
	CNPJID           string           `json:"cnpj_id,omitempty"`
	StateRegistration string          `json:"state_registration,omitempty"`
	CNAEMain         string           `json:"cnae_main,omitempty"`
	CNAESecondary    string           `json:"cnae_secondary,omitempty"`
	MEICheck         string           `json:"mei_check,omitempty"`
	SubmissionTime   int32            `json:"submission_time,omitempty"`
	Nationality      string           `json:"nationality,omitempty"`
	BillingAddress   *BillingAddress  `json:"billing_address,omitempty"`
}

type GetBROnboardingInfoResponse struct {
	BaseResponse
	Response *BROnboardingInfo `json:"response,omitempty"`
}

func (s *ShopService) GetBROnboardingInfo() (*GetBROnboardingInfoResponse, error) {
	result := &GetBROnboardingInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetBROnboardingInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type HolidayMode struct {
	HolidayModeOn          bool   `json:"holiday_mode_on"`
	HolidayModeMtime       int64  `json:"holiday_mode_mtime,omitempty"`
	HolidayModeType        int32  `json:"holiday_mode_type,omitempty"`
	HolidayModeStartTime   int64  `json:"holiday_mode_start_time,omitempty"`
	HolidayModeEndTime     int64  `json:"holiday_mode_end_time,omitempty"`
	HolidayModeDescription string `json:"holiday_mode_description,omitempty"`
	DebugMsg               string `json:"debug_msg,omitempty"`
}

type GetHolidayModeResponse struct {
	BaseResponse
	Response *HolidayMode `json:"response,omitempty"`
}

type SetHolidayModeRequest struct {
	HolidayModeOn bool `json:"holiday_mode_on"`
}

func (s *ShopService) GetHolidayMode() (*GetHolidayModeResponse, error) {
	result := &GetHolidayModeResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetHolidayMode, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) SetHolidayMode(params *SetHolidayModeRequest) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathShopSetHolidayMode, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type MerchantInfo struct {
	MerchantID   int64  `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	Region       string `json:"region"`
}

type GetMerchantInfoResponse struct {
	BaseResponse

	MerchantName     string `json:"merchant_name"`
	Region           string `json:"region,omitempty"`
	MerchantRegion   string `json:"merchant_region,omitempty"`
	MerchantCurrency string `json:"merchant_currency,omitempty"`
	IsCNSC           bool   `json:"is_cnsc"`
	IsUpgradedCBSC   bool   `json:"is_upgraded_cbsc"`
	AuthTime         int64  `json:"auth_time"`
	ExpireTime       int64  `json:"expire_time"`
}

func (s *ShopService) GetMerchantInfo() (*GetMerchantInfoResponse, error) {
	result := &GetMerchantInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetMerchantInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type MerchantShop struct {
	ShopID int64 `json:"shop_id"`
}

type GetShopListByMerchantResponse struct {
	BaseResponse
	ShopList []MerchantShop `json:"shop_list"`
	More     bool           `json:"more"`
	IsCNSC   bool           `json:"is_cnsc"`
}

func (s *ShopService) GetShopListByMerchant() (*GetShopListByMerchantResponse, error) {
	result := &GetShopListByMerchantResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetListByMerchant, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetMerchantWarehouseLocationList() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetMerchantWarehouseLocations, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetMerchantWarehouseList() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetMerchantWarehouseList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetWarehouseEligibleShopList() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetWarehouseEligibleShopList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetMerchantPrepaidAccountList() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetMerchantPrepaidAccountList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShopPerformanceMetric struct {
	MetricName  string  `json:"metric_name"`
	MetricValue float64 `json:"metric_value"`
	MetricUnit  string  `json:"metric_unit,omitempty"`
}

type GetShopPerformanceResponse struct {
	BaseResponse
	Response struct {
		PerformanceList []ShopPerformanceMetric `json:"performance_list"`
	} `json:"response"`
}

func (s *ShopService) GetShopPerformance() (*GetShopPerformanceResponse, error) {
	result := &GetShopPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetPerformance, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetMetricSourceDetail(metricName string) (*BaseResponse, error) {
	q := map[string]string{"metric_name": metricName}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetMetricSourceDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PenaltyPointInfo struct {
	PenaltyID    int64  `json:"penalty_id"`
	PenaltyPoint int    `json:"penalty_point"`
	Reason       string `json:"reason"`
	CreatedTime  int64  `json:"create_time"`
}

type GetPenaltyHistoryResponse struct {
	BaseResponse
	Response struct {
		PenaltyList []PenaltyPointInfo `json:"penalty_list"`
	} `json:"response"`
}

func (s *ShopService) GetPenaltyPointHistory() (*GetPenaltyHistoryResponse, error) {
	result := &GetPenaltyHistoryResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetPenaltyHistory, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetPunishmentHistory() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetPunishmentHistory, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetListingsWithIssues(pageSize, pageNumber int) (*BaseResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetListingsWithIssues, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetLateOrders(pageSize int, cursor string) (*BaseResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if cursor != "" {
		q["cursor"] = cursor
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetLateOrders, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetToggleInfo() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetToggleInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
