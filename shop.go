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

type GetShopInfoResponse struct {
	BaseResponse
	Response *ShopInfo `json:"response,omitempty"`
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
	ShopName        string `json:"shop_name"`
	ShopDescription string `json:"shop_description,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
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
	WarehouseID   int64  `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	Region        string `json:"region"`
	Address       string `json:"address"`
}

type GetWarehouseDetailResponse struct {
	BaseResponse
	Response *WarehouseDetail `json:"response,omitempty"`
}

func (s *ShopService) GetWarehouseDetail(warehouseID int64) (*GetWarehouseDetailResponse, error) {
	q := map[string]string{"warehouse_id": strconv.FormatInt(warehouseID, 10)}
	result := &GetWarehouseDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetWarehouseDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShopNotification struct {
	NotificationID   int64  `json:"notification_id"`
	NotificationType string `json:"notification_type"`
	Content          string `json:"content"`
	CreateTime       int64  `json:"create_time"`
	IsRead           bool   `json:"is_read"`
}

type GetShopNotificationResponse struct {
	BaseResponse
	Response struct {
		NotificationList []ShopNotification `json:"notification_list"`
		TotalCount       int                `json:"total_count"`
	} `json:"response"`
}

func (s *ShopService) GetShopNotification(pageSize, pageNumber int) (*GetShopNotificationResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
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

func (s *ShopService) GetAuthorisedResellerBrand() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetAuthResellerBrand, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ShopService) GetBROnboardingInfo() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetBROnboardingInfo, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type HolidayMode struct {
	IsHolidayModeEnabled bool  `json:"is_holiday_mode_enabled"`
	StartTime            int64 `json:"start_time,omitempty"`
	EndTime              int64 `json:"end_time,omitempty"`
}

type GetHolidayModeResponse struct {
	BaseResponse
	Response *HolidayMode `json:"response,omitempty"`
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

func (s *ShopService) SetHolidayMode(params *HolidayMode) (*BaseResponse, error) {
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
	Response *MerchantInfo `json:"response,omitempty"`
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

func (s *ShopService) GetShopListByMerchant() (*BaseResponse, error) {
	result := &BaseResponse{}
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

func (s *ShopService) GetTotalBalance() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathShopGetTotalBalance, map[string]string{}, result); err != nil {
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
