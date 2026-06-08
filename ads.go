package shopee

import "strconv"

type AdsService struct {
	client *Client
}

func NewAdsService(client *Client) *AdsService {
	return &AdsService{client: client}
}

type RecommendedKeyword struct {
	Keyword string `json:"keyword"`
	Score   float64 `json:"score,omitempty"`
}

type GetRecommendedKeywordsResponse struct {
	BaseResponse
	Response struct {
		KeywordList []RecommendedKeyword `json:"keyword_list"`
	} `json:"response"`
}

func (s *AdsService) GetRecommendedKeywords(itemID int64, limit int) (*GetRecommendedKeywordsResponse, error) {
	q := map[string]string{
		"item_id": strconv.FormatInt(itemID, 10),
		"limit":   strconv.Itoa(limit),
	}
	result := &GetRecommendedKeywordsResponse{}
	if err := s.client.DoGet(PathAdsGetRecommendedKeywords, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type RecommendedItem struct {
	ItemID   int64  `json:"item_id"`
	ItemName string `json:"item_name"`
}

type GetRecommendedItemsResponse struct {
	BaseResponse
	Response struct {
		ItemList []RecommendedItem `json:"item_list"`
	} `json:"response"`
}

func (s *AdsService) GetRecommendedItems(keyword string, limit int) (*GetRecommendedItemsResponse, error) {
	q := map[string]string{
		"keyword": keyword,
		"limit":   strconv.Itoa(limit),
	}
	result := &GetRecommendedItemsResponse{}
	if err := s.client.DoGet(PathAdsGetRecommendedItems, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CPCAdsPerformance struct {
	Impression int64   `json:"impression"`
	Click      int64   `json:"click"`
	Spend      float64 `json:"spend"`
	Sales      float64 `json:"sales"`
	ROI        float64 `json:"roi"`
}

type GetAllCPCAdsPerformanceResponse struct {
	BaseResponse
	Response *CPCAdsPerformance `json:"response,omitempty"`
}

func (s *AdsService) GetAllCPCHourlyPerformance(date int64) (*GetAllCPCAdsPerformanceResponse, error) {
	q := map[string]string{"date": strconv.FormatInt(date, 10)}
	result := &GetAllCPCAdsPerformanceResponse{}
	if err := s.client.DoGet(PathAdsGetAllCPCHourlyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetAllCPCDailyPerformance(dateFrom, dateTo int64) (*GetAllCPCAdsPerformanceResponse, error) {
	q := map[string]string{
		"date_from": strconv.FormatInt(dateFrom, 10),
		"date_to":   strconv.FormatInt(dateTo, 10),
	}
	result := &GetAllCPCAdsPerformanceResponse{}
	if err := s.client.DoGet(PathAdsGetAllCPCDailyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetProductCampaignDailyPerformance(campaignID int64, dateFrom, dateTo int64) (*GetAllCPCAdsPerformanceResponse, error) {
	q := map[string]string{
		"campaign_id": strconv.FormatInt(campaignID, 10),
		"date_from":   strconv.FormatInt(dateFrom, 10),
		"date_to":     strconv.FormatInt(dateTo, 10),
	}
	result := &GetAllCPCAdsPerformanceResponse{}
	if err := s.client.DoGet(PathAdsGetProductCampaignDailyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetProductCampaignHourlyPerformance(campaignID int64, date int64) (*GetAllCPCAdsPerformanceResponse, error) {
	q := map[string]string{
		"campaign_id": strconv.FormatInt(campaignID, 10),
		"date":        strconv.FormatInt(date, 10),
	}
	result := &GetAllCPCAdsPerformanceResponse{}
	if err := s.client.DoGet(PathAdsGetProductCampaignHourlyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CampaignIDInfo struct {
	ItemID     int64 `json:"item_id"`
	CampaignID int64 `json:"campaign_id"`
}

type GetProductLevelCampaignIDsResponse struct {
	BaseResponse
	Response struct {
		CampaignList []CampaignIDInfo `json:"campaign_list"`
	} `json:"response"`
}

func (s *AdsService) GetProductLevelCampaignIDs(itemIDs []int64) (*GetProductLevelCampaignIDsResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetProductLevelCampaignIDsResponse{}
	if err := s.client.DoGet(PathAdsGetProductLevelCampaignIDs, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CampaignSettingInfo struct {
	CampaignID     int64   `json:"campaign_id"`
	CampaignName   string  `json:"campaign_name"`
	Budget         float64 `json:"budget"`
	BudgetType     string  `json:"budget_type"`
	Status         string  `json:"status"`
}

type GetProductLevelCampaignSettingResponse struct {
	BaseResponse
	Response *CampaignSettingInfo `json:"response,omitempty"`
}

func (s *AdsService) GetProductLevelCampaignSetting(campaignID int64) (*GetProductLevelCampaignSettingResponse, error) {
	q := map[string]string{"campaign_id": strconv.FormatInt(campaignID, 10)}
	result := &GetProductLevelCampaignSettingResponse{}
	if err := s.client.DoGet(PathAdsGetProductLevelCampaignSetting, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) CreateManualProductAds(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAdsCreateManualProductAds, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditManualProductAdKeywords(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAdsEditManualProductAdKeywords, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditManualProductAds(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAdsEditManualProductAds, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BudgetSuggestion struct {
	SuggestedBudget float64 `json:"suggested_budget"`
	MinBudget       float64 `json:"min_budget"`
	MaxBudget       float64 `json:"max_budget"`
}

type GetCreateAdBudgetSuggestionResponse struct {
	BaseResponse
	Response *BudgetSuggestion `json:"response,omitempty"`
}

func (s *AdsService) GetCreateAdBudgetSuggestion(campaignID int64) (*GetCreateAdBudgetSuggestionResponse, error) {
	q := map[string]string{"campaign_id": strconv.FormatInt(campaignID, 10)}
	result := &GetCreateAdBudgetSuggestionResponse{}
	if err := s.client.DoGet(PathAdsGetCreateAdBudgetSuggestion, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ROITargetInfo struct {
	RecommendedROI float64 `json:"recommended_roi"`
	MinROI         float64 `json:"min_roi"`
	MaxROI         float64 `json:"max_roi"`
}

type GetRecommendedROITargetResponse struct {
	BaseResponse
	Response *ROITargetInfo `json:"response,omitempty"`
}

func (s *AdsService) GetRecommendedROITarget(campaignID int64) (*GetRecommendedROITargetResponse, error) {
	q := map[string]string{"campaign_id": strconv.FormatInt(campaignID, 10)}
	result := &GetRecommendedROITargetResponse{}
	if err := s.client.DoGet(PathAdsGetRecommendedROITarget, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type FacilShopRateInfo struct {
	ShopID int64   `json:"shop_id"`
	Rate   float64 `json:"rate"`
}

type GetFacilShopRateResponse struct {
	BaseResponse
	Response *FacilShopRateInfo `json:"response,omitempty"`
}

func (s *AdsService) GetFacilShopRate() (*GetFacilShopRateResponse, error) {
	result := &GetFacilShopRateResponse{}
	if err := s.client.DoGet(PathAdsGetFacilShopRate, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GMSEligibilityInfo struct {
	Eligible bool   `json:"eligible"`
	Reason   string `json:"reason,omitempty"`
}

type CheckGMSCampaignEligibilityResponse struct {
	BaseResponse
	Response *GMSEligibilityInfo `json:"response,omitempty"`
}

func (s *AdsService) CheckCreateGMSCampaignEligibility() (*CheckGMSCampaignEligibilityResponse, error) {
	result := &CheckGMSCampaignEligibilityResponse{}
	if err := s.client.DoGet(PathAdsCheckCreateGMSCampaignElig, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) CreateGMSCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAdsCreateGMSCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditGMSCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAdsEditGMSCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditGMSItemCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(PathAdsEditGMSItemCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetGMSCampaignPerformance(campaignID int64, dateFrom, dateTo int64) (*GetAllCPCAdsPerformanceResponse, error) {
	q := map[string]string{
		"campaign_id": strconv.FormatInt(campaignID, 10),
		"date_from":   strconv.FormatInt(dateFrom, 10),
		"date_to":     strconv.FormatInt(dateTo, 10),
	}
	result := &GetAllCPCAdsPerformanceResponse{}
	if err := s.client.DoGet(PathAdsGetGMSCampaignPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetGMSItemPerformance(campaignID int64, dateFrom, dateTo int64, pageSize, pageNumber int) (*BaseResponse, error) {
	q := map[string]string{
		"campaign_id": strconv.FormatInt(campaignID, 10),
		"date_from":   strconv.FormatInt(dateFrom, 10),
		"date_to":     strconv.FormatInt(dateTo, 10),
		"page_size":   strconv.Itoa(pageSize),
		"page_number": strconv.Itoa(pageNumber),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(PathAdsGetGMSItemPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
