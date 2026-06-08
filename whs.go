package shopee

import "strconv"

type WHSService struct {
	client *Client
}

func NewWHSService(client *Client) *WHSService {
	return &WHSService{client: client}
}

type WHSInventoryItem struct {
	SKU          string `json:"sku"`
	WarehouseID  int64  `json:"warehouse_id"`
	CurrentStock int    `json:"current_stock"`
	ReservedStock int   `json:"reserved_stock"`
	AvailableStock int  `json:"available_stock"`
}

type GetCurrentInventoryResponse struct {
	BaseResponse
	Response struct {
		InventoryList []WHSInventoryItem `json:"inventory_list"`
		Total         int                `json:"total"`
		PageNum       int                `json:"page_num"`
		PageSize      int                `json:"page_size"`
	} `json:"response"`
}

func (s *WHSService) GetCurrentInventory(warehouseID int64, skuList []string, pageSize, pageNumber int) (*GetCurrentInventoryResponse, error) {
	q := map[string]string{
		"warehouse_id": strconv.FormatInt(warehouseID, 10),
		"page_size":    strconv.Itoa(pageSize),
		"page_number":  strconv.Itoa(pageNumber),
	}
	if len(skuList) > 0 {
		q["sku_list"] = stringsJoin(skuList, ",")
	}
	result := &GetCurrentInventoryResponse{}
	if err := s.client.DoGet(PathWHSGetCurrentInventory, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type WHSExpiryItem struct {
	SKU         string `json:"sku"`
	WarehouseID int64  `json:"warehouse_id"`
	ExpiryDate  string `json:"expiry_date"`
	Quantity    int    `json:"quantity"`
}

type GetExpiryReportResponse struct {
	BaseResponse
	Response struct {
		ExpiryList []WHSExpiryItem `json:"expiry_list"`
		Total      int             `json:"total"`
		PageNum    int             `json:"page_num"`
		PageSize   int             `json:"page_size"`
	} `json:"response"`
}

func (s *WHSService) GetExpiryReport(warehouseID int64, pageSize, pageNumber int) (*GetExpiryReportResponse, error) {
	q := map[string]string{
		"warehouse_id": strconv.FormatInt(warehouseID, 10),
		"page_size":    strconv.Itoa(pageSize),
		"page_number":  strconv.Itoa(pageNumber),
	}
	result := &GetExpiryReportResponse{}
	if err := s.client.DoGet(PathWHSGetExpiryReport, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type WHSStockAgingItem struct {
	SKU            string `json:"sku"`
	WarehouseID    int64  `json:"warehouse_id"`
	DaysInStock    int    `json:"days_in_stock"`
	Quantity       int    `json:"quantity"`
	AgingCategory  string `json:"aging_category,omitempty"`
}

type GetStockAgingResponse struct {
	BaseResponse
	Response struct {
		AgingList []WHSStockAgingItem `json:"aging_list"`
		Total     int                 `json:"total"`
		PageNum   int                 `json:"page_num"`
		PageSize  int                 `json:"page_size"`
	} `json:"response"`
}

func (s *WHSService) GetStockAging(warehouseID int64, pageSize, pageNumber int) (*GetStockAgingResponse, error) {
	q := map[string]string{
		"warehouse_id": strconv.FormatInt(warehouseID, 10),
		"page_size":    strconv.Itoa(pageSize),
		"page_number":  strconv.Itoa(pageNumber),
	}
	result := &GetStockAgingResponse{}
	if err := s.client.DoGet(PathWHSGetStockAging, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}

type WHSStockMovementItem struct {
	SKU          string `json:"sku"`
	WarehouseID  int64  `json:"warehouse_id"`
	MovementType string `json:"movement_type"`
	Quantity     int    `json:"quantity"`
	CreatedAt    int64  `json:"create_time"`
}

type GetStockMovementResponse struct {
	BaseResponse
	Response struct {
		MovementList []WHSStockMovementItem `json:"movement_list"`
		Total        int                    `json:"total"`
		PageNum      int                    `json:"page_num"`
		PageSize     int                    `json:"page_size"`
	} `json:"response"`
}

func (s *WHSService) GetStockMovement(warehouseID int64, dateFrom, dateTo int64, pageSize, pageNumber int) (*GetStockMovementResponse, error) {
	q := map[string]string{
		"warehouse_id": strconv.FormatInt(warehouseID, 10),
		"page_size":    strconv.Itoa(pageSize),
		"page_number":  strconv.Itoa(pageNumber),
	}
	if dateFrom > 0 {
		q["date_from"] = strconv.FormatInt(dateFrom, 10)
	}
	if dateTo > 0 {
		q["date_to"] = strconv.FormatInt(dateTo, 10)
	}
	result := &GetStockMovementResponse{}
	if err := s.client.DoGet(PathWHSGetStockMovement, q, result); err != nil {
		return nil, err
	}
	if result.HasError() {
		return nil, &APIError{ErrorCode: result.Error, Message: result.Message, RequestID: result.RequestID}
	}
	return result, nil
}
