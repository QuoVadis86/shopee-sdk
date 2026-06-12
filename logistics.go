package shopee

import (
	"context"
	"strconv"
)

type LogisticsService struct {
	client *Client
}

func NewLogisticsService(client *Client) *LogisticsService {
	return &LogisticsService{client: client}
}

type DropoffBranch struct {
	BranchID int64  `json:"branch_id"`
	Region   string `json:"region"`
	State    string `json:"state"`
	City     string `json:"city"`
	Address  string `json:"address"`
	ZipCode  string `json:"zipcode"`
}

type PickupAddress struct {
	AddressID int64  `json:"address_id"`
	Region    string `json:"region"`
	State     string `json:"state"`
	City      string `json:"city"`
	District  string `json:"district"`
	Address   string `json:"address"`
}

type ShippingParameterResponse struct {
	BaseResponse
	Response *struct {
		InfoNeededInfo []string `json:"info_needed,omitempty"`
		Dropoff        *struct {
			BranchList []DropoffBranch `json:"branch_list"`
		} `json:"dropoff,omitempty"`
		Pickup *struct {
			AddressList []PickupAddress `json:"address_list"`
		} `json:"pickup,omitempty"`
	} `json:"response,omitempty"`
}

func (s *LogisticsService) GetShippingParameter(orderSN, packageNumber string) (*ShippingParameterResponse, error) {
	q := map[string]string{"order_sn": orderSN}
	if packageNumber != "" {
		q["package_number"] = packageNumber
	}
	result := &ShippingParameterResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetShippingParam, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetMassShippingParameter(orderList []string) (*BaseResponse, error) {
	payload := map[string]any{"order_list": orderList}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsGetMassShippingParam, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShipPickup struct {
	AddressID    int64  `json:"address_id"`
	PickupTimeID string `json:"pickup_time_id,omitempty"`
}

type ShipDropoff struct {
	BranchID       int64  `json:"branch_id,omitempty"`
	SenderRealName string `json:"sender_real_name,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
}

type ShipNonIntegrated struct {
	TrackingNumber string `json:"tracking_number"`
}

type ShipOrderParams struct {
	OrderSN       string           `json:"order_sn"`
	PackageNumber string           `json:"package_number,omitempty"`
	Pickup        *ShipPickup      `json:"pickup,omitempty"`
	Dropoff       *ShipDropoff     `json:"dropoff,omitempty"`
	NonIntegrated *ShipNonIntegrated `json:"non_integrated,omitempty"`
}

func (s *LogisticsService) ShipOrder(params *ShipOrderParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsShipOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) MassShipOrder(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsMassShipOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) UpdateShippingOrder(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsUpdateShippingOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetTrackingNumberResponse struct {
	BaseResponse
	Response *struct {
		TrackingNumber          string `json:"tracking_number"`
		FirstMileTrackingNumber string `json:"first_mile_tracking_number,omitempty"`
		LastMileTrackingNumber  string `json:"last_mile_tracking_number,omitempty"`
		PlpNumber               string `json:"plp_number,omitempty"`
	} `json:"response,omitempty"`
}

func (s *LogisticsService) GetTrackingNumber(orderSN, packageNumber string) (*GetTrackingNumberResponse, error) {
	q := map[string]string{"order_sn": orderSN}
	if packageNumber != "" {
		q["package_number"] = packageNumber
	}
	result := &GetTrackingNumberResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetTrackingNumber, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetMassTrackingNumber(orderSNs []string) (*BaseResponse, error) {
	payload := map[string]any{"order_sn_list": orderSNs}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsGetMassTrackingNumber, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetShippingDocParamResponse struct {
	BaseResponse
	Response *struct {
		SuggestedDocType   string   `json:"suggested_shipping_document_type"`
		SelectableDocTypes []string `json:"selectable_shipping_document_type"`
	} `json:"response,omitempty"`
}

type orderSNItem struct {
	OrderSN string `json:"order_sn"`
}

func (s *LogisticsService) GetShippingDocumentParameter(orderSNs []string) (*GetShippingDocParamResponse, error) {
	list := make([]orderSNItem, len(orderSNs))
	for i, sn := range orderSNs {
		list[i] = orderSNItem{OrderSN: sn}
	}
	result := &GetShippingDocParamResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsGetShippingDocParam, map[string]any{"order_list": list}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) CreateShippingDocument(orderSNs []string, docType string) (*BaseResponse, error) {
	list := make([]orderSNItem, len(orderSNs))
	for i, sn := range orderSNs {
		list[i] = orderSNItem{OrderSN: sn}
	}
	payload := map[string]any{
		"order_list":             list,
		"shipping_document_type": docType,
	}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsCreateShippingDoc, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetShippingDocumentResult(orderListJSON, docType string) (*BaseResponse, error) {
	q := map[string]string{
		"order_list":             orderListJSON,
		"shipping_document_type": docType,
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetShippingDocResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetShippingDocumentDataInfo(resultID string) (*BaseResponse, error) {
	q := map[string]string{"shipping_document_result_id": resultID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetShippingDocDataInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type TrackingEvent struct {
	UpdateTime      int64  `json:"update_time"`
	Description     string `json:"description"`
	LogisticsStatus string `json:"logistics_status"`
}

type GetTrackingInfoResponse struct {
	BaseResponse
	Response *struct {
		OrderSN         string          `json:"order_sn"`
		PackageNumber   string          `json:"package_number"`
		LogisticsStatus string          `json:"logistics_status"`
		TrackingInfo    []TrackingEvent `json:"tracking_info"`
	} `json:"response,omitempty"`
}

func (s *LogisticsService) GetTrackingInfo(orderSN, packageNumber string) (*GetTrackingInfoResponse, error) {
	q := map[string]string{"order_sn": orderSN}
	if packageNumber != "" {
		q["package_number"] = packageNumber
	}
	result := &GetTrackingInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetTrackingInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type Address struct {
	AddressID int64  `json:"address_id"`
	Region    string `json:"region"`
	State     string `json:"state"`
	City      string `json:"city"`
	District  string `json:"district"`
	Address   string `json:"address"`
	ZipCode   string `json:"zipcode"`
}

type GetAddressListResponse struct {
	BaseResponse
	Response struct {
		AddressList []Address `json:"address_list"`
	} `json:"response"`
}

func (s *LogisticsService) GetAddressList() (*GetAddressListResponse, error) {
	result := &GetAddressListResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetAddressList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) SetAddressConfig(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsSetAddressConfig, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) UpdateAddress(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsUpdateAddress, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) DeleteAddress(addressID int64) (*BaseResponse, error) {
	payload := map[string]any{"address_id": addressID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsDeleteAddress, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type LogisticsChannel struct {
	LogisticsChannelID   int64  `json:"logistics_channel_id"`
	LogisticsChannelName string `json:"logistics_channel_name"`
	Enabled              bool   `json:"enabled"`
	CODenabled           bool   `json:"cod_enabled"`
	FeeType              string `json:"fee_type"`
}

type GetChannelListResponse struct {
	BaseResponse
	Response struct {
		LogisticsChannelList []LogisticsChannel `json:"logistics_channel_list"`
	} `json:"response"`
}

func (s *LogisticsService) GetChannelList() (*GetChannelListResponse, error) {
	result := &GetChannelListResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetChannelList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) UpdateChannel(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsUpdateChannel, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetOperatingHours(addressID int64, date string) (*BaseResponse, error) {
	q := map[string]string{
		"address_id": strconv.FormatInt(addressID, 10),
		"date":       date,
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetOperatingHours, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetOperatingHourRestrictions() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetOpHoursRestrictions, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) UpdateOperatingHours(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsUpdateOperatingHours, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) DeleteSpecialOperatingHour(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsDeleteSpecialOpHour, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) BatchUpdateTPFWarehouseTrackingStatus(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsBatchUpdateTPFTracking, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) BatchShipOrder(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsBatchShipOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) UpdateTrackingStatus(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsUpdateTrackingStatus, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetBookingShippingParameter(bookingID int64) (*BaseResponse, error) {
	q := map[string]string{"booking_id": strconv.FormatInt(bookingID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetBookingShippingParam, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) ShipBooking(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsShipBooking, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetBookingTrackingNumber(bookingID int64) (*BaseResponse, error) {
	q := map[string]string{"booking_id": strconv.FormatInt(bookingID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetBookingTrackingNum, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetBookingShippingDocumentParameter(bookingID int64) (*BaseResponse, error) {
	q := map[string]string{"booking_id": strconv.FormatInt(bookingID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetBookingShipDocParam, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) CreateBookingShippingDocument(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsCreateBookingShipDoc, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetBookingShippingDocumentResult(bookingID int64) (*BaseResponse, error) {
	q := map[string]string{"booking_id": strconv.FormatInt(bookingID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetBookingShipDocResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetBookingShippingDocumentDataInfo(resultID string) (*BaseResponse, error) {
	q := map[string]string{"shipping_document_result_id": resultID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetBookingShipDocData, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetBookingTrackingInfo(bookingID int64) (*BaseResponse, error) {
	q := map[string]string{"booking_id": strconv.FormatInt(bookingID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetBookingTrackingInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) CreateShippingDocumentJob(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsCreateShipDocJob, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetShippingDocumentJobStatus(jobID string) (*BaseResponse, error) {
	q := map[string]string{"job_id": jobID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetShipDocJobStatus, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) UpdateSelfCollectionOrderLogistics(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsUpdateSelfCollectOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetMartPackagingInfo(martItemID int64) (*BaseResponse, error) {
	q := map[string]string{"mart_item_id": strconv.FormatInt(martItemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetMartPackagingInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) SetMartPackagingInfo(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsSetMartPackagingInfo, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) CheckPolygonUpdateStatus(polygonVersion string) (*BaseResponse, error) {
	q := map[string]string{"polygon_version": polygonVersion}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsCheckPolygonUpdate, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PauseStatus struct {
	IsPaused   bool   `json:"is_paused"`
	PauseType  string `json:"pause_type,omitempty"`
	Region     string `json:"region,omitempty"`
}

type GetPauseStatusResponse struct {
	BaseResponse
	Response *PauseStatus `json:"response,omitempty"`
}

func (s *LogisticsService) GetPauseStatus() (*GetPauseStatusResponse, error) {
	result := &GetPauseStatusResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetPauseStatus, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) SetPauseStatus(params *PauseStatus) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLogisticsSetPauseStatus, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetUnbindOrderList(orderSNs []string) (*BaseResponse, error) {
	q := map[string]string{"order_sn_list": stringsJoin(orderSNs, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetUnbindOrderList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetDetail(trackingNumber string, channelID int64) (*BaseResponse, error) {
	q := map[string]string{
		"tracking_number": trackingNumber,
		"channel_id":      strconv.FormatInt(channelID, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetTrackingNumberList(orderSNs []string) (*BaseResponse, error) {
	q := map[string]string{"order_sn_list": stringsJoin(orderSNs, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetTrackingNumList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetWaybill(orderSNs []string, channelID int64) (*BaseResponse, error) {
	q := map[string]string{
		"order_sn_list": stringsJoin(orderSNs, ","),
		"channel_id":    strconv.FormatInt(channelID, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetWaybill, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetCourierDeliveryChannelList(addressID int64, orderAmount float64) (*BaseResponse, error) {
	q := map[string]string{
		"address_id":   strconv.FormatInt(addressID, 10),
		"order_amount": strconv.FormatFloat(orderAmount, 'f', 2, 64),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetCourierChannelList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetTransitWarehouseList() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetTransitWarehouseList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetCourierDeliveryDetail(trackingNumber string) (*BaseResponse, error) {
	q := map[string]string{"tracking_number": trackingNumber}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetCourierDeliveryDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetCourierDeliveryWaybill(trackingNumber string) (*BaseResponse, error) {
	q := map[string]string{"tracking_number": trackingNumber}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetCourierDeliveryWaybill, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LogisticsService) GetCourierDeliveryTrackingNumberList(orderSNs []string) (*BaseResponse, error) {
	q := map[string]string{"order_sn_list": stringsJoin(orderSNs, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLogisticsGetCourierTrackingList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
