package shopee

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func getClient(t *testing.T) *Client {
	t.Helper()

	pidStr := os.Getenv("SHOPEE_PARTNER_ID")
	key := os.Getenv("SHOPEE_PARTNER_KEY")
	token := os.Getenv("SHOPEE_ACCESS_TOKEN")
	sidStr := os.Getenv("SHOPEE_SHOP_ID")

	if pidStr == "" || key == "" || token == "" || sidStr == "" {
		t.Skip("set SHOPEE_PARTNER_ID, SHOPEE_PARTNER_KEY, SHOPEE_ACCESS_TOKEN, SHOPEE_SHOP_ID")
	}

	pid, _ := strconv.ParseInt(pidStr, 10, 64)
	sid, _ := strconv.ParseInt(sidStr, 10, 64)

	return NewClient(pid, key, token, sid, WithRegion(RegionSandbox))
}

func daysAgo(days int) int64 {
	return time.Now().AddDate(0, 0, -days).Unix()
}

func nowUnix() int64 {
	return time.Now().Unix()
}

// ---------------------------------------------------------------------------
// Product
// ---------------------------------------------------------------------------

func TestIntegrationProduct_GetCategory(t *testing.T) {
	c := getClient(t)
	s := &ProductService{client: c}

	resp, err := s.GetCategory("en")
	if err != nil {
		t.Fatalf("GetCategory failed: %v", err)
	}
	t.Logf("Categories: %d", len(resp.Response.CategoryList))
	if len(resp.Response.CategoryList) == 0 {
		t.Fatal("expected at least 1 category")
	}
	t.Logf("First: ID=%d Name=%s", resp.Response.CategoryList[0].CategoryID, resp.Response.CategoryList[0].CategoryName)
}

func TestIntegrationProduct_GetItemList(t *testing.T) {
	c := getClient(t)
	s := &ProductService{client: c}

	from := daysAgo(30)
	to := nowUnix()
	resp, err := s.GetItemList(0, 10, from, to, "")
	if err != nil {
		t.Logf("GetItemList (sandbox may have no items): %v", err)
		return
	}
	t.Logf("Items: %d total, hasNextPage: %v", resp.Response.TotalCount, resp.Response.HasNextPage)
}

func TestIntegrationProduct_AddItem(t *testing.T) {
	c := getClient(t)
	s := &ProductService{client: c}

	catResp, err := s.GetCategory("en")
	if err != nil {
		t.Fatal(err)
	}
	var leafID int64
	for _, cat := range catResp.Response.CategoryList {
		if !cat.HasChildren {
			leafID = cat.CategoryID
			break
		}
	}
	if leafID == 0 {
		t.Fatal("no leaf category found")
	}
	t.Logf("Using leaf category: ID=%d", leafID)

	params := &AddItemParams{
		CategoryID:  leafID,
		ItemName:    "Go SDK Test Item",
		Description: "Created by integration test",
		Price:       99.99,
		Stock:       10,
		Weight:      0.5,
		Image: &ImageInfo{
			ImageURLList: []string{"https://cf.shopee.sg/file/demo-product-image"},
		},
		Condition: "NEW",
	}

	resp, err := s.AddItem(params)
	if err != nil {
		t.Logf("AddItem (sandbox may require brand): %v", err)
		return
	}
	if resp.Response.ItemID <= 0 {
		t.Fatal("expected valid item_id")
	}
	t.Logf("Created item ID: %d", resp.Response.ItemID)
}

func TestIntegrationProduct_BoostItem(t *testing.T) {
	c := getClient(t)
	s := &ProductService{client: c}

	resp, err := s.BoostItem([]int64{100004})
	if err != nil {
		t.Logf("BoostItem (expected in sandbox): %v", err)
		return
	}
	t.Logf("BoostItem: success=%v", resp.Response)
}

// ---------------------------------------------------------------------------
// Order
// ---------------------------------------------------------------------------

func TestIntegrationOrder_GetOrderList(t *testing.T) {
	c := getClient(t)
	s := &OrderService{client: c}

	from := daysAgo(7)
	to := nowUnix()
	resp, err := s.GetOrderList("create_time", from, to, 10, "", "")
	if err != nil {
		t.Logf("GetOrderList (sandbox may have no orders): %v", err)
		return
	}
	t.Logf("Orders: %d, more: %v", len(resp.Response.OrderList), resp.Response.More)

	if len(resp.Response.OrderList) > 0 {
		first := resp.Response.OrderList[0]
		detail, err := s.GetOrderDetail([]string{first.OrderSN}, "item_list,package_list,recipient_address")
		if err != nil {
			t.Fatalf("GetOrderDetail failed: %v", err)
		}
		t.Logf("Order %s: items=%d packages=%d", first.OrderSN,
			len(detail.Response.OrderList[0].ItemList),
			len(detail.Response.OrderList[0].PackageList))
	}
}

func TestIntegrationOrder_GetShipmentList(t *testing.T) {
	c := getClient(t)
	s := &OrderService{client: c}

	resp, err := s.GetShipmentList(10, "")
	if err != nil {
		t.Logf("GetShipmentList: %v", err)
		return
	}
	t.Logf("Shipments: %d, more: %v", len(resp.Response.OrderList), resp.Response.More)
}

func TestIntegrationOrder_SetNote(t *testing.T) {
	c := getClient(t)
	s := &OrderService{client: c}

	resp, err := s.SetNote(&SetNoteParams{
		OrderSN: "sandbox-test-order-sn",
		Note:    "Integration test note",
	})
	if err != nil {
		t.Logf("SetNote (expected in sandbox): %v", err)
		return
	}
	t.Logf("SetNote: request_id=%s", resp.RequestID)
}

// ---------------------------------------------------------------------------
// Logistics
// ---------------------------------------------------------------------------

func TestIntegrationLogistics_GetChannelList(t *testing.T) {
	c := getClient(t)
	s := &LogisticsService{client: c}

	resp, err := s.GetChannelList()
	if err != nil {
		t.Fatalf("GetChannelList failed: %v", err)
	}
	t.Logf("Channels: %d", len(resp.Response.LogisticsChannelList))
}

func TestIntegrationLogistics_GetAddressList(t *testing.T) {
	c := getClient(t)
	s := &LogisticsService{client: c}

	resp, err := s.GetAddressList()
	if err != nil {
		t.Fatalf("GetAddressList failed: %v", err)
	}
	t.Logf("Addresses: %d", len(resp.Response.AddressList))
}

// ---------------------------------------------------------------------------
// Shop
// ---------------------------------------------------------------------------

func TestIntegrationShop_GetShopInfo(t *testing.T) {
	c := getClient(t)
	s := &ShopService{client: c}

	result, err := s.GetShopInfo()
	if err != nil {
		t.Fatalf("GetShopInfo failed: %v", err)
	}
	if result.Response == nil {
		t.Logf("Response is nil (sandbox limitation for shop_info).")
		return
	}
	t.Logf("Shop: %s (ID: %d, Region: %s, Status: %s)",
		result.Response.ShopName, result.Response.ShopID, result.Response.Region, result.Response.Status)
}

func TestIntegrationShop_GetProfile(t *testing.T) {
	c := getClient(t)
	s := &ShopService{client: c}

	resp, err := s.GetProfile()
	if err != nil {
		t.Fatalf("GetProfile failed: %v", err)
	}
	if resp.Response != nil {
		t.Logf("Shop name: %s", resp.Response.ShopName)
	}
}

func TestIntegrationShop_GetWarehouseDetail(t *testing.T) {
	c := getClient(t)
	s := &ShopService{client: c}

	resp, err := s.GetWarehouseDetail(1)
	if err != nil {
		t.Logf("GetWarehouseDetail: %v", err)
		return
	}
	t.Logf("Warehouse: %s", resp.Response.WarehouseName)
}

// ---------------------------------------------------------------------------
// AMS
// ---------------------------------------------------------------------------

func TestIntegrationAMS_GetOpenCampaignAddedProduct(t *testing.T) {
	c := getClient(t)
	s := &AMSService{client: c}

	resp, err := s.GetOpenCampaignAddedProduct(10, "", "", "", "")
	if err != nil {
		t.Logf("GetOpenCampaignAddedProduct: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("AMS products: %d total, has_more=%v", resp.Response.TotalCount, resp.Response.HasMore)
	}
}

func TestIntegrationAMS_GetShopSuggestedRate(t *testing.T) {
	c := getClient(t)
	s := &AMSService{client: c}

	resp, err := s.GetShopSuggestedRate()
	if err != nil {
		t.Logf("GetShopSuggestedRate: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("Suggested rate: %.4f", resp.Response.SuggestedRate)
	}
}

func TestIntegrationAMS_GetAutoAddToggleStatus(t *testing.T) {
	c := getClient(t)
	s := &AMSService{client: c}

	resp, err := s.GetAutoAddToggleStatus()
	if err != nil {
		t.Logf("GetAutoAddToggleStatus: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("Auto-add enabled: %v", resp.Response.AutoAddEnabled)
	}
}

// ---------------------------------------------------------------------------
// Payment
// ---------------------------------------------------------------------------

func TestIntegrationPayment_GetPaymentMethodList(t *testing.T) {
	c := getClient(t)
	s := &PaymentService{client: c}

	resp, err := s.GetPaymentMethodList()
	if err != nil {
		t.Logf("GetPaymentMethodList: %v", err)
		return
	}
	t.Logf("Payment methods: %d", len(resp.Response.PaymentMethodList))
}

func TestIntegrationPayment_GetShopInstallmentStatus(t *testing.T) {
	c := getClient(t)
	s := &PaymentService{client: c}

	resp, err := s.GetShopInstallmentStatus()
	if err != nil {
		t.Logf("GetShopInstallmentStatus: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("Installment enabled: %v", resp.Response.Enabled)
	}
}

// ---------------------------------------------------------------------------
// Promotion
// ---------------------------------------------------------------------------

func TestIntegrationPromotion_GetDiscountList(t *testing.T) {
	c := getClient(t)
	s := &PromotionService{client: c}

	resp, err := s.GetDiscountList(10, 1, "")
	if err != nil {
		t.Logf("GetDiscountList: %v", err)
		return
	}
	t.Logf("Discounts: %d", len(resp.Response.DiscountList))
}

func TestIntegrationPromotion_GetVoucherList(t *testing.T) {
	c := getClient(t)
	s := &PromotionService{client: c}

	resp, err := s.GetVoucherList(10, 1, "")
	if err != nil {
		t.Logf("GetVoucherList: %v", err)
		return
	}
	t.Logf("Vouchers: %d", len(resp.Response.VoucherList))
}

func TestIntegrationPromotion_GetShopCategoryList(t *testing.T) {
	c := getClient(t)
	s := &PromotionService{client: c}

	resp, err := s.GetShopCategoryList()
	if err != nil {
		t.Logf("GetShopCategoryList: %v", err)
		return
	}
	t.Logf("Shop categories: %d", len(resp.Response.CategoryList))
}

// ---------------------------------------------------------------------------
// Returns
// ---------------------------------------------------------------------------

func TestIntegrationReturns_GetShippingCarrier(t *testing.T) {
	c := getClient(t)
	s := &ReturnsService{client: c}

	resp, err := s.GetShippingCarrier()
	if err != nil {
		t.Logf("GetShippingCarrier: %v", err)
		return
	}
	t.Logf("Carriers: %d", len(resp.Response.CarrierList))
}

func TestIntegrationReturns_GetReturnList(t *testing.T) {
	c := getClient(t)
	s := &ReturnsService{client: c}

	resp, err := s.GetReturnList(10, "", 0, 0, "")
	if err != nil {
		t.Logf("GetReturnList: %v", err)
		return
	}
	t.Logf("Returns: %d, more: %v", len(resp.Response.ReturnList), resp.Response.More)
}

// ---------------------------------------------------------------------------
// Ads
// ---------------------------------------------------------------------------

func TestIntegrationAds_CheckCreateGMSCampaignEligibility(t *testing.T) {
	c := getClient(t)
	s := &AdsService{client: c}

	resp, err := s.CheckCreateGMSCampaignEligibility()
	if err != nil {
		t.Logf("CheckCreateGMSCampaignEligibility: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("GMS eligible: %v", resp.Response.Eligible)
	}
}

func TestIntegrationAds_GetFacilShopRate(t *testing.T) {
	c := getClient(t)
	s := &AdsService{client: c}

	resp, err := s.GetFacilShopRate()
	if err != nil {
		t.Logf("GetFacilShopRate: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("Facil rate: %.4f", resp.Response.Rate)
	}
}

// ---------------------------------------------------------------------------
// Partner
// ---------------------------------------------------------------------------

func TestIntegrationPartner_GetShopeeIPRanges(t *testing.T) {
	c := getClient(t)
	s := &PartnerService{client: c}

	resp, err := s.GetShopeeIPRanges()
	if err != nil {
		t.Logf("GetShopeeIPRanges: %v", err)
		return
	}
	t.Logf("IP ranges: %d", len(resp.Response.IPRangeList))
}

func TestIntegrationPartner_GetShopsByPartner(t *testing.T) {
	c := getClient(t)
	s := &PartnerService{client: c}

	resp, err := s.GetShopsByPartner(10, 1)
	if err != nil {
		t.Logf("GetShopsByPartner: %v", err)
		return
	}
	t.Logf("Partner shops: %d", len(resp.Response.ShopList))
}

// ---------------------------------------------------------------------------
// BR
// ---------------------------------------------------------------------------

func TestIntegrationBR_QueryShopEnrollmentStatus(t *testing.T) {
	c := getClient(t)
	s := &BRService{client: c}

	resp, err := s.QueryShopEnrollmentStatus([]int64{227601935})
	if err != nil {
		t.Logf("QueryShopEnrollmentStatus: %v", err)
		return
	}
	t.Logf("Enrollment results: %d", len(resp.Response.ShopList))
}

// ---------------------------------------------------------------------------
// WHS
// ---------------------------------------------------------------------------

func TestIntegrationWHS_GetCurrentInventory(t *testing.T) {
	c := getClient(t)
	s := &WHSService{client: c}

	resp, err := s.GetCurrentInventory(0, nil, 10, 1)
	if err != nil {
		t.Logf("GetCurrentInventory: %v", err)
		return
	}
	t.Logf("Inventory items: %d", len(resp.Response.InventoryList))
}

// ---------------------------------------------------------------------------
// Live
// ---------------------------------------------------------------------------

func TestIntegrationLive_GetSessionDetail(t *testing.T) {
	c := getClient(t)
	s := &LiveService{client: c}

	resp, err := s.GetSessionDetail(1)
	if err != nil {
		t.Logf("GetSessionDetail: %v", err)
		return
	}
	if resp.Response != nil {
		t.Logf("Session: %s (status: %s)", resp.Response.Name, resp.Response.Status)
	}
}
