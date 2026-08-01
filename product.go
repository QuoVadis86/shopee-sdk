package shopee

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"
)

type ProductService struct {
	client *Client
}

func NewProductService(client *Client) *ProductService {
	return &ProductService{client: client}
}

type CategoryInfo struct {
	CategoryID       int64  `json:"category_id"`
	ParentCategoryID int64  `json:"parent_category_id"`
	CategoryName     string `json:"category_name"`
	HasChildren      bool   `json:"has_children"`
}

type GetCategoryResponse struct {
	BaseResponse
	Response struct {
		CategoryList []CategoryInfo `json:"category_list"`
	} `json:"response"`
}

func (s *ProductService) GetCategory(language string) (*GetCategoryResponse, error) {
	q := map[string]string{}
	if language != "" {
		q["language"] = language
	}
	result := &GetCategoryResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetCategory, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AttributeInfo struct {
	AttributeID       int64  `json:"attribute_id"`
	AttributeName     string `json:"attribute_name"`
	AttributeType     string `json:"attribute_type"`
	Mandatory         bool   `json:"mandatory"`
	ExampleValue      string `json:"example_value"`
	AttributeUnit     string `json:"attribute_unit"`
	OptionList        []struct {
		OptionID   int64  `json:"option_id"`
		OptionName string `json:"option_name"`
	} `json:"option_list,omitempty"`
}

type GetAttributeTreeResponse struct {
	BaseResponse
	Response struct {
		AttributeList []AttributeInfo `json:"attribute_list"`
	} `json:"response"`
}

func (s *ProductService) GetAttributeTree(categoryIDs []int64, language string) (*GetAttributeTreeResponse, error) {
	ids := make([]string, len(categoryIDs))
	for i, id := range categoryIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{
		"category_id_list": stringsJoin(ids, ","),
	}
	if language != "" {
		q["language"] = language
	}
	result := &GetAttributeTreeResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetAttributeTree, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BrandInfo struct {
	BrandID          int64  `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
	BrandStatus      string `json:"brand_status,omitempty"`
}

type GetBrandListResponse struct {
	BaseResponse
	Response struct {
		BrandList []BrandInfo `json:"brand_list"`
	} `json:"response"`
}

func (s *ProductService) GetBrandList(categoryID int64, language string, offset, pageSize int, status int64) (*GetBrandListResponse, error) {
	q := map[string]string{
		"category_id": strconv.FormatInt(categoryID, 10),
		"offset":      strconv.Itoa(offset),
		"page_size":   strconv.Itoa(pageSize),
		"status":      strconv.FormatInt(status, 10),
	}
	if language != "" {
		q["language"] = language
	}
	result := &GetBrandListResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetBrandList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ItemLimitInfo struct {
	ItemLimit int `json:"item_limit"`
	CategoryID int64 `json:"category_id,omitempty"`
}

type GetItemLimitResponse struct {
	BaseResponse
	Response struct {
		ItemLimit ItemLimitInfo `json:"item_limit"`
	} `json:"response"`
}

func (s *ProductService) GetItemLimit(itemName string, categoryID int64) (*GetItemLimitResponse, error) {
	q := map[string]string{
		"item_name": itemName,
	}
	if categoryID > 0 {
		q["category_id"] = strconv.FormatInt(categoryID, 10)
	}
	result := &GetItemLimitResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetItemLimit, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ItemListItem struct {
	ItemID     int64      `json:"item_id"`
	ItemStatus ItemStatus `json:"item_status"`
	UpdateTime int64      `json:"update_time"`
}

type GetItemListResponse struct {
	BaseResponse
	Response struct {
		Item        []ItemListItem `json:"item"`
		TotalCount  int            `json:"total_count"`
		HasNextPage bool           `json:"has_next_page"`
		NextOffset  int            `json:"next_offset"`
	} `json:"response"`
}

func (s *ProductService) GetItemList(offset, pageSize int, updateTimeFrom, updateTimeTo int64, itemStatus []string) (*GetItemListResponse, error) {
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
	result := &GetItemListResponse{}
	if err := s.client.DoGetMulti(context.Background(), PathProductGetItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type PriceInfo struct {
	Currency              string  `json:"currency"`
	OriginalPrice         float64 `json:"original_price"`
	CurrentPrice          float64 `json:"current_price"`
	InflatedPriceOriginal float64 `json:"inflated_price_of_original_price"`
	InflatedPriceCurrent  float64 `json:"inflated_price_of_current_price"`
}

type ImageInfo struct {
	ImageURLList []string `json:"image_url_list"`
	ImageIDList  []string `json:"image_id_list"`
	ImageRatio   string   `json:"image_ratio,omitempty"`
}

type Dimension struct {
	PackageLength float64 `json:"package_length"`
	PackageWidth  float64 `json:"package_width"`
	PackageHeight float64 `json:"package_height"`
}

type LogisticInfo struct {
	LogisticsID   int64   `json:"logistic_id"`
	LogisticsName string  `json:"logistic_name"`
	Enabled       bool    `json:"enabled"`
	FeeType       string  `json:"fee_type,omitempty"`
	ShippingFee   float64 `json:"shipping_fee,omitempty"`
	IsFree        bool    `json:"is_free"`
}

type PreOrder struct {
	IsPreOrder bool `json:"is_pre_order"`
	DaysToShip int  `json:"days_to_ship"`
}

type VideoInfo struct {
	VideoURL     string `json:"video_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Duration     int    `json:"duration"`
}

type ItemBaseInfo struct {
	ItemID       int64          `json:"item_id"`
	CategoryID   int64          `json:"category_id"`
	ItemName     string         `json:"item_name"`
	Description  string         `json:"description,omitempty"`
	ItemSKU      string         `json:"item_sku,omitempty"`
	CreateTime   int64          `json:"create_time"`
	UpdateTime   int64          `json:"update_time"`
	Image        *ImageInfo     `json:"image,omitempty"`
	Weight       string         `json:"weight,omitempty"`
	Dimension    *Dimension     `json:"dimension,omitempty"`
	LogisticInfo []LogisticInfo `json:"logistic_info,omitempty"`
	PreOrder     *PreOrder      `json:"pre_order,omitempty"`
	Condition    string         `json:"condition,omitempty"`
	ItemStatus   ItemStatus     `json:"item_status"`
	HasModel     bool           `json:"has_model"`
	Brand        *BrandInfo     `json:"brand,omitempty"`
	VideoInfo    []VideoInfo    `json:"video_info,omitempty"`
	PriceInfo    []PriceInfo    `json:"price_info,omitempty"`
}

type GetItemBaseInfoResponse struct {
	BaseResponse
	Response struct {
		ItemList []ItemBaseInfo `json:"item_list"`
	} `json:"response"`
}

func (s *ProductService) GetItemBaseInfo(itemIDs []int64) (*GetItemBaseInfoResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetItemBaseInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetItemBaseInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ItemExtraInfo struct {
	ItemID         int64  `json:"item_id"`
	Sales          int    `json:"sales,omitempty"`
	Views          int    `json:"views,omitempty"`
	RatingStar     float64 `json:"rating_star,omitempty"`
	CmtCount       int    `json:"cmt_count,omitempty"`
}

type GetItemExtraInfoResponse struct {
	BaseResponse
	Response struct {
		ItemList []ItemExtraInfo `json:"item_list"`
	} `json:"response"`
}

func (s *ProductService) GetItemExtraInfo(itemIDs []int64) (*GetItemExtraInfoResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetItemExtraInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetItemExtraInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddItemParams struct {
	CategoryID     int64           `json:"category_id"`
	ItemName       string          `json:"item_name"`
	Description    string          `json:"description"`
	ItemSKU        string          `json:"item_sku,omitempty"`
	Price          float64         `json:"price"`
	Stock          int             `json:"stock"`
	LogisticInfo   []LogisticInfo  `json:"logistic_info,omitempty"`
	Weight         float64         `json:"weight,omitempty"`
	Dimension      *Dimension      `json:"dimension,omitempty"`
	Image          *ImageInfo      `json:"image,omitempty"`
	Brand          *BrandInfo      `json:"brand,omitempty"`
	TierVariation  []TierVariation `json:"tier_variation,omitempty"`
	Condition      string          `json:"condition,omitempty"`
	VideoInfo      []VideoInfo     `json:"video_info,omitempty"`
	ItemStatus     ItemStatus      `json:"item_status,omitempty"`
}

type AddItemResponse struct {
	BaseResponse
	Response struct {
		ItemID int64 `json:"item_id"`
	} `json:"response,omitempty"`
}

func (s *ProductService) AddItem(params *AddItemParams) (*AddItemResponse, error) {
	result := &AddItemResponse{}
	if err := s.client.DoPost(context.Background(), PathProductAddItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdateItemParams struct {
	ItemID       int64          `json:"item_id"`
	ItemName     string         `json:"item_name,omitempty"`
	Description  string         `json:"description,omitempty"`
	ItemSKU      string         `json:"item_sku,omitempty"`
	Weight       string         `json:"weight,omitempty"`
	Dimension    *Dimension     `json:"dimension,omitempty"`
	Image        *ImageInfo     `json:"image,omitempty"`
	LogisticInfo []LogisticInfo `json:"logistic_info,omitempty"`
	Brand        *BrandInfo     `json:"brand,omitempty"`
	Condition    string         `json:"condition,omitempty"`
	VideoInfo    []VideoInfo    `json:"video_info,omitempty"`
}

type UpdateItemResponse struct {
	BaseResponse
	Response struct {
		ItemID int64 `json:"item_id"`
	} `json:"response,omitempty"`
}

func (s *ProductService) UpdateItem(params *UpdateItemParams) (*UpdateItemResponse, error) {
	result := &UpdateItemResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdateItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) DeleteItem(itemIDs []int64) (*BaseResponse, error) {
	payload := map[string]any{"item_id_list": itemIDs}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductDeleteItem, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type TierVariation struct {
	Name       string `json:"name"`
	OptionList []struct {
		Option string `json:"option"`
		Image  *struct {
			ImageID  string `json:"image_id"`
			ImageURL string `json:"image_url"`
		} `json:"image,omitempty"`
	} `json:"option_list"`
}

type UpdateTierVariationParams struct {
	ItemID        int64           `json:"item_id"`
	TierVariation []TierVariation `json:"tier_variation"`
}

func (s *ProductService) UpdateTierVariation(params *UpdateTierVariationParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdateTierVariation, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ModelPriceInfo struct {
	Currency              string  `json:"currency"`
	CurrentPrice          float64 `json:"current_price"`
	OriginalPrice         float64 `json:"original_price"`
	InflatedPriceOriginal float64 `json:"inflated_price_of_original_price"`
	InflatedPriceCurrent  float64 `json:"inflated_price_of_current_price"`
}

type ModelInfo struct {
	ModelID     *int64          `json:"model_id,omitempty"`
	TierIndex   []int           `json:"tier_index,omitempty"`
	ModelSKU    string          `json:"model_sku,omitempty"`
	ModelName   string          `json:"model_name,omitempty"`
	ModelStatus string          `json:"model_status,omitempty"`
	PriceInfo   []ModelPriceInfo `json:"price_info"`
	PromotionID int64           `json:"promotion_id,omitempty"`
}

type GetModelListResponse struct {
	BaseResponse
	Response *struct {
		TierVariation []TierVariation `json:"tier_variation"`
		Model         []ModelInfo     `json:"model"`
	} `json:"response,omitempty"`
}

func (s *ProductService) GetModelList(itemID int64) (*GetModelListResponse, error) {
	q := map[string]string{"item_id": strconv.FormatInt(itemID, 10)}
	result := &GetModelListResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetModelList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AddModelParams struct {
	ItemID         int64   `json:"item_id"`
	TierIndex      []int   `json:"tier_index"`
	ModelSKU       string  `json:"model_sku,omitempty"`
	OriginalPrice  float64 `json:"original_price"`
	Stock          int     `json:"stock"`
}

func (s *ProductService) AddModel(params *AddModelParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductAddModel, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdateModelParams struct {
	ModelID       int64   `json:"model_id"`
	ModelSKU      string  `json:"model_sku,omitempty"`
	OriginalPrice float64 `json:"original_price,omitempty"`
	Stock         int     `json:"stock,omitempty"`
}

type UpdateModelResponse struct {
	BaseResponse
	Response *struct {
		ModelID int64 `json:"model_id"`
	} `json:"response,omitempty"`
}

func (s *ProductService) UpdateModel(params *UpdateModelParams) (*UpdateModelResponse, error) {
	result := &UpdateModelResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdateModel, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) DeleteModel(modelID int64) (*BaseResponse, error) {
	payload := map[string]any{"model_id": modelID}
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductDeleteModel, payload, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UpdatePriceItem struct {
	ModelID       int64   `json:"model_id"`
	OriginalPrice float64 `json:"original_price"`
}

type UpdatePriceParams struct {
	ItemID    int64             `json:"item_id"`
	PriceList []UpdatePriceItem `json:"price_list"`
}

type UpdatePriceResponse struct {
	BaseResponse
	Response *struct {
		SuccessList []UpdatePriceItem `json:"success_list,omitempty"`
		FailureList []struct {
			ModelID      int64  `json:"model_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list,omitempty"`
	} `json:"response,omitempty"`
}

func (s *ProductService) UpdatePrice(params *UpdatePriceParams) (*UpdatePriceResponse, error) {
	result := &UpdatePriceResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdatePrice, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SellerStockUpdate struct {
	LocationID string `json:"location_id"`
	Stock      int    `json:"stock"`
}

type StockUpdateItem struct {
	ModelID     int64               `json:"model_id"`
	SellerStock []SellerStockUpdate `json:"seller_stock"`
}

type UpdateStockParams struct {
	ItemID    int64             `json:"item_id"`
	StockList []StockUpdateItem `json:"stock_list"`
}

type UpdateStockResponse struct {
	BaseResponse
	Response *struct {
		SuccessList []StockUpdateItem `json:"success_list,omitempty"`
		FailureList []struct {
			ModelID      int64  `json:"model_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list,omitempty"`
	} `json:"response,omitempty"`
}

func (s *ProductService) UpdateStock(params *UpdateStockParams) (*UpdateStockResponse, error) {
	result := &UpdateStockResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdateStock, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BoostedItemInfo struct {
	ItemID    int64  `json:"item_id"`
	Boosted   bool   `json:"boosted"`
	BoostCost int    `json:"boost_cost,omitempty"`
}

type GetBoostedListResponse struct {
	BaseResponse
	Response struct {
		ItemList []BoostedItemInfo `json:"item_list"`
		Total    int               `json:"total"`
	} `json:"response"`
}

func (s *ProductService) GetBoostedList(pageSize, pageNumber int) (*GetBoostedListResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
	}
	result := &GetBoostedListResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetBoostedList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetItemPromotion(itemIDs []int64) (*GetItemBaseInfoResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetItemBaseInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetItemPromotion, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SIPItemPriceParams struct {
	ItemID int64   `json:"item_id"`
	Price  float64 `json:"price"`
}

func (s *ProductService) UpdateSIPItemPrice(params *SIPItemPriceParams) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdateSIPItemPrice, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ProductCommentInfo struct {
	CommentID   int64  `json:"comment_id"`
	ItemID      int64  `json:"item_id"`
	CommentText string `json:"comment_text"`
	Rating      int    `json:"rating"`
	CreateTime  int64  `json:"create_time"`
	BuyerName   string `json:"buyer_username,omitempty"`
}

type GetCommentResponse struct {
	BaseResponse
	Response struct {
		CommentList []ProductCommentInfo `json:"comment_list"`
		TotalCount  int           `json:"total_count"`
	} `json:"response"`
}

func (s *ProductService) GetComment(itemID int64, pageSize int, cursor string, createTimeFrom, createTimeTo int64) (*GetCommentResponse, error) {
	q := map[string]string{
		"item_id":   strconv.FormatInt(itemID, 10),
		"page_size": strconv.Itoa(pageSize),
	}
	if cursor != "" {
		q["cursor"] = cursor
	}
	if createTimeFrom > 0 {
		q["create_time_from"] = strconv.FormatInt(createTimeFrom, 10)
	}
	if createTimeTo > 0 {
		q["create_time_to"] = strconv.FormatInt(createTimeTo, 10)
	}
	result := &GetCommentResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetComment, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type RecommendAttributeInfo struct {
	AttributeID   int64  `json:"attribute_id"`
	AttributeName string `json:"attribute_name"`
	AttributeType string `json:"attribute_type"`
}

type GetRecommendAttributeResponse struct {
	BaseResponse
	Response struct {
		AttributeList []RecommendAttributeInfo `json:"attribute_list"`
	} `json:"response"`
}

func (s *ProductService) GetRecommendAttribute(categoryID int64, itemName, language string) (*GetRecommendAttributeResponse, error) {
	q := map[string]string{
		"category_id": strconv.FormatInt(categoryID, 10),
		"item_name":   itemName,
	}
	if language != "" {
		q["language"] = language
	}
	result := &GetRecommendAttributeResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetRecommendAttr, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type WeightRecommendationInfo struct {
	MinWeight float64 `json:"min_weight"`
	MaxWeight float64 `json:"max_weight"`
	WeightUnit string `json:"weight_unit"`
}

type GetWeightRecommendationResponse struct {
	BaseResponse
	Response struct {
		WeightRecommendation WeightRecommendationInfo `json:"weight_recommendation"`
	} `json:"response"`
}

func (s *ProductService) GetWeightRecommendation(categoryID, brandID int64, itemName, coverImageID, descriptionType string, attributeList []map[string]any) (*GetWeightRecommendationResponse, error) {
	attrs, _ := json.Marshal(attributeList)
	q := map[string]string{
		"category_id":      strconv.FormatInt(categoryID, 10),
		"brand_id":         strconv.FormatInt(brandID, 10),
		"item_name":        itemName,
		"cover_image_id":   coverImageID,
		"description_type": descriptionType,
	}
	if len(attributeList) > 0 {
		q["attribute_list"] = string(attrs)
	}
	result := &GetWeightRecommendationResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetWeightRec, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SizeChartInfo struct {
	SizeChartID   int64  `json:"size_chart_id"`
	SizeChartName string `json:"size_chart_name"`
}

type GetSizeChartListResponse struct {
	BaseResponse
	Response struct {
		SizeChartList []SizeChartInfo `json:"size_chart_list"`
	} `json:"response"`
}

func (s *ProductService) GetSizeChartList() (*GetSizeChartListResponse, error) {
	result := &GetSizeChartListResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetSizeChartList, map[string]string{}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetSizeChartDetail(sizeChartID int64) (*BaseResponse, error) {
	q := map[string]string{"size_chart_id": strconv.FormatInt(sizeChartID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetSizeChartDetail, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type VariationInfo struct {
	ItemID      int64   `json:"item_id"`
	VariationID int64   `json:"variation_id"`
	VariationName string `json:"variation_name"`
	OriginalPrice float64 `json:"original_price"`
	Stock        int     `json:"stock"`
}

type GetVariationsResponse struct {
	BaseResponse
	Response struct {
		VariationList []VariationInfo `json:"variation_list"`
	} `json:"response"`
}

func (s *ProductService) GetVariations(categoryID int64) (*GetVariationsResponse, error) {
	q := map[string]string{"category_id": strconv.FormatInt(categoryID, 10)}
	result := &GetVariationsResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetVariations, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type VehicleInfo struct {
	VehicleID   int64  `json:"vehicle_id"`
	VehicleName string `json:"vehicle_name"`
}

type GetAllVehicleListResponse struct {
	BaseResponse
	Response struct {
		VehicleList []VehicleInfo `json:"vehicle_list"`
	} `json:"response"`
}

func (s *ProductService) GetAllVehicleList(pageSize int64) (*GetAllVehicleListResponse, error) {
	q := map[string]string{"page_size": strconv.FormatInt(pageSize, 10)}
	result := &GetAllVehicleListResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetAllVehicleList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetVehicleListByCompatibilityDetail(compatibilityDetails string) (*BaseResponse, error) {
	q := map[string]string{
		"compatibility_details": compatibilityDetails,
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetVehicleCompList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetItemContentDiagnosisResult(itemIDs []int64) (*BaseResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{
		"item_id_list": stringsJoin(ids, ","),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetContentDiagResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetItemListByContentDiagnosis(diagnosisType string, offset, pageSize int) (*BaseResponse, error) {
	q := map[string]string{
		"diagnosis_type": diagnosisType,
		"offset":         strconv.Itoa(offset),
		"page_size":      strconv.Itoa(pageSize),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetContentDiagList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type KitLimitInfo struct {
	KitItemLimit int `json:"kit_item_limit"`
}

type GetKitItemLimitResponse struct {
	BaseResponse
	Response struct {
		KitItemLimit KitLimitInfo `json:"kit_item_limit"`
	} `json:"response"`
}

func (s *ProductService) GetKitItemLimit(categoryID int64) (*GetKitItemLimitResponse, error) {
	q := map[string]string{"category_id": strconv.FormatInt(categoryID, 10)}
	result := &GetKitItemLimitResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetKitLimit, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) AddKitItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductAddKitItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) UpdateKitItem(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUpdateKitItem, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetKitItemInfo(itemID int64) (*BaseResponse, error) {
	q := map[string]string{"item_id": strconv.FormatInt(itemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetKitItemInfo, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type SSPInfo struct {
	SSPID    int64  `json:"ssp_id"`
	SSPName  string `json:"ssp_name"`
	SSPStatus string `json:"ssp_status"`
}

type GetSSPListResponse struct {
	BaseResponse
	Response struct {
		SSPList      []SSPInfo `json:"ssp_list"`
		TotalCount   int       `json:"total_count"`
	} `json:"response"`
}




func (s *ProductService) GetAItemByPItemID(pItemID int64) (*BaseResponse, error) {
	q := map[string]string{"p_item_id": strconv.FormatInt(pItemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetAItemByPItemID, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type AttributeValueInfo struct {
	AttributeValueID   int64  `json:"attribute_value_id"`
	AttributeValueName string `json:"attribute_value_name"`
}

type SearchAttributeValueResponse struct {
	BaseResponse
	Response struct {
		AttributeValueList []AttributeValueInfo `json:"attribute_value_list"`
	} `json:"response"`
}

func (s *ProductService) SearchAttributeValueList(attributeID, cursor, limit int64) (*SearchAttributeValueResponse, error) {
	q := map[string]string{
		"attribute_id": strconv.FormatInt(attributeID, 10),
		"cursor":       strconv.FormatInt(cursor, 10),
		"limit":        strconv.FormatInt(limit, 10),
	}
	result := &SearchAttributeValueResponse{}
	if err := s.client.DoGet(context.Background(), PathProductSearchAttrValue, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetMainItemList(directItemIDs []int64) (*GetItemBaseInfoResponse, error) {
	ids := make([]string, len(directItemIDs))
	for i, id := range directItemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"direct_item_id": stringsJoin(ids, ",")}
	result := &GetItemBaseInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetMainItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetDirectItemList(mainItemIDs []int64) (*BaseResponse, error) {
	ids := make([]string, len(mainItemIDs))
	for i, id := range mainItemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"main_item_id": stringsJoin(ids, ",")}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetDirectItemList, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetDirectShopRecommendedPrice(itemID int64) (*BaseResponse, error) {
	q := map[string]string{"item_id": strconv.FormatInt(itemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetDirectShopPrice, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetProductCertificationRule(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductGetCertRule, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) PublishItemToOutletShop(params any) (*BaseResponse, error) {
	result := &BaseResponse{}
	if err := s.client.DoPost(context.Background(), PathProductPublishToOutlet, params, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetMartItemMappingByID(martItemID int64) (*BaseResponse, error) {
	q := map[string]string{"mart_item_id": strconv.FormatInt(martItemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetMartItemMapping, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) SearchUnpackagedModelList(pageSize, pageNumber int) (*BaseResponse, error) {
	q := map[string]string{
		"page_size":   strconv.Itoa(pageSize),
		"page_no": strconv.Itoa(pageNumber),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductSearchUnpackaged, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) GetMartItemByOutletItemID(outletItemID int64) (*BaseResponse, error) {
	q := map[string]string{"outlet_item_id": strconv.FormatInt(outletItemID, 10)}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetMartItemByOutlet, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type ItemViolationInfo struct {
	ItemID       int64  `json:"item_id"`
	ViolationType string `json:"violation_type"`
	ViolationReason string `json:"violation_reason"`
	PenaltyPoints int    `json:"penalty_points,omitempty"`
}

type GetItemViolationInfoResponse struct {
	BaseResponse
	Response struct {
		ItemList []ItemViolationInfo `json:"item_list"`
	} `json:"response"`
}

func (s *ProductService) GetItemViolationInfo(itemIDs []int64) (*GetItemViolationInfoResponse, error) {
	ids := make([]string, len(itemIDs))
	for i, id := range itemIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := map[string]string{"item_id_list": stringsJoin(ids, ",")}
	result := &GetItemViolationInfoResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetItemViolation, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type BoostItemParams struct {
	ItemIDList []int64 `json:"item_id_list"`
}

type BoostItemResponse struct {
	BaseResponse
	Response *struct {
		FailureList []struct {
			ItemID       int64  `json:"item_id"`
			FailedReason string `json:"failed_reason"`
		} `json:"failure_list,omitempty"`
		SuccessList []struct {
			ItemIDList []int64 `json:"item_id_list"`
		} `json:"success_list,omitempty"`
	} `json:"response,omitempty"`
}

func (s *ProductService) BoostItem(itemIDs []int64) (*BoostItemResponse, error) {
	result := &BoostItemResponse{}
	if err := s.client.DoPost(context.Background(), PathProductBoostItem, &BoostItemParams{ItemIDList: itemIDs}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type UnlistItemReq struct {
	ItemID int64 `json:"item_id"`
	Unlist bool  `json:"unlist"`
}

type UnlistItemParams struct {
	ItemList []UnlistItemReq `json:"item_list"`
}

type UnlistItemResponse struct {
	BaseResponse
	Response *struct {
		Result []struct {
			ItemID       int64  `json:"item_id"`
			Success      bool   `json:"success"`
			FailedReason string `json:"failed_reason,omitempty"`
		} `json:"result,omitempty"`
	} `json:"response,omitempty"`
}

func (s *ProductService) UnlistItem(itemIDs []int64, unlist bool) (*UnlistItemResponse, error) {
	list := make([]UnlistItemReq, len(itemIDs))
	for i, id := range itemIDs {
		list[i] = UnlistItemReq{ItemID: id, Unlist: unlist}
	}
	result := &UnlistItemResponse{}
	if err := s.client.DoPost(context.Background(), PathProductUnlistItem, &UnlistItemParams{ItemList: list}, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

func (s *ProductService) BatchAddItem(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductBatchAddItem, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) BatchPublishItemToOutletShop(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductBatchPublishItemToOutletShop, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) BatchUpdateOutletPrice(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductBatchUpdateOutletPrice, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) BatchUpdateOutletStock(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductBatchUpdateOutletStock, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) CategoryRecommend(itemname string) (*BaseResponse, error) {
    q := map[string]string{
		"item_name": itemname,
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathProductCategoryRecommend, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) GenerateKitImage(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductGenerateKitImage, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) InitTierVariation(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductInitTierVariation, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) RegisterBrand(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductRegisterBrand, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) ReplyComment(params any) (*BaseResponse, error) {
    result := &BaseResponse{}
    if err := s.client.DoPost(context.Background(), PathProductReplyComment, params, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) SearchItem(pagesize int64) (*BaseResponse, error) {
    q := map[string]string{
		"page_size": strconv.FormatInt(pagesize, 10),
	}
    result := &BaseResponse{}
    if err := s.client.DoGet(context.Background(), PathProductSearchItem, q, result); err != nil {
    	return nil, err
    }
    if result.HasError() {
    	return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
    }
    return result, nil
}

func (s *ProductService) GetBatchTaskResult(taskType string, taskID int64) (*BaseResponse, error) {
	q := map[string]string{
		"task_type": taskType,
		"task_id":   strconv.FormatInt(taskID, 10),
	}
	result := &BaseResponse{}
	if err := s.client.DoGet(context.Background(), PathProductGetBatchTaskResult, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
