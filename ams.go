package shopee

import "strconv"

type AMSService struct {
	client *Client
}

func NewAMSService(client *Client) *AMSService {
	return &AMSService{client: client}
}

type AMSItem struct {
	ItemID       int64   `json:"item_id"`
	ItemName     string  `json:"item_name"`
	ItemStatus   string  `json:"item_status"`
	CurrentBid   float64 `json:"current_bid,omitempty"`
	SuggestedBid float64 `json:"suggested_bid,omitempty"`
	ImageURL     string  `json:"image_url,omitempty"`
}

type AMSPaginatedResponse struct {
	BaseResponse
	Response *struct {
		ItemList []AMSItem `json:"item_list"`
		Total    int       `json:"total"`
		PageNum  int       `json:"page_num"`
		PageSize int       `json:"page_size"`
	} `json:"response,omitempty"`
}

type GetOpenCampaignAddedProductRequest struct {
	PageSize      int
	Cursor        string
	SortBy        string
	SearchType    string
	SearchContent string
}

type CommissionProtection struct {
	CommissionRate          float64 `json:"commission_rate"`
	ProtectionPeriodEndTime int64   `json:"protection_period_end_time"`
}

type OpenCampaignAddedProduct struct {
	ItemID                      int64                  `json:"item_id"`
	ItemName                    string                 `json:"item_name"`
	CampaignID                  int64                  `json:"campaign_id"`
	CampaignStatus              string                 `json:"campaign_status"`
	CommissionRate              float64                `json:"commission_rate"`
	PeriodStartTime             int64                  `json:"period_start_time"`
	PeriodEndTime               int64                  `json:"period_end_time"`
	PendingTerminatedTime       int64                  `json:"pending_terminated_time"`
	CommissionProtectionList    []CommissionProtection `json:"commission_protection_list"`
	MaxCommissionRateCurrentDay float64                `json:"max_commission_rate_current_day"`
}

type GetOpenCampaignAddedProductResponse struct {
	BaseResponse
	Response *struct {
		ItemList   []OpenCampaignAddedProduct `json:"item_list"`
		TotalCount int                        `json:"total_count"`
		Cursor     string                     `json:"cursor"`
		HasMore    bool                       `json:"has_more"`
	} `json:"response,omitempty"`
}

func (s *AMSService) GetOpenCampaignAddedProduct(
	request GetOpenCampaignAddedProductRequest,
) (*GetOpenCampaignAddedProductResponse, error) {
	q := map[string]string{
		"page_size": strconv.Itoa(request.PageSize),
	}
	if request.Cursor != "" {
		q["cursor"] = request.Cursor
	}
	if request.SortBy != "" {
		q["sort_by"] = request.SortBy
	}
	if request.SearchType != "" {
		q["search_type"] = request.SearchType
	}
	if request.SearchContent != "" {
		q["search_content"] = request.SearchContent
	}
	result := &GetOpenCampaignAddedProductResponse{}
	if err := s.client.DoGet(PathAMSGetOpenCampaignAddedProduct, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) GetOpenCampaignNotAddedProduct(pageNum, pageSize int, keyword string) (*AMSPaginatedResponse, error) {
	q := map[string]string{
		"page_num":  strconv.Itoa(pageNum),
		"page_size": strconv.Itoa(pageSize),
	}
	if keyword != "" {
		q["keyword"] = keyword
	}
	result := &AMSPaginatedResponse{}
	if err := s.client.DoGet(PathAMSGetOpenCampaignNotAddedProduct, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchAddProductsToOpenCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSBatchAddProductsToOpenCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) AddAllProductsToOpenCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSAddAllProductsToOpenCampaign, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetAutoAddToggleStatus, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) UpdateAutoAddNewProductSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSUpdateAutoAddNewProductSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchEditProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSBatchEditProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditAllProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSEditAllProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchRemoveProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSBatchRemoveProductsOCSetting, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) RemoveAllProductsOCSetting(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSRemoveAllProductsOCSetting, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetOpenCampaignBatchTaskResult, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetOptimizationSuggestionProduct, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) BatchGetProductsSuggestedRate(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSBatchGetProductsSuggestedRate, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetShopSuggestedRate, map[string]string{}, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetTargetedCampaignAddableProducts, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetRecommendedAffiliateList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetManagedAffiliateList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSQueryAffiliateList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) CreateNewTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSCreateTargetedCampaign, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetTargetedCampaignList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetTargetedCampaignSettings, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) UpdateBasicInfoOfTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSUpdateTargetedCampaignBasicInfo, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditProductListOfTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSEditTargetedCampaignProductList, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditAffiliateListOfTargetedCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSEditTargetedCampaignAffiliateList, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetPerfDataUpdateTime, map[string]string{}, result); err != nil {
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
	Impression int64   `json:"impression"`
	Click      int64   `json:"click"`
	Spend      float64 `json:"spend"`
	Sales      float64 `json:"sales"`
	ROI        float64 `json:"roi"`
	CTR        float64 `json:"ctr"`
	CPC        float64 `json:"cpc"`
	Conversion float64 `json:"conversion"`
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
	if err := s.client.DoGet(PathAMSGetShopPerformance, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ProductPerformance struct {
	ItemID     int64   `json:"item_id"`
	ItemName   string  `json:"item_name"`
	Impression int64   `json:"impression"`
	Click      int64   `json:"click"`
	Spend      float64 `json:"spend"`
	Sales      float64 `json:"sales"`
	ROI        float64 `json:"roi"`
	Conversion float64 `json:"conversion"`
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
	if err := s.client.DoGet(PathAMSGetProductPerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetAffiliatePerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetContentPerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetCampaignKeyMetricsPerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetOpenCampaignPerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetTargetedCampaignPerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetConversionReport, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetValidationList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetValidationReport, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetCoverList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) EditVideoInfo(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSEditVideoInfo, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AMSService) DeleteVideo(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAMSDeleteVideo, params, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetOverviewPerformance, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetMetricTrend, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetUserDemographics, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoPerfList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetProductPerfList, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoDetailPerf, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoDetailMetricTrend, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoDetailAudienceDist, q, result); err != nil {
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
	if err := s.client.DoGet(PathAMSGetVideoDetailProductPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
