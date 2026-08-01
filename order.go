package shopee

import (
	"context"
	"encoding/json"
	"strconv"
)

type OrderService struct {
	client *Client
}

func NewOrderService(client *Client) *OrderService {
	return &OrderService{client: client}
}

type OrderListItem struct {
	OrderSN     string      `json:"order_sn"`
	OrderStatus OrderStatus `json:"order_status,omitempty"`
}

type GetOrderListResponse struct {
	BaseResponse
	Response struct {
		More       bool            `json:"more"`
		NextCursor string          `json:"next_cursor"`
		OrderList  []OrderListItem `json:"order_list"`
	} `json:"response"`
}

func (s *OrderService) GetOrderList(timeRangeField string, timeFrom, timeTo int64, pageSize int, cursor, orderStatus string) (*GetOrderListResponse, error) {
	q := map[string]string{
		"time_range_field": timeRangeField,
		"time_from":        strconv.FormatInt(timeFrom, 10),
		"time_to":          strconv.FormatInt(timeTo, 10),
		"page_size":        strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if orderStatus != "" {
		q["order_status"] = orderStatus
	}
	result := &GetOrderListResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetOrderList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type RecipientAddress struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Town        string `json:"town"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Region      string `json:"region"`
	ZipCode     string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

type OrderItem struct {
	ItemID               int64   `json:"item_id"`
	ItemName             string  `json:"item_name"`
	ItemSKU              string  `json:"item_sku"`
	ModelID              int64   `json:"model_id"`
	ModelName            string  `json:"model_name"`
	ModelSKU             string  `json:"model_sku"`
	ModelQuantity        int     `json:"model_quantity_purchased"`
	ModelOriginalPrice   float64 `json:"model_original_price"`
	ModelDiscountedPrice float64 `json:"model_discounted_price"`
	OrderItemID          int64   `json:"order_item_id"`
	Weight               float64 `json:"weight"`
	PromotionType        string  `json:"promotion_type,omitempty"`
	PromotionID          int64   `json:"promotion_id,omitempty"`
	ImageInfo            *struct {
		ImageURL string `json:"image_url"`
	} `json:"image_info,omitempty"`
}

type Package struct {
	PackageNumber      string `json:"package_number"`
	LogisticsStatus    string `json:"logistics_status"`
	LogisticsChannelID int64  `json:"logistics_channel_id"`
	ShippingCarrier    string `json:"shipping_carrier"`
	TrackingNumber     string `json:"tracking_number,omitempty"`
}

type OrderDetail struct {
	OrderSN              string            `json:"order_sn"`
	Region               string            `json:"region"`
	Currency             string            `json:"currency"`
	TotalAmount          float64           `json:"total_amount"`
	OrderStatus          OrderStatus       `json:"order_status"`
	ShippingCarrier      string            `json:"shipping_carrier"`
	PaymentMethod        string            `json:"payment_method"`
	EstimatedShippingFee float64           `json:"estimated_shipping_fee"`
	MessageToSeller      string            `json:"message_to_seller"`
	CreateTime           int64             `json:"create_time"`
	UpdateTime           int64             `json:"update_time"`
	DaysToShip           int               `json:"days_to_ship"`
	ShipByDate           int64             `json:"ship_by_date"`
	BuyerUserID          int64             `json:"buyer_user_id"`
	BuyerUsername        string            `json:"buyer_username"`
	RecipientAddress     *RecipientAddress `json:"recipient_address,omitempty"`
	ActualShippingFee    float64           `json:"actual_shipping_fee"`
	Note                 string            `json:"note,omitempty"`
	PayTime              int64             `json:"pay_time,omitempty"`
	ItemList             []OrderItem       `json:"item_list,omitempty"`
	PackageList          []Package         `json:"package_list,omitempty"`
	FulfillmentFlag      string            `json:"fulfillment_flag,omitempty"`
}

type GetOrderDetailResponse struct {
	BaseResponse
	Response struct {
		OrderList []OrderDetail `json:"order_list"`
	} `json:"response"`
}

func (s *OrderService) GetOrderDetail(orderSNs []string, optionalFields string) (*GetOrderDetailResponse, error) {
	q := map[string]string{
		"order_sn_list": stringsJoin(orderSNs, ","),
	}
	if optionalFields != "" {
		q["response_optional_fields"] = optionalFields
	}
	result := &GetOrderDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetOrderDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShipmentOrder struct {
	OrderSN       string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
}

type GetShipmentListResponse struct {
	BaseResponse
	Response struct {
		OrderList  []ShipmentOrder `json:"order_list"`
		More       bool            `json:"more"`
		NextCursor string          `json:"next_cursor"`
	} `json:"response"`
}

func (s *OrderService) GetShipmentList(pageSize int, cursor string) (*GetShipmentListResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if cursor != "" {
		q["cursor"] = cursor
	}
	result := &GetShipmentListResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetShipmentList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SearchPackageListParams struct {
	Pagination struct {
		PageSize int    `json:"page_size"`
		Cursor   string `json:"cursor,omitempty"`
	} `json:"pagination"`
	Filter *struct {
		PackageStatus   int `json:"package_status,omitempty"`
		FulfillmentType int `json:"fulfillment_type,omitempty"`
	} `json:"filter,omitempty"`
}

type SearchPackageListItem struct {
	OrderSN           string `json:"order_sn"`
	PackageNumber     string `json:"package_number"`
	FulfillmentStatus string `json:"fulfillment_status"`
	UpdateTime        int64  `json:"update_time"`
	CreateTime        int64  `json:"create_time"`
}

type SearchPackageListResponse struct {
	BaseResponse
	Response struct {
		More        bool                  `json:"more"`
		NextCursor  string                `json:"next_cursor"`
		PackageList []SearchPackageListItem `json:"package_list"`
	} `json:"response"`
}

func (s *OrderService) SearchPackageList(params *SearchPackageListParams) (*SearchPackageListResponse, error) {
	result := &SearchPackageListResponse{}
	if err := s.client.DoPost(context.Background(), PathOrderSearchPackageList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PackageDetail struct {
	OrderSN            string        `json:"order_sn"`
	PackageNumber      string        `json:"package_number"`
	FulfillmentStatus  string        `json:"fulfillment_status"`
	LogisticsChannelID int64         `json:"logistics_channel_id"`
	ShippingCarrier    string        `json:"shipping_carrier"`
	TrackingNumber     string        `json:"tracking_number,omitempty"`
	ItemList           []PackageItem `json:"item_list"`
}

type PackageItem struct {
	ItemID   int64 `json:"item_id"`
	ModelID  int64 `json:"model_id"`
	ModelQty int   `json:"model_quantity"`
}

type GetPackageDetailResponse struct {
	BaseResponse
	Response struct {
		PackageList []PackageDetail `json:"package_list"`
	} `json:"response"`
}

func (s *OrderService) GetPackageDetail(packageNumbers []string) (*GetPackageDetailResponse, error) {
	q := map[string]string{"package_number_list": stringsJoin(packageNumbers, ",")}
	result := &GetPackageDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetPackageDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SplitOrderParams struct {
	OrderSN  string `json:"order_sn"`
	ItemList []struct {
		ItemID  int64 `json:"item_id"`
		ModelID int64 `json:"model_id"`
	} `json:"item_list"`
}

func (s *OrderService) SplitOrder(params *SplitOrderParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathOrderSplitOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) UnsplitOrder(orderSN string) (*BaseResponse, error) {
	payload := map[string]any{"order_sn": orderSN}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathOrderUnsplitOrder, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CancelOrderParams struct {
	OrderSN      string `json:"order_sn"`
	CancelReason string `json:"cancel_reason"`
	ItemList     []struct {
		ItemID  int64 `json:"item_id"`
		ModelID int64 `json:"model_id"`
	} `json:"item_list,omitempty"`
}

type CancelOrderResponse struct {
	BaseResponse
	Response *struct {
		UpdateTime int64 `json:"update_time"`
	} `json:"response,omitempty"`
}

func (s *OrderService) CancelOrder(params *CancelOrderParams) (*CancelOrderResponse, error) {
	result := &CancelOrderResponse{}
	if err := s.client.DoPost(context.Background(), PathOrderCancelOrder, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SetNoteParams struct {
	OrderSN string `json:"order_sn"`
	Note    string `json:"note"`
}

func (s *OrderService) SetNote(params *SetNoteParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathOrderSetNote, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetPendingBuyerInvoiceOrderList(pageSize int, cursor string) (*GetOrderListResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if cursor != "" {
		q["cursor"] = cursor
	}
	result := &GetOrderListResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetPendingInvoiceList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetBuyerInvoiceInfo(orderSNs []string) (*BaseResponse, error) {
	q := map[string]string{"order_sn_list": stringsJoin(orderSNs, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetBuyerInvoiceInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetWarehouseFilterConfig() (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetWarehouseFilterCfg, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetBookingList(timeRangeField string, timeFrom, timeTo int64, pageSize int, cursor string) (*BaseResponse, error) {
	q := map[string]string{
		"time_range_field": timeRangeField,
		"page_size":        strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if timeFrom > 0 {
		q["time_from"] = strconv.FormatInt(timeFrom, 10)
	}
	if timeTo > 0 {
		q["time_to"] = strconv.FormatInt(timeTo, 10)
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetBookingList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetBookingDetail(bookingSNs []string) (*BaseResponse, error) {
	q := map[string]string{"booking_sn_list": stringsJoin(bookingSNs, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetBookingDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetFBSInvoicesResult(requestIDs []string) (*BaseResponse, error) {
	q := map[string]string{"request_id_list": stringsJoin(requestIDs, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetFBSInvoicesResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) GetEstimateCancelValue(orderSN string, partialCancelItems []map[string]any) (*BaseResponse, error) {
	q := map[string]string{
		"order_sn": orderSN,
	}
	if payload, err := json.Marshal(partialCancelItems); err == nil {
		q["partial_cancel_item_list"] = string(payload)
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathOrderGetEstimateCancelValue, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *OrderService) DownloadFbsInvoices(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathOrderDownloadFbsInvoices, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *OrderService) DownloadInvoiceDoc(ordersn string) (*BaseResponse, error) {
    q := map[string]string{
		"order_sn": ordersn,
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathOrderDownloadInvoiceDoc, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *OrderService) GenerateFbsInvoices(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathOrderGenerateFbsInvoices, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *OrderService) HandleBuyerCancellation(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathOrderHandleBuyerCancellation, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *OrderService) HandlePrescriptionCheck(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathOrderHandlePrescriptionCheck, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *OrderService) UploadInvoiceDoc(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathOrderUploadInvoiceDoc, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}
