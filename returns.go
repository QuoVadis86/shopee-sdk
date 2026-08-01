package shopee

import (
	"context"
	"strconv"
)

type ReturnsService struct {
	client *Client
}

func NewReturnsService(client *Client) *ReturnsService {
	return &ReturnsService{client: client}
}

type ReturnItem struct {
	ReturnID      int64        `json:"return_id"`
	OrderSN       string       `json:"order_sn"`
	ItemName      string       `json:"item_name"`
	ReturnStatus  ReturnStatus `json:"return_status"`
	Amount        float64      `json:"amount"`
	Reason        string       `json:"reason"`
	CreateTime    int64        `json:"create_time"`
	DisputeStatus string       `json:"dispute_status,omitempty"`
}

type GetReturnListResponse struct {
	BaseResponse
	Response struct {
		ReturnList []ReturnItem `json:"return_list"`
		More       bool         `json:"more"`
		NextCursor string       `json:"next_cursor"`
	} `json:"response"`
}

func (s *ReturnsService) GetReturnList(pageSize, pageNo int, createTimeFrom, createTimeTo int64, status string) (*GetReturnListResponse, error) {
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
	if status != "" {
		q["status"] = status
	}
	result := &GetReturnListResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnGetList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ReturnDetail struct {
	ReturnID      int64        `json:"return_id"`
	OrderSN       string       `json:"order_sn"`
	ItemName      string       `json:"item_name"`
	ReturnStatus  ReturnStatus `json:"return_status"`
	Amount        float64      `json:"amount"`
	Reason        string       `json:"reason"`
	DisputeStatus string       `json:"dispute_status,omitempty"`
	ItemList      []struct {
		ItemID    int64  `json:"item_id"`
		ModelID   int64  `json:"model_id"`
		Quantity  int    `json:"quantity"`
	} `json:"item_list,omitempty"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type GetReturnDetailResponse struct {
	BaseResponse
	Response *ReturnDetail `json:"response,omitempty"`
}

func (s *ReturnsService) GetReturnDetail(returnSN string) (*GetReturnDetailResponse, error) {
	q := map[string]string{"return_sn": returnSN}
	result := &GetReturnDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnGetDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AvailableSolution struct {
	SolutionID   int64  `json:"solution_id"`
	SolutionType string `json:"solution_type"`
	Description  string `json:"description"`
}

type GetAvailableSolutionsResponse struct {
	BaseResponse
	Response struct {
		SolutionList []AvailableSolution `json:"solution_list"`
	} `json:"response"`
}

func (s *ReturnsService) GetAvailableSolutions(returnSN string) (*GetAvailableSolutionsResponse, error) {
	q := map[string]string{"return_sn": returnSN}
	result := &GetAvailableSolutionsResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnGetAvailableSolutions, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ReturnsService) QueryProof(returnSN string) (*BaseResponse, error) {
	q := map[string]string{"return_sn": returnSN}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnQueryProof, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type DisputeReason struct {
	ReasonID   int64  `json:"reason_id"`
	ReasonName string `json:"reason_name"`
}

type GetReturnDisputeReasonResponse struct {
	BaseResponse
	Response struct {
		ReasonList []DisputeReason `json:"reason_list"`
	} `json:"response"`
}

func (s *ReturnsService) GetReturnDisputeReason(returnSN string) (*GetReturnDisputeReasonResponse, error) {
	q := map[string]string{"return_sn": returnSN}
	result := &GetReturnDisputeReasonResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnGetDisputeReason, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ReturnsService) CancelDispute(returnSN string) (*BaseResponse, error) {
	payload := map[string]any{"return_sn": returnSN}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathReturnCancelDispute, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShippingCarrier struct {
	CarrierID   int64  `json:"carrier_id"`
	CarrierName string `json:"carrier_name"`
}

type GetShippingCarrierResponse struct {
	BaseResponse
	Response struct {
		CarrierList []ShippingCarrier `json:"carrier_list"`
	} `json:"response"`
}

func (s *ReturnsService) GetShippingCarrier(returnSN string) (*GetShippingCarrierResponse, error) {
	q := map[string]string{"return_sn": returnSN}
	result := &GetShippingCarrierResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnGetShippingCarrier, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ReverseTrackingInfo struct {
	TrackingNumber string          `json:"tracking_number"`
	Carrier        string          `json:"carrier"`
	Status         string          `json:"status"`
	Events         []TrackingEvent `json:"events,omitempty"`
}

type GetReverseTrackingInfoResponse struct {
	BaseResponse
	Response *ReverseTrackingInfo `json:"response,omitempty"`
}

func (s *ReturnsService) GetReverseTrackingInfo(returnSN string) (*GetReverseTrackingInfoResponse, error) {
	q := map[string]string{"return_sn": returnSN}
	result := &GetReverseTrackingInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathReturnGetReverseTracking, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
