package shopee

import (
	"context"
	"strconv"
)

type AdsService struct {
	client *Client
}

func NewAdsService(client *Client) *AdsService {
	return &AdsService{client: client}
}

type RecommendedKeyword struct {
	Keyword string  `json:"keyword"`
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
	if err := s.client.DoGet(context.Background(), PathAdsGetRecommendedKeywords, q, result); err != nil {
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
	if err := s.client.DoGet(context.Background(), PathAdsGetRecommendedItems, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CPCAdsDailyPerformance struct {
	Date              string  `json:"date"`
	Impression        int64   `json:"impression"`
	Clicks            int64   `json:"clicks"`
	CTR               float64 `json:"ctr"`
	DirectOrder       int64   `json:"direct_order"`
	BroadOrder        int64   `json:"broad_order"`
	DirectConversions float64 `json:"direct_conversions"`
	BroadConversions  float64 `json:"broad_conversions"`
	DirectItemSold    int64   `json:"direct_item_sold"`
	BroadItemSold     int64   `json:"broad_item_sold"`
	DirectGMV         float64 `json:"direct_gmv"`
	BroadGMV          float64 `json:"broad_gmv"`
	Expense           float64 `json:"expense"`
	CostPerConversion float64 `json:"cost_per_conversion"`
	DirectROAS        float64 `json:"direct_roas"`
	BroadROAS         float64 `json:"broad_roas"`
}

type GetAllCPCAdsDailyPerformanceResponse struct {
	BaseResponse
	Warning  string                   `json:"warning,omitempty"`
	Response []CPCAdsDailyPerformance `json:"response"`
}

func (s *AdsService) GetAllCPCHourlyPerformance(date int64) (*BaseResponse, error) {
	q := map[string]string{"date": strconv.FormatInt(date, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetAllCPCHourlyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetAllCPCDailyPerformance(
	startDate, endDate string,
) (*GetAllCPCAdsDailyPerformanceResponse, error) {
	q := map[string]string{
		"start_date": startDate,
		"end_date":   endDate,
	}
	result := &GetAllCPCAdsDailyPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetAllCPCDailyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ProductCampaignDailyMetric struct {
	Date              string  `json:"date"`
	Impression        int64   `json:"impression"`
	Clicks            int64   `json:"clicks"`
	CTR               float64 `json:"ctr"`
	Expense           float64 `json:"expense"`
	BroadGMV          float64 `json:"broad_gmv"`
	BroadOrder        int64   `json:"broad_order"`
	BroadOrderAmount  int64   `json:"broad_order_amount"`
	BroadROI          float64 `json:"broad_roi"`
	BroadCIR          float64 `json:"broad_cir"`
	CR                float64 `json:"cr"`
	CPC               float64 `json:"cpc"`
	DirectOrder       int64   `json:"direct_order"`
	DirectOrderAmount int64   `json:"direct_order_amount"`
	DirectGMV         float64 `json:"direct_gmv"`
	DirectROI         float64 `json:"direct_roi"`
	DirectCIR         float64 `json:"direct_cir"`
	DirectCR          float64 `json:"direct_cr"`
	CPDC              float64 `json:"cpdc"`
}

type ProductCampaignDaily struct {
	CampaignID        int64                        `json:"campaign_id"`
	AdType            string                       `json:"ad_type"`
	CampaignPlacement string                       `json:"campaign_placement"`
	AdName            string                       `json:"ad_name"`
	MetricsList       []ProductCampaignDailyMetric `json:"metrics_list"`
}

type ProductCampaignDailyShop struct {
	ShopID       int64                  `json:"shop_id"`
	Region       string                 `json:"region"`
	CampaignList []ProductCampaignDaily `json:"campaign_list"`
}

type GetProductCampaignDailyPerformanceResponse struct {
	BaseResponse
	Warning  string                     `json:"warning,omitempty"`
	Response []ProductCampaignDailyShop `json:"response"`
}

func (s *AdsService) GetProductCampaignDailyPerformance(
	campaignIDs []int64,
	startDate, endDate string,
) (*GetProductCampaignDailyPerformanceResponse, error) {
	ids := make([]string, len(campaignIDs))
	for i, id := range campaignIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{
		"campaign_id_list": stringsJoin(ids, ","),
		"start_date":       startDate,
		"end_date":         endDate,
	}
	result := &GetProductCampaignDailyPerformanceResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetProductCampaignDailyPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetProductCampaignHourlyPerformance(campaignIDs []int64, performanceDate string) (*BaseResponse, error) {
	ids := make([]string, len(campaignIDs))
	for i, id := range campaignIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{
		"campaign_id_list": stringsJoin(ids, ","),
		"performance_date": performanceDate,
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetProductCampaignHourlyPerf, q, result); err != nil {
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
	if err := s.client.DoGet(context.Background(), PathAdsGetProductLevelCampaignIDs, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CampaignSettingInfo struct {
	CampaignID   int64   `json:"campaign_id"`
	CampaignName string  `json:"campaign_name"`
	Budget       float64 `json:"budget"`
	BudgetType   string  `json:"budget_type"`
	Status       string  `json:"status"`
}

type GetProductLevelCampaignSettingResponse struct {
	BaseResponse
	Response *CampaignSettingInfo `json:"response,omitempty"`
}

func (s *AdsService) GetProductLevelCampaignSetting(campaignIDs []int64, infoTypeList string) (*GetProductLevelCampaignSettingResponse, error) {
	ids := make([]string, len(campaignIDs))
	for i, id := range campaignIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{
		"campaign_id_list": stringsJoin(ids, ","),
		"info_type_list":   infoTypeList,
	}
	result := &GetProductLevelCampaignSettingResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetProductLevelCampaignSetting, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) CreateManualProductAds(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAdsCreateManualProductAds, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditManualProductAdKeywords(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAdsEditManualProductAdKeywords, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditManualProductAds(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAdsEditManualProductAds, params, result); err != nil {
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

func (s *AdsService) GetCreateAdBudgetSuggestion(campaignID int64, biddingMethod, campaignPlacement string) (*GetCreateAdBudgetSuggestionResponse, error) {
	q := map[string]string{
		"campaign_id":        strconv.FormatInt(campaignID, 10),
		"bidding_method":     biddingMethod,
		"campaign_placement": campaignPlacement,
	}
	result := &GetCreateAdBudgetSuggestionResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetCreateAdBudgetSuggestion, q, result); err != nil {
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

func (s *AdsService) GetRecommendedROITarget(itemID, referenceID int64) (*GetRecommendedROITargetResponse, error) {
	q := map[string]string{
		"item_id":      strconv.FormatInt(itemID, 10),
		"reference_id": strconv.FormatInt(referenceID, 10),
	}
	result := &GetRecommendedROITargetResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetRecommendedROITarget, q, result); err != nil {
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
	if err := s.client.DoGet(context.Background(), PathAdsGetFacilShopRate, map[string]string{}, result); err != nil {
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
	if err := s.client.DoGet(context.Background(), PathAdsCheckCreateGMSCampaignElig, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) CreateGMSCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAdsCreateGMSCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditGMSCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAdsEditGMSCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) EditGMSItemCampaign(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathAdsEditGMSItemCampaign, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *AdsService) GetGMSCampaignPerformance(campaignID int64, dateFrom, dateTo int64) (*BaseResponse, error) {
	q := map[string]string{
		"campaign_id": strconv.FormatInt(campaignID, 10),
		"date_from":   strconv.FormatInt(dateFrom, 10),
		"date_to":     strconv.FormatInt(dateTo, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetGMSCampaignPerf, q, result); err != nil {
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
		"page_no": strconv.Itoa(pageNumber),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathAdsGetGMSItemPerf, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
