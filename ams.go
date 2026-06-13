package shopee

import (
	"context"
	"strconv"
)

type AMSService struct {
	client *Client
}

func NewAMSService(client *Client) *AMSService {
	return &AMSService{client: client}
}

type AMSItem struct {
	ItemID                   int64                    `json:"item_id"`
	ItemName                 string                   `json:"item_name"`
	CampaignID               int64                    `json:"campaign_id"`
	CampaignStatus           string                   `json:"campaign_status"`
	CommissionRate           float64                  `json:"commission_rate"`
	PeriodStartTime          int64                    `json:"period_start_time"`
	PeriodEndTime            int64                    `json:"period_end_time"`
	PendingTerminatedTime    int64                    `json:"pending_terminated_time"`
	CommissionProtectionList []AMSCommissionProtection `json:"commission_protection_list"`
	MaxCommissionRateCurrentDay float64               `json:"max_commission_rate_current_day"`
}

type AMSCommissionProtection struct {
	CommissionRate          float64 `json:"commission_rate"`
	ProtectionPeriodEndTime int64   `json:"protection_period_end_time"`
}

type AMSCursorResponse struct {
	BaseResponse
	Response *struct {
		ItemList   []AMSItem `json:"item_list"`
		TotalCount int       `json:"total_count"`
		Cursor     string    `json:"cursor"`
		HasMore    bool      `json:"has_more"`
	} `json:"response,omitempty"`
}

// AMSPaginatedResponse is used by AMS endpoints that use page_num/page_size pagination.
type AMSPaginatedResponse struct {
	BaseResponse
	Response *struct {
		ItemList []AMSItem `json:"item_list"`
		Total    int       `json:"total"`
		PageNum  int       `json:"page_num"`
		PageSize int       `json:"page_size"`
	} `json:"response,omitempty"`
}

// GetOpenCampaignAddedProduct retrieves all products currently in the Open Campaign.
// Official params: page_size, cursor, sort_by, search_type, search_content.
func (s *AMSService) GetOpenCampaignAddedProduct(pageSize int, cursor, sortBy, searchType, searchContent string) (*AMSCursorResponse, error) {
	q := map[string]string{
		"page_size": strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if sortBy != "" {
		q["sort_by"] = sortBy
	}
	if searchType != "" {
		q["search_type"] = searchType
	}
	if searchContent != "" {
		q["search_content"] = searchContent
	}
	result := &AMSCursorResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetOpenCampaignAddedProduct, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

// GetOpenCampaignNotAddedProduct retrieves products not yet in the Open Campaign.
func (s *AMSService) GetOpenCampaignNotAddedProduct(pageSize int, cursor, sortBy, searchType, searchContent string) (*AMSCursorResponse, error) {
	q := map[string]string{
		"page_size": strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if sortBy != "" {
		q["sort_by"] = sortBy
	}
	if searchType != "" {
		q["search_type"] = searchType
	}
	if searchContent != "" {
		q["search_content"] = searchContent
	}
	result := &AMSCursorResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetOpenCampaignNotAddedProduct, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchAddProductsToOpenCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSBatchAddProductsToOpenCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) AddAllProductsToOpenCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSAddAllProductsToOpenCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AutoAddToggleStatus struct {
	AutoAddEnabled bool `json:"auto_add_enabled"`
}

type GetAutoAddToggleStatusResponse struct {
	BaseResponse
	Response *AutoAddToggleStatus `json:"response,omitempty"`
}

func (s *AMSService) GetAutoAddToggleStatus() (*GetAutoAddToggleStatusResponse, error) {
	result := &GetAutoAddToggleStatusResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetAutoAddToggleStatus, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) UpdateAutoAddNewProductSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSUpdateAutoAddNewProductSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchEditProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSBatchEditProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditAllProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSEditAllProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchRemoveProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSBatchRemoveProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) RemoveAllProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSRemoveAllProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BatchTaskResult struct {
	BatchID string `json:"batch_id"`
	Status  string `json:"status"`
	Total   int    `json:"total"`
	Success int    `json:"success"`
	Failed  int    `json:"failed"`
}

type GetBatchTaskResultResponse struct {
	BaseResponse
	Response *BatchTaskResult `json:"response,omitempty"`
}

func (s *AMSService) GetOpenCampaignBatchTaskResult(batchID string) (*GetBatchTaskResultResponse, error) {
	q := map[string]string{"batch_id": batchID}
	result := &GetBatchTaskResultResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetOpenCampaignBatchTaskResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetOptimizationSuggestionProduct(pageNum, pageSize int) (*AMSPaginatedResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &AMSPaginatedResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetOptimizationSuggestionProduct, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchGetProductsSuggestedRate(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSBatchGetProductsSuggestedRate, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SuggestedRateInfo struct {
	ShopID        int64   `json:"shop_id"`
	SuggestedRate float64 `json:"suggested_rate"`
}

type GetShopSuggestedRateResponse struct {
	BaseResponse
	Response *SuggestedRateInfo `json:"response,omitempty"`
}

func (s *AMSService) GetShopSuggestedRate() (*GetShopSuggestedRateResponse, error) {
	result := &GetShopSuggestedRateResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetShopSuggestedRate, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetTargetedCampaignAddableProductList(campaignID string, pageNum, pageSize int) (*AMSPaginatedResponse, error) {
	q := map[string]string{
		"campaign_id": campaignID,
		"page_num":    strconv.Itoa(pageNum),
		"page_size":   strconv.Itoa(pageSize),
	}
	result := &AMSPaginatedResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetTargetedCampaignAddableProducts, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetRecommendedAffiliateList(campaignID string) (*BaseResponse, error) {
	q := map[string]string{"campaign_id": campaignID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetRecommendedAffiliateList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetManagedAffiliateList(pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetManagedAffiliateList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) QueryAffiliateList(keyword string, pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	if keyword != "" {
		q["keyword"] = keyword
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSQueryAffiliateList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) CreateNewTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSCreateTargetedCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetTargetedCampaignList(pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetTargetedCampaignList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetTargetedCampaignSettings(campaignID string) (*BaseResponse, error) {
	q := map[string]string{"campaign_id": campaignID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetTargetedCampaignSettings, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) UpdateBasicInfoOfTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSUpdateTargetedCampaignBasicInfo, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditProductListOfTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSEditTargetedCampaignProductList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditAffiliateListOfTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSEditTargetedCampaignAffiliateList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PerfDataUpdateTime struct {
	UpdateTime int64 `json:"update_time"`
}

type GetPerfDataUpdateTimeResponse struct {
	BaseResponse
	Response *PerfDataUpdateTime `json:"response,omitempty"`
}

func (s *AMSService) GetPerformanceDataUpdateTime() (*GetPerfDataUpdateTimeResponse, error) {
	result := &GetPerfDataUpdateTimeResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetPerfDataUpdateTime, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AMSDateRangeParams struct {
	DateFrom int64 `json:"date_from"`
	DateTo   int64 `json:"date_to"`
}

type AMSShopPerformanceData struct {
	Impression  int64   `json:"impression"`
	Click       int64   `json:"click"`
	Spend       float64 `json:"spend"`
	Sales       float64 `json:"sales"`
	ROI         float64 `json:"roi"`
	CTR         float64 `json:"ctr"`
	CPC         float64 `json:"cpc"`
	Conversion  float64 `json:"conversion"`
}

type GetAMSShopPerformanceResponse struct {
	BaseResponse
	Response *AMSShopPerformanceData `json:"response,omitempty"`
}

func (s *AMSService) GetShopPerformance(dateFrom, dateTo int64) (*GetAMSShopPerformanceResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &GetAMSShopPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetShopPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ProductPerformance struct {
	ItemID      int64   `json:"item_id"`
	ItemName    string  `json:"item_name"`
	Impression  int64   `json:"impression"`
	Click       int64   `json:"click"`
	Spend       float64 `json:"spend"`
	Sales       float64 `json:"sales"`
	ROI         float64 `json:"roi"`
	Conversion  float64 `json:"conversion"`
}

type GetProductPerformanceResponse struct {
	BaseResponse
	Response *struct {
		PerformanceList []ProductPerformance `json:"performance_list"`
		Total           int                  `json:"total"`
		PageNum         int                  `json:"page_num"`
		PageSize        int                  `json:"page_size"`
	} `json:"response,omitempty"`
}

func (s *AMSService) GetProductPerformance(dateFrom, dateTo int64, pageNum, pageSize int) (*GetProductPerformanceResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &GetProductPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetProductPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetAffiliatePerformance(dateFrom, dateTo int64, pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetAffiliatePerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetContentPerformance(dateFrom, dateTo int64, pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetContentPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetCampaignKeyMetricsPerformance(campaignID string, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"campaign_id": campaignID,
		"date_from":   strconv.FormatInt(dateFrom, 10),
		"date_to":     strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetCampaignKeyMetricsPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetOpenCampaignPerformance(dateFrom, dateTo int64) (*GetAMSShopPerformanceResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &GetAMSShopPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetOpenCampaignPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetTargetedCampaignPerformance(campaignID string, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"campaign_id": campaignID,
		"date_from":   strconv.FormatInt(dateFrom, 10),
		"date_to":     strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetTargetedCampaignPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetConversionReport(dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetConversionReport, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetValidationList(pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetValidationList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetValidationReport(validationID string) (*BaseResponse, error) {
	q := map[string]string{"validation_id": validationID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetValidationReport, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetCoverList(pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetCoverList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditVideoInfo(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSEditVideoInfo, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoList(pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoDetail(videoID string) (*BaseResponse, error) {
	q := map[string]string{"video_id": videoID}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) DeleteVideo(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAMSDeleteVideo, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetOverviewPerformance(dateFrom, dateTo int64) (*GetAMSShopPerformanceResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &GetAMSShopPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetOverviewPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetMetricTrend(metric string, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"metric":    metric,
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetMetricTrend, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetUserDemographics(dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetUserDemographics, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoPerformanceList(dateFrom, dateTo int64, pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoPerfList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetProductPerfList(dateFrom, dateTo int64, pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetProductPerfList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoDetailPerformance(videoID string, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"video_id":  videoID,
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoDetailPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoDetailMetricTrend(videoID, metric string, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"video_id":  videoID,
		"metric":    metric,
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoDetailMetricTrend, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoDetailAudienceDistribution(videoID string, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"video_id":  videoID,
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoDetailAudienceDist, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetVideoDetailProductPerformance(videoID string, dateFrom, dateTo int64, pageNum, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"video_id":  videoID,
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAMSGetVideoDetailProductPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
