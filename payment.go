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
	OrderSN    string  `json:"order_sn"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	CreateTime int64   `json:"create_time"`
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
	TransactionID   string  `json:"transaction_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	CreateTime      int64   `json:"create_time"`
}

type GetWalletTransactionListResponse struct {
	BaseResponse
	Response struct {
		TransactionList []WalletTransaction `json:"transaction_list"`
		More            bool                `json:"more"`
		NextCursor      string              `json:"next_cursor"`
	} `json:"response"`
}

func (s *PaymentService) GetWalletTransactionList(pageSize int, cursor string, timeFrom, timeTo int64) (*GetWalletTransactionListResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if timeFrom > 0 {
		q["time_from"] = strconv.FormatInt(timeFrom, 10)
	}
	if timeTo > 0 {
		q["time_to"] = strconv.FormatInt(timeTo, 10)
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
	OrderSN    string  `json:"order_sn"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	CreateTime int64   `json:"create_time"`
}

type GetEscrowListResponse struct {
	BaseResponse
	Response struct {
		EscrowList  []EscrowListItem `json:"escrow_list"`
		More        bool             `json:"more"`
		NextCursor  string           `json:"next_cursor"`
	} `json:"response"`
}

func (s *PaymentService) GetEscrowList(pageSize int, cursor string, timeFrom, timeTo int64) (*GetEscrowListResponse, error) {
	q := map[string]string{"page_size": strconv.Itoa(pageSize)}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if timeFrom > 0 {
		q["time_from"] = strconv.FormatInt(timeFrom, 10)
	}
	if timeTo > 0 {
		q["time_to"] = strconv.FormatInt(timeTo, 10)
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

func (s *PaymentService) GetEscrowDetailBatch(orderSNs []string) (*BaseResponse, error) {
	payload := map[string]any{"order_sn_list": orderSNs}
	result := &BaseResponse{}
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

type IncomeOverview struct {
	TotalIncome     float64 `json:"total_income"`
	PendingIncome   float64 `json:"pending_income"`
	CompletedIncome float64 `json:"completed_income"`
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

type IncomeDetail struct {
	TransactionID string  `json:"transaction_id"`
	OrderSN       string  `json:"order_sn"`
	Amount        float64 `json:"amount"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	CreateTime    int64   `json:"create_time"`
}

type GetIncomeDetailResponse struct {
	BaseResponse
	Response struct {
		IncomeList  []IncomeDetail `json:"income_list"`
		More        bool           `json:"more"`
		NextCursor  string         `json:"next_cursor"`
	} `json:"response"`
}

func (s *PaymentService) GetIncomeDetail(dateFrom, dateTo int64, pageSize int, cursor string) (*GetIncomeDetailResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_size": strconv.Itoa(pageSize),
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
