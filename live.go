package shopee

import (
	"context"
	"strconv"
)

type LiveService struct {
	client *Client
}

func NewLiveService(client *Client) *LiveService {
	return &LiveService{client: client}
}

type SessionInfo struct {
	SessionID int64  `json:"session_id"`
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Status    string `json:"status"`
}

type CreateSessionResponse struct {
	BaseResponse
	Response struct {
		SessionID int64 `json:"session_id"`
	} `json:"response,omitempty"`
}

func (s *LiveService) CreateSession(params any) (*CreateSessionResponse, error) {
	result := &CreateSessionResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveCreateSession, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) UpdateSession(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveUpdateSession, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GetSessionDetailResponse struct {
	BaseResponse
	Response *SessionInfo `json:"response,omitempty"`
}

func (s *LiveService) GetSessionDetail(sessionID int64) (*GetSessionDetailResponse, error) {
	q := map[string]string{"session_id": strconv.FormatInt(sessionID, 10)}
	result := &GetSessionDetailResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetSessionDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) AddItemList(sessionID int64, itemIDs []int64) (*BaseResponse, error) {
	payload := map[string]any{
		"session_id":  sessionID,
		"item_id_list": itemIDs,
	}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveAddItemList, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) DeleteItemList(sessionID int64, itemIDs []int64) (*BaseResponse, error) {
	payload := map[string]any{
		"session_id":   sessionID,
		"item_id_list": itemIDs,
	}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveDeleteItemList, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) UpdateItemList(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveUpdateItemList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type LiveItemCount struct {
	TotalItems int `json:"total_items"`
}

type GetLiveItemCountResponse struct {
	BaseResponse
	Response *LiveItemCount `json:"response,omitempty"`
}

func (s *LiveService) GetItemCount(sessionID int64) (*GetLiveItemCountResponse, error) {
	q := map[string]string{"session_id": strconv.FormatInt(sessionID, 10)}
	result := &GetLiveItemCountResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetItemCount, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type LiveItem struct {
	ItemID    int64  `json:"item_id"`
	ItemName  string `json:"item_name,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

type GetLiveItemListResponse struct {
	BaseResponse
	Response struct {
		ItemList   []LiveItem `json:"item_list"`
		Total      int        `json:"total"`
		PageNum    int        `json:"page_num"`
		PageSize   int        `json:"page_size"`
	} `json:"response"`
}

func (s *LiveService) GetItemList(sessionID int64, pageSize, pageNumber int) (*GetLiveItemListResponse, error) {
	q := map[string]string{
		"session_id":  strconv.FormatInt(sessionID, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetLiveItemListResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) UpdateShowItem(sessionID int64, itemID int64) (*BaseResponse, error) {
	payload := map[string]any{
		"session_id": sessionID,
		"item_id":    itemID,
	}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveUpdateShowItem, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) DeleteShowItem(sessionID int64, itemID int64) (*BaseResponse, error) {
	payload := map[string]any{
		"session_id": sessionID,
		"item_id":    itemID,
	}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathLiveDeleteShowItem, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShowItemInfo struct {
	SessionID int64 `json:"session_id"`
	ItemID    int64 `json:"item_id"`
}

type GetShowItemResponse struct {
	BaseResponse
	Response *ShowItemInfo `json:"response,omitempty"`
}

func (s *LiveService) GetShowItem(sessionID int64) (*GetShowItemResponse, error) {
	q := map[string]string{"session_id": strconv.FormatInt(sessionID, 10)}
	result := &GetShowItemResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetShowItem, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) GetLikeItemList(sessionID int64, pageSize, pageNumber int) (*BaseResponse, error) {
	q := map[string]string{
		"session_id":  strconv.FormatInt(sessionID, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetLikeItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) GetRecentItemList(sessionID int64, pageSize, pageNumber int) (*GetLiveItemListResponse, error) {
	q := map[string]string{
		"session_id":  strconv.FormatInt(sessionID, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetLiveItemListResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetRecentItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ItemSet struct {
	ItemSetID   int64  `json:"item_set_id"`
	ItemSetName string `json:"item_set_name"`
}

type GetItemSetListResponse struct {
	BaseResponse
	Response struct {
		ItemSetList []ItemSet `json:"item_set_list"`
		Total       int       `json:"total"`
		PageNum     int       `json:"page_num"`
		PageSize    int       `json:"page_size"`
	} `json:"response"`
}

func (s *LiveService) GetItemSetList(pageSize, pageNumber int) (*GetItemSetListResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetItemSetListResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetItemSetList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *LiveService) GetItemSetItemList(itemSetID int64, pageSize, pageNumber int) (*GetLiveItemListResponse, error) {
	q := map[string]string{
		"item_set_id": strconv.FormatInt(itemSetID, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetLiveItemListResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetItemSetItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SessionMetric struct {
	SessionID    int64 `json:"session_id"`
	TotalViews   int64 `json:"total_views"`
	UniqueViewers int64 `json:"unique_viewers"`
	TotalLikes   int64 `json:"total_likes"`
	TotalShares  int64 `json:"total_shares"`
	TotalComments int64 `json:"total_comments"`
	TotalOrders  int64 `json:"total_orders"`
	TotalSales   float64 `json:"total_sales"`
}

type GetSessionMetricResponse struct {
	BaseResponse
	Response *SessionMetric `json:"response,omitempty"`
}

func (s *LiveService) GetSessionMetric(sessionID int64) (*GetSessionMetricResponse, error) {
	q := map[string]string{"session_id": strconv.FormatInt(sessionID, 10)}
	result := &GetSessionMetricResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetSessionMetric, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SessionItemMetric struct {
	ItemID      int64 `json:"item_id"`
	Views       int64 `json:"views"`
	Likes       int64 `json:"likes"`
	Orders      int64 `json:"orders"`
	Sales       float64 `json:"sales"`
}

type GetSessionItemMetricResponse struct {
	BaseResponse
	Response *SessionItemMetric `json:"response,omitempty"`
}

func (s *LiveService) GetSessionItemMetric(sessionID, itemID int64) (*GetSessionItemMetricResponse, error) {
	q := map[string]string{
		"session_id": strconv.FormatInt(sessionID, 10),
		"item_id":    strconv.FormatInt(itemID, 10),
	}
	result := &GetSessionItemMetricResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetSessionItemMetric, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type LiveCommentInfo struct {
	CommentID   int64  `json:"comment_id"`
	Content     string `json:"content"`
	BuyerName   string `json:"buyer_username,omitempty"`
	CreateTime  int64  `json:"create_time"`
}

type GetLatestCommentListResponse struct {
	BaseResponse
	Response struct {
		CommentList []LiveCommentInfo `json:"comment_list"`
		Total       int               `json:"total"`
	} `json:"response"`
}

func (s *LiveService) GetLatestCommentList(sessionID int64, pageSize, pageNumber int) (*GetLatestCommentListResponse, error) {
	q := map[string]string{
		"session_id":  strconv.FormatInt(sessionID, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &GetLatestCommentListResponse{}
	if err := s.client.DoGet(context.Background(), PathLiveGetLatestCommentList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
