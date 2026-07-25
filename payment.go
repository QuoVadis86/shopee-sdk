package shopee

import (
	"context"
	"strconv"
)

type PaymentService struct {
	client *Client
}

func NewPaymentService(client *Client) *PaymentService {
	return &PaymentService{client: client}
}

type EscrowDetail struct {
	OrderSN           string            `json:"order_sn"`
	BuyerUserName     string            `json:"buyer_user_name"`
	ReturnOrderSNList []string          `json:"return_order_sn_list"`
	OrderIncome       EscrowOrderIncome `json:"order_income"`
	BuyerPaymentInfo  BuyerPaymentInfo  `json:"buyer_payment_info"`
}

type EscrowOrderItem struct {
	ItemID                    int64   `json:"item_id"`
	ItemName                  string  `json:"item_name"`
	ItemSKU                   string  `json:"item_sku"`
	ModelID                   int64   `json:"model_id"`
	ModelName                 string  `json:"model_name"`
	ModelSKU                  string  `json:"model_sku"`
	OriginalPrice             float64 `json:"original_price"`
	SellingPrice              float64 `json:"selling_price"`
	DiscountedPrice           float64 `json:"discounted_price"`
	SellerDiscount            float64 `json:"seller_discount"`
	ShopeeDiscount            float64 `json:"shopee_discount"`
	DiscountFromCoin          float64 `json:"discount_from_coin"`
	DiscountFromVoucherShopee float64 `json:"discount_from_voucher_shopee"`
	DiscountFromVoucherSeller float64 `json:"discount_from_voucher_seller"`
	ActivityType              string  `json:"activity_type"`
	ActivityID                int64   `json:"activity_id"`
	QuantityPurchased         int64   `json:"quantity_purchased"`
	AMSCommissionFee          float64 `json:"ams_commission_fee"`
}

type EscrowOrderIncome struct {
	EscrowAmount                        float64           `json:"escrow_amount"`
	BuyerTotalAmount                    float64           `json:"buyer_total_amount"`
	OrderOriginalPrice                  float64           `json:"order_original_price"`
	OriginalPrice                       float64           `json:"original_price"`
	OrderDiscountedPrice                float64           `json:"order_discounted_price"`
	OrderSellingPrice                   float64           `json:"order_selling_price"`
	OrderSellerDiscount                 float64           `json:"order_seller_discount"`
	SellerDiscount                      float64           `json:"seller_discount"`
	ShopeeDiscount                      float64           `json:"shopee_discount"`
	VoucherFromSeller                   float64           `json:"voucher_from_seller"`
	VoucherFromShopee                   float64           `json:"voucher_from_shopee"`
	Coins                               float64           `json:"coins"`
	BuyerPaidShippingFee                float64           `json:"buyer_paid_shipping_fee"`
	BuyerTransactionFee                 float64           `json:"buyer_transaction_fee"`
	CrossBorderTax                      float64           `json:"cross_border_tax"`
	PaymentPromotion                    float64           `json:"payment_promotion"`
	CommissionFee                       float64           `json:"commission_fee"`
	ServiceFee                          float64           `json:"service_fee"`
	SellerTransactionFee                float64           `json:"seller_transaction_fee"`
	SellerLostCompensation              float64           `json:"seller_lost_compensation"`
	SellerCoinCashBack                  float64           `json:"seller_coin_cash_back"`
	EscrowTax                           float64           `json:"escrow_tax"`
	EstimatedShippingFee                float64           `json:"estimated_shipping_fee"`
	FinalShippingFee                    float64           `json:"final_shipping_fee"`
	ActualShippingFee                   float64           `json:"actual_shipping_fee"`
	ShopeeShippingRebate                float64           `json:"shopee_shipping_rebate"`
	ShippingFeeDiscountFrom3PL          float64           `json:"shipping_fee_discount_from_3pl"`
	SellerShippingDiscount              float64           `json:"seller_shipping_discount"`
	DRCAdjustableRefund                 float64           `json:"drc_adjustable_refund"`
	CostOfGoodsSold                     float64           `json:"cost_of_goods_sold"`
	OriginalCostOfGoodsSold             float64           `json:"original_cost_of_goods_sold"`
	OriginalShopeeDiscount              float64           `json:"original_shopee_discount"`
	SellerReturnRefund                  float64           `json:"seller_return_refund"`
	CampaignFee                         float64           `json:"campaign_fee"`
	OrderAMSCommissionFee               float64           `json:"order_ams_commission_fee"`
	SellerOrderProcessingFee            float64           `json:"seller_order_processing_fee"`
	FBSFee                              float64           `json:"fbs_fee"`
	NetCommissionFee                    float64           `json:"net_commission_fee"`
	NetServiceFee                       float64           `json:"net_service_fee"`
	AdsEscrowTopUpOrTechnicalSupportFee float64           `json:"ads_escrow_top_up_fee_or_technical_support_fee"`
	Items                               []EscrowOrderItem `json:"items"`
}

type BuyerPaymentInfo struct {
	BuyerPaymentMethod  string  `json:"buyer_payment_method"`
	BuyerServiceFee     float64 `json:"buyer_service_fee"`
	BuyerTaxAmount      float64 `json:"buyer_tax_amount"`
	BuyerTotalAmount    float64 `json:"buyer_total_amount"`
	MerchantSubtotal    float64 `json:"merchant_subtotal"`
	SellerVoucher       float64 `json:"seller_voucher"`
	ShippingFee         float64 `json:"shipping_fee"`
	ShopeeVoucher       float64 `json:"shopee_voucher"`
	ShopeeCoinsRedeemed float64 `json:"shopee_coins_redeemed"`
}

type GetEscrowDetailResponse struct {
	BaseResponse
	Response *EscrowDetail `json:"response,omitempty"`
}

func (s *PaymentService) GetEscrowDetail(orderSN string) (*GetEscrowDetailResponse, error) {
	q := map[string]string{"order_sn": orderSN}
	result := &GetEscrowDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetEscrowDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type InstallmentStatus struct {
	Enabled bool `json:"enabled"`
}

func (s *PaymentService) SetShopInstallmentStatus(enabled bool) (*BaseResponse, error) {
	payload := map[string]any{"installment_status": enabled}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPaymentSetShopInstallment, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetInstallmentStatusResponse struct {
	BaseResponse
	Response *InstallmentStatus `json:"response,omitempty"`
}

func (s *PaymentService) GetShopInstallmentStatus() (*GetInstallmentStatusResponse, error) {
	result := &GetInstallmentStatusResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetShopInstallment, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PayoutDetail struct {
	PayoutID   int64   `json:"payout_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	CreateTime int64   `json:"create_time"`
}

type GetPayoutDetailResponse struct {
	BaseResponse
	Response *PayoutDetail `json:"response,omitempty"`
}

func (s *PaymentService) GetPayoutDetail(payoutID int64) (*GetPayoutDetailResponse, error) {
	q := map[string]string{"payout_id": strconv.FormatInt(payoutID, 10)}
	result := &GetPayoutDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetPayoutDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PaymentService) SetItemInstallmentStatus(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathPaymentSetItemInstallment, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PaymentService) GetItemInstallmentStatus(itemID int64) (*BaseResponse, error) {
	q := map[string]string{"item_id": strconv.FormatInt(itemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetItemInstallment, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PaymentMethod struct {
	PaymentMethodID   int64  `json:"payment_method_id"`
	PaymentMethodName string `json:"payment_method_name"`
	Enabled           bool   `json:"enabled"`
}

type GetPaymentMethodListResponse struct {
	BaseResponse
	Response struct {
		PaymentMethodList []PaymentMethod `json:"payment_method_list"`
	} `json:"response"`
}

func (s *PaymentService) GetPaymentMethodList() (*GetPaymentMethodListResponse, error) {
	result := &GetPaymentMethodListResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetPaymentMethodList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type WalletTransaction struct {
	TransactionID       string  `json:"transaction_id"`
	TransactionType     string  `json:"transaction_type"`
	Amount              float64 `json:"amount"`
	CurrentBalance      float64 `json:"current_balance,omitempty"`
	Status              string  `json:"status"`
	CreateTime          int64   `json:"create_time"`
	OrderSN             string  `json:"order_sn,omitempty"`
	RefundSN            string  `json:"refund_sn,omitempty"`
	TransactionFee      float64 `json:"transaction_fee,omitempty"`
	MoneyFlow           string  `json:"money_flow,omitempty"`
	WithdrawalType      string  `json:"withdrawal_type,omitempty"`
	TxnTitle            string  `json:"txn_title,omitempty"`
}

type GetWalletTransactionListResponse struct {
	BaseResponse
	Response struct {
		TransactionList []WalletTransaction `json:"transaction_list"`
		More            bool                `json:"more"`
		NextCursor      string              `json:"next_cursor"`
	} `json:"response"`
}

func (s *PaymentService) GetWalletTransactionList(pageSize, pageNo int, createTimeFrom, createTimeTo int64) (*GetWalletTransactionListResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if pageNo > 0 {
		q["page_no"] = strconv.Itoa(pageNo)
	}
	if createTimeFrom > 0 {
		q["create_time_from"] = strconv.FormatInt(createTimeFrom, 10)
	}
	if createTimeTo > 0 {
		q["create_time_to"] = strconv.FormatInt(createTimeTo, 10)
	}
	result := &GetWalletTransactionListResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetWalletTransactionList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type EscrowListItem struct {
	OrderSN          string  `json:"order_sn"`
	Amount           float64 `json:"payout_amount"`
	Status           string  `json:"-"`
	EscrowReleaseTime int64  `json:"escrow_release_time"`
	CreateTime       int64   `json:"create_time,omitempty"`
}

type GetEscrowListResponse struct {
	BaseResponse
	Response struct {
		EscrowList []EscrowListItem `json:"escrow_list"`
		More       bool             `json:"more"`
		NextCursor string           `json:"next_cursor"`
	} `json:"response"`
}

func (s *PaymentService) GetEscrowList(pageSize, pageNo int, timeFrom, timeTo int64) (*GetEscrowListResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if pageNo > 0 {
		q["page_no"] = strconv.Itoa(pageNo)
	}
	if timeFrom > 0 {
		q["release_time_from"] = strconv.FormatInt(timeFrom, 10)
	}
	if timeTo > 0 {
		q["release_time_to"] = strconv.FormatInt(timeTo, 10)
	}
	result := &GetEscrowListResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetEscrowList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PaymentService) GetPayoutInfo(payoutID int64) (*BaseResponse, error) {
	q := map[string]string{"payout_id": strconv.FormatInt(payoutID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetPayoutInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PaymentService) GetBillingTransactionInfo(transactionID string) (*BaseResponse, error) {
	q := map[string]string{"transaction_id": transactionID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetBillingTransaction, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type EscrowDetailBatchItem struct {
	EscrowDetail EscrowDetail `json:"escrow_detail"`
}

type GetEscrowDetailBatchResponse struct {
	BaseResponse
	Response []EscrowDetailBatchItem `json:"response"`
}

func (s *PaymentService) GetEscrowDetailBatch(orderSNs []string) (*GetEscrowDetailBatchResponse, error) {
	payload := map[string]any{"order_sn_list": orderSNs}
	result := &GetEscrowDetailBatchResponse{}
	if err := s.client.DoPost(context.Background(), PathPaymentGetEscrowDetailBatch, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type IncomeStatement struct {
	StatementID string  `json:"statement_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	CreateTime  int64   `json:"create_time"`
}

type GetIncomeStatementResponse struct {
	BaseResponse
	Response struct {
		StatementList []IncomeStatement `json:"statement_list"`
		More          bool              `json:"more"`
		NextCursor    string            `json:"next_cursor"`
	} `json:"response"`
}

func (s *PaymentService) GetIncomeStatement(dateFrom, dateTo int64, pageSize int, cursor string) (*GetIncomeStatementResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_size": strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	result := &GetIncomeStatementResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetIncomeStatement, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *PaymentService) GetIncomeReport(dateFrom, dateTo int64, pageSize int, cursor string) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_size": strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetIncomeReport, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type IncomeTotal struct {
	PendingAmount   float64 `json:"pending_amount"`
	ReleasedAmount  float64 `json:"released_amount"`
	ToReleaseAmount float64 `json:"to_release_amount"`
}

type IncomeOverview struct {
	LatestPayoutDate string      `json:"latest_payout_date"`
	TotalIncome      IncomeTotal `json:"total_income"`
}

type GetIncomeOverviewResponse struct {
	BaseResponse
	Response *IncomeOverview `json:"response,omitempty"`
}

func (s *PaymentService) GetIncomeOverview(dateFrom, dateTo int64) (*GetIncomeOverviewResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &GetIncomeOverviewResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetIncomeOverview, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type IncomeDetailListItem struct {
	PaymentMethod      string  `json:"payment_method,omitempty"`
	OrderSN            string  `json:"order_sn,omitempty"`
	Description        string  `json:"description,omitempty"`
	Status             string  `json:"status,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	EstimatedEscrowAmt float64 `json:"estimated_escrow_amount,omitempty"`
	EstimatedPayoutTime int64  `json:"estimated_payout_time,omitempty"`
	ToReleaseAmount    float64 `json:"to_release_amount,omitempty"`
	CreationDate       int64   `json:"creation_date,omitempty"`
	ReleasedAmount     float64 `json:"released_amount,omitempty"`
	ActualPayoutTime   int64   `json:"actual_payout_time,omitempty"`
}

type GetIncomeDetailResponse struct {
	BaseResponse
	IncomeDetailList *struct {
		NextPage *struct {
			Cursor   string `json:"cursor"`
			PageSize int    `json:"page_size"`
		} `json:"next_page,omitempty"`
		List []IncomeDetailListItem `json:"list,omitempty"`
	} `json:"income_detail_list,omitempty"`
}

func (s *PaymentService) GetIncomeDetail(dateFrom, dateTo string, incomeStatus int, pageSize int, cursor string) (*GetIncomeDetailResponse, error) {
	q := map[string]string{
		"date_from":     dateFrom,
		"date_to":       dateTo,
		"income_status": strconv.Itoa(incomeStatus),
		"page_size":     strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	result := &GetIncomeDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathPaymentGetIncomeDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
