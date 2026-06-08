package shopee

// BaseResponse is embedded in all API responses.
type BaseResponse struct {
	RequestID string `json:"request_id"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Warning   string `json:"warning,omitempty"`
}

func (r *BaseResponse) HasError() bool {
	return r.Error != ""
}

// --- Common Enums ---

type ItemStatus string

const (
	ItemStatusNormal       ItemStatus = "NORMAL"
	ItemStatusBanned       ItemStatus = "BANNED"
	ItemStatusUnlist       ItemStatus = "UNLIST"
	ItemStatusReviewing    ItemStatus = "REVIEWING"
	ItemStatusSellerDelete ItemStatus = "SELLER_DELETE"
	ItemStatusShopeeDelete ItemStatus = "SHOPEE_DELETE"
)

type OrderStatus string

const (
	OrderStatusUnpaid      OrderStatus = "UNPAID"
	OrderStatusReadyToShip OrderStatus = "READY_TO_SHIP"
	OrderStatusProcessed   OrderStatus = "PROCESSED"
	OrderStatusRetryShip   OrderStatus = "RETRY_SHIP"
	OrderStatusShipped     OrderStatus = "SHIPPED"
	OrderStatusToConfirm   OrderStatus = "TO_CONFIRM_RECEIVE"
	OrderStatusInCancel    OrderStatus = "IN_CANCEL"
	OrderStatusCancelled   OrderStatus = "CANCELLED"
	OrderStatusToReturn    OrderStatus = "TO_RETURN"
	OrderStatusCompleted   OrderStatus = "COMPLETED"
)

type FulfillmentStatus string

const (
	FulfillmentUnfulfilled  FulfillmentStatus = "UNFULFILLED"
	FulfillmentFulfilled    FulfillmentStatus = "FULFILLED"
	FulfillmentPartiallyFulfilled FulfillmentStatus = "PARTIALLY_FULFILLED"
)

type LogisticsStatus string

const (
	LogisticsStatusInTransit    LogisticsStatus = "IN_TRANSIT"
	LogisticsStatusDelivered    LogisticsStatus = "DELIVERED"
	LogisticsStatusFailed       LogisticsStatus = "FAILED"
	LogisticsStatusReturned     LogisticsStatus = "RETURNED"
	LogisticsStatusCancelled    LogisticsStatus = "CANCELLED"
)

type ReturnStatus string

const (
	ReturnStatusPending        ReturnStatus = "PENDING"
	ReturnStatusAccepted       ReturnStatus = "ACCEPTED"
	ReturnStatusRejected       ReturnStatus = "REJECTED"
	ReturnStatusProcessing     ReturnStatus = "PROCESSING"
	ReturnStatusCompleted      ReturnStatus = "COMPLETED"
	ReturnStatusCancelled      ReturnStatus = "CANCELLED"
)

type DiscountType string

const (
	DiscountTypePercentage DiscountType = "PERCENTAGE"
	DiscountTypeFixed      DiscountType = "FIXED"
)

type VoucherType string

const (
	VoucherTypeFreeShipping  VoucherType = "FREE_SHIPPING"
	VoucherTypeDiscount      VoucherType = "DISCOUNT"
)
