package shopee

import (
	"context"
	"net/url"
	"strconv"
)

type GlobalProductService struct {
	client *Client
}

func NewGlobalProductService(client *Client) *GlobalProductService {
	return &GlobalProductService{client: client}
}

func (s *GlobalProductService) GetCategory(language string) (*GetCategoryResponse, error) {
	q := map[string]string{}
	if language != "" {
		q["language"] = language
	}
	result := &GetCategoryResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetCategory, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetAttributeTree(categoryID int64, language string) (*GetAttributeTreeResponse, error) {
	q := map[string]string{
		"category_id": strconv.FormatInt(categoryID, 10),
	}
	if language != "" {
		q["language"] = language
	}
	result := &GetAttributeTreeResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetAttributeTree, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetBrandList(categoryID int64, language string, offset, pageSize int) (*GetBrandListResponse, error) {
	q := map[string]string{
		"category_id": strconv.FormatInt(categoryID, 10),
		"offset":      strconv.Itoa(offset),
		"page_size":   strconv.Itoa(pageSize),
	}
	if language != "" {
		q["language"] = language
	}
	result := &GetBrandListResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetBrandList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetItemLimit(itemName string, categoryID int64) (*GetItemLimitResponse, error) {
	q := map[string]string{"item_name": itemName}
	if categoryID > 0 {
		q["category_id"] = strconv.FormatInt(categoryID, 10)
	}
	result := &GetItemLimitResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetItemLimit, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

// GetGlobalItemListResponse is the response for global_product.get_global_item_list.
// Uses item_list (not item) to match the global product API response structure.
type GetGlobalItemListResponse struct {
	BaseResponse
	Response struct {
		ItemList    []ItemListItem `json:"item_list"`
		TotalCount  int            `json:"total_count"`
		HasNextPage bool           `json:"has_next_page"`
		NextOffset  int            `json:"next_offset"`
	} `json:"response"`
}

func (s *GlobalProductService) GetItemList(offset, pageSize int, updateTimeFrom, updateTimeTo int64, itemStatus []string) (*GetGlobalItemListResponse, error) {
	q := url.Values{
		"offset":    {strconv.Itoa(offset)},
		"page_size": {strconv.Itoa(pageSize)},
	}
	if updateTimeFrom > 0 {
		q.Set("update_time_from", strconv.FormatInt(updateTimeFrom, 10))
	}
	if updateTimeTo > 0 {
		q.Set("update_time_to", strconv.FormatInt(updateTimeTo, 10))
	}
	for _, status := range itemStatus {
		q.Add("item_status", status)
	}
	result := &GetGlobalItemListResponse{}
	if err := s.client.DoGetMulti(context.Background(), PathGlobalProductGetItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GlobalItemInfo struct {
	ItemID       int64      `json:"item_id"`
	CategoryID   int64      `json:"category_id"`
	ItemName     string     `json:"item_name"`
	Description  string     `json:"description,omitempty"`
	ItemSKU      string     `json:"item_sku,omitempty"`
	CreateTime   int64      `json:"create_time"`
	UpdateTime   int64      `json:"update_time"`
	Image        *ImageInfo `json:"image,omitempty"`
	Weight       string     `json:"weight,omitempty"`
	Dimension    *Dimension `json:"dimension,omitempty"`
	ItemStatus   ItemStatus  `json:"item_status"`
	HasModel     bool        `json:"has_model"`
	PriceInfo    []PriceInfo `json:"price_info,omitempty"`
}

type GetGlobalItemInfoResponse struct {
	BaseResponse
	Response struct {
		ItemList []GlobalItemInfo `json:"item_list"`
	} `json:"response"`
}

func (s *GlobalProductService) GetItemInfo(itemIDs []int64) (*GetGlobalItemInfoResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetGlobalItemInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetItemInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddGlobalItemParams struct {
	CategoryID    int64           `json:"category_id"`
	ItemName      string          `json:"item_name"`
	Description   string          `json:"description"`
	ItemSKU       string          `json:"item_sku,omitempty"`
	Price         float64         `json:"price"`
	Stock         int             `json:"stock"`
	Weight        string          `json:"weight,omitempty"`
	Dimension     *Dimension      `json:"dimension,omitempty"`
	Image         *ImageInfo      `json:"image,omitempty"`
	TierVariation []TierVariation `json:"tier_variation,omitempty"`
}

type AddGlobalItemResponse struct {
	BaseResponse
	Response struct {
		ItemID int64 `json:"item_id"`
	} `json:"response,omitempty"`
}

func (s *GlobalProductService) AddItem(params *AddGlobalItemParams) (*AddGlobalItemResponse, error) {
	result := &AddGlobalItemResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductAddItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdateGlobalItemParams struct {
	ItemID      int64      `json:"item_id"`
	ItemName    string     `json:"item_name,omitempty"`
	Description string     `json:"description,omitempty"`
	ItemSKU     string     `json:"item_sku,omitempty"`
	Weight      string     `json:"weight,omitempty"`
	Dimension   *Dimension `json:"dimension,omitempty"`
	Image       *ImageInfo `json:"image,omitempty"`
}

func (s *GlobalProductService) UpdateItem(params *UpdateGlobalItemParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdateItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) DeleteItem(itemID int64) (*BaseResponse, error) {
	payload := map[string]any{"item_id": itemID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductDeleteItem, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) UpdateTierVariation(params *UpdateTierVariationParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdateTierVariation, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddGlobalModelParams struct {
	ItemID        int64   `json:"item_id"`
	TierIndex     []int   `json:"tier_index"`
	OriginalPrice float64 `json:"original_price"`
	Stock         int     `json:"stock"`
	ModelSKU      string  `json:"model_sku,omitempty"`
}

func (s *GlobalProductService) AddModel(params *AddGlobalModelParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductAddModel, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdateGlobalModelParams struct {
	ModelID       int64   `json:"model_id"`
	OriginalPrice float64 `json:"original_price,omitempty"`
	Stock         int     `json:"stock,omitempty"`
	ModelSKU      string  `json:"model_sku,omitempty"`
}

func (s *GlobalProductService) UpdateModel(params *UpdateGlobalModelParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdateModel, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) DeleteModel(modelID int64) (*BaseResponse, error) {
	payload := map[string]any{"model_id": modelID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductDeleteModel, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GlobalModelInfo struct {
	ModelID     *int64          `json:"model_id,omitempty"`
	TierIndex   []int           `json:"tier_index,omitempty"`
	ModelSKU    string          `json:"model_sku,omitempty"`
	ModelName   string          `json:"model_name,omitempty"`
	PriceInfo   []ModelPriceInfo `json:"price_info"`
}

type GetGlobalModelListResponse struct {
	BaseResponse
	Response *struct {
		TierVariation []TierVariation `json:"tier_variation"`
		Model         []GlobalModelInfo `json:"model"`
	} `json:"response,omitempty"`
}

func (s *GlobalProductService) GetModelList(itemID int64) (*GetGlobalModelListResponse, error) {
	q := map[string]string{"item_id": strconv.FormatInt(itemID, 10)}
	result := &GetGlobalModelListResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetModelList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) UpdateSizeChart(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdateSizeChart, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type CreatePublishTaskParams struct {
	GlobalItemID int64 `json:"global_item_id"`
	ShopID       int64 `json:"shop_id"`
}

type CreatePublishTaskResponse struct {
	BaseResponse
	Response struct {
		PublishTaskID int64 `json:"publish_task_id"`
	} `json:"response,omitempty"`
}

func (s *GlobalProductService) CreatePublishTask(params *CreatePublishTaskParams) (*CreatePublishTaskResponse, error) {
	result := &CreatePublishTaskResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductCreatePublishTask, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PublishableShopInfo struct {
	ShopID         int64  `json:"shop_id"`
	ShopName       string `json:"shop_name"`
	Region         string `json:"region"`
	Publishable    bool   `json:"publishable"`
}

type GetPublishableShopResponse struct {
	BaseResponse
	Response struct {
		ShopList []PublishableShopInfo `json:"shop_list"`
	} `json:"response"`
}

func (s *GlobalProductService) GetPublishableShop() (*GetPublishableShopResponse, error) {
	result := &GetPublishableShopResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetPublishableShop, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PublishTaskResultInfo struct {
	TaskID     int64  `json:"publish_task_id"`
	TaskStatus string `json:"task_status"`
	ItemID     int64  `json:"item_id,omitempty"`
	FailedReason string `json:"failed_reason,omitempty"`
}

type GetPublishTaskResultResponse struct {
	BaseResponse
	Response struct {
		TaskList []PublishTaskResultInfo `json:"task_list"`
	} `json:"response"`
}

func (s *GlobalProductService) GetPublishTaskResult(publishTaskID int64) (*GetPublishTaskResultResponse, error) {
	q := map[string]string{"publish_task_id": strconv.FormatInt(publishTaskID, 10)}
	result := &GetPublishTaskResultResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetPublishTaskResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PublishedItemInfo struct {
	ItemID      int64  `json:"item_id"`
	ItemName    string `json:"item_name"`
	ItemStatus  string `json:"item_status"`
	PublishedAt int64  `json:"published_at,omitempty"`
}

type GetPublishedListResponse struct {
	BaseResponse
	Response struct {
		ItemList    []PublishedItemInfo `json:"item_list"`
		TotalCount  int                 `json:"total_count"`
		HasNextPage bool                `json:"has_next_page"`
		NextOffset  int                 `json:"next_offset"`
	} `json:"response"`
}

func (s *GlobalProductService) GetPublishedList(offset, pageSize int) (*GetPublishedListResponse, error) {
	q := map[string]string{
		"offset":    strconv.Itoa(offset),
		"page_size": strconv.Itoa(pageSize),
	}
	result := &GetPublishedListResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetPublishedList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdateGlobalPriceParams struct {
	ItemID    int64             `json:"item_id"`
	PriceList []UpdatePriceItem `json:"price_list"`
}

func (s *GlobalProductService) UpdatePrice(params *UpdateGlobalPriceParams) (*UpdatePriceResponse, error) {
	result := &UpdatePriceResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdatePrice, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdateGlobalStockParams struct {
	ItemID    int64             `json:"item_id"`
	StockList []StockUpdateItem `json:"stock_list"`
}

func (s *GlobalProductService) UpdateStock(params *UpdateGlobalStockParams) (*UpdateStockResponse, error) {
	result := &UpdateStockResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdateStock, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SetSyncFieldParams struct {
	GlobalItemID int64    `json:"global_item_id"`
	SyncFields   []string `json:"sync_fields"`
}

func (s *GlobalProductService) SetSyncField(params *SetSyncFieldParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductSetSyncField, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type GlobalItemIDMapping struct {
	ItemID        int64 `json:"item_id"`
	GlobalItemID  int64 `json:"global_item_id"`
}

type GetGlobalItemIDResponse struct {
	BaseResponse
	Response struct {
		ItemList []GlobalItemIDMapping `json:"item_list"`
	} `json:"response"`
}

func (s *GlobalProductService) GetGlobalItemID(itemIDs []int64) (*GetGlobalItemIDResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetGlobalItemIDResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetGlobalItemID, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetRecommendAttribute(categoryID int64, language string) (*GetRecommendAttributeResponse, error) {
	q := map[string]string{
		"category_id": strconv.FormatInt(categoryID, 10),
	}
	if language != "" {
		q["language"] = language
	}
	result := &GetRecommendAttributeResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetRecommendAttr, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ShopPublishableStatus struct {
	ShopID            int64 `json:"shop_id"`
	IsPublishable     bool   `json:"is_publishable"`
}

func (s *GlobalProductService) GetShopPublishableStatus(shopIDs []int64) (*BaseResponse, error) {
	ids := make([]string, len(shopIDs))
	for i, id := range shopIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"shop_id_list": stringsJoin(ids, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetShopPublishable, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetVariations(itemIDs []int64) (*GetVariationsResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetVariationsResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetVariations, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetSizeChartList() (*GetSizeChartListResponse, error) {
	result := &GetSizeChartListResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetSizeChartList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) GetSizeChartDetail(sizeChartID int64) (*BaseResponse, error) {
	q := map[string]string{"size_chart_id": strconv.FormatInt(sizeChartID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetSizeChartDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) SearchAttributeValueList(categoryID, attributeID int64, keyword, language string) (*SearchAttributeValueResponse, error) {
	q := map[string]string{
		"category_id":  strconv.FormatInt(categoryID, 10),
		"attribute_id": strconv.FormatInt(attributeID, 10),
	}
	if keyword != "" {
		q["keyword"] = keyword
	}
	if language != "" {
		q["language"] = language
	}
	result := &SearchAttributeValueResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductSearchAttrValue, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type LocalAdjustmentRateInfo struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

type GetLocalAdjustmentRateResponse struct {
	BaseResponse
	Response struct {
		AdjustmentRate LocalAdjustmentRateInfo `json:"adjustment_rate"`
	} `json:"response"`
}

func (s *GlobalProductService) GetLocalAdjustmentRate(currency string) (*GetLocalAdjustmentRateResponse, error) {
	q := map[string]string{"currency": currency}
	result := &GetLocalAdjustmentRateResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetLocalAdjRate, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) UpdateLocalAdjustmentRate(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductUpdateLocalAdjRate, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type VideoUploadResultInfo struct {
	VideoUploadID string `json:"video_upload_id"`
	VideoURL      string `json:"video_url,omitempty"`
	UploadStatus  string `json:"upload_status"`
}

type GetVideoUploadResultResponse struct {
	BaseResponse
	Response struct {
		VideoInfo VideoUploadResultInfo `json:"video_info"`
	} `json:"response"`
}

func (s *GlobalProductService) GetVideoUploadResult(videoUploadID string) (*GetVideoUploadResultResponse, error) {
	q := map[string]string{"video_upload_id": videoUploadID}
	result := &GetVideoUploadResultResponse{}
	if err := s.client.DoGet(context.Background(), PathGlobalProductGetVideoUploadResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) CancelVideoUpload(videoUploadID string) (*BaseResponse, error) {
	payload := map[string]any{"video_upload_id": videoUploadID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathGlobalProductCancelVideoUpload, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *GlobalProductService) CategoryRecommend(globalitemname string) (*BaseResponse, error) {
    q := map[string]string{
		"global_item_name": globalitemname,
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathGlobalProductCategoryRecommend, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *GlobalProductService) InitTierVariation(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathGlobalProductInitTierVariation, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *GlobalProductService) SupportSizeChart(categoryid int64) (*BaseResponse, error) {
    q := map[string]string{
		"category_id": strconv.FormatInt(categoryid, 10),
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathGlobalProductSupportSizeChart, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}
