package shopee

import (
	"os"
	"strconv"
	"testing"
)

func getTestCredentials(t *testing.T) (partnerID int64, partnerKey, accessToken string, shopID int64, ok bool) {
	t.Helper()

	partnerIDStr := os.Getenv("SHOPEE_PARTNER_ID")
	partnerKey = os.Getenv("SHOPEE_PARTNER_KEY")

	if partnerIDStr == "" || partnerKey == "" {
		return 0, "", "", 0, false
	}

	id, err := strconv.ParseInt(partnerIDStr, 10, 64)
	if err != nil {
		t.Fatalf("invalid SHOPEE_PARTNER_ID: %v", err)
	}

	accessToken = os.Getenv("SHOPEE_ACCESS_TOKEN")
	shopIDStr := os.Getenv("SHOPEE_SHOP_ID")
	if shopIDStr != "" {
		sid, err := strconv.ParseInt(shopIDStr, 10, 64)
		if err != nil {
			t.Fatalf("invalid SHOPEE_SHOP_ID: %v", err)
		}
		shopID = sid
	}

	return id, partnerKey, accessToken, shopID, true
}

func TestIntegrationGetShopeeIPRanges(t *testing.T) {
	partnerID, partnerKey, accessToken, shopID, ok := getTestCredentials(t)
	if !ok {
		t.Skip("set SHOPEE_PARTNER_ID and SHOPEE_PARTNER_KEY env vars to run integration tests")
	}

	client := NewClient(partnerID, partnerKey, accessToken, shopID, WithRegion(RegionSandbox))

	result := &GetShopeeIPRangesResponse{}
	err := client.DoGet(PathPartnerGetShopeeIPRanges, map[string]string{}, result)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Logf("IP ranges response: request_id=%s error=%s message=%s",
		result.RequestID, result.Error, result.Message)

	if result.HasError() {
		t.Logf("API returned error (expected without full auth): %s - %s", result.Error, result.Message)
	} else if len(result.Response.IPRangeList) > 0 {
		t.Logf("Got %d IP ranges, first: %s", len(result.Response.IPRangeList), result.Response.IPRangeList[0].IPRange)
	}
}

func TestIntegrationGetShopInfo(t *testing.T) {
	partnerID, partnerKey, accessToken, shopID, ok := getTestCredentials(t)
	if !ok {
		t.Skip("set SHOPEE_PARTNER_ID and SHOPEE_PARTNER_KEY env vars to run integration tests")
	}

	client := NewClient(partnerID, partnerKey, accessToken, shopID, WithRegion(RegionSandbox))

	result := &GetShopInfoResponse{}
	err := client.DoGet(PathShopGetInfo, map[string]string{}, result)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Logf("Shop info response: request_id=%s error=%s message=%s",
		result.RequestID, result.Error, result.Message)

	if result.HasError() {
		t.Logf("API returned error (expected without auth): %s - %s", result.Error, result.Message)
	} else if result.ShopName != "" {
		t.Logf("Shop: %s (Region: %s, Status: %s)",
			result.ShopName, result.Region, result.Status)
	}
}

func TestIntegrationGetCategory(t *testing.T) {
	partnerID, partnerKey, accessToken, shopID, ok := getTestCredentials(t)
	if !ok {
		t.Skip("set SHOPEE_PARTNER_ID and SHOPEE_PARTNER_KEY env vars to run integration tests")
	}

	client := NewClient(partnerID, partnerKey, accessToken, shopID, WithRegion(RegionSandbox))

	result := &GetCategoryResponse{}
	err := client.DoGet(PathProductGetCategory, map[string]string{"language": "en"}, result)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Logf("Category response: request_id=%s error=%s message=%s category_count=%d",
		result.RequestID, result.Error, result.Message, len(result.Response.CategoryList))

	if result.HasError() {
		t.Logf("API returned error: %s - %s", result.Error, result.Message)
	} else if len(result.Response.CategoryList) > 0 {
		t.Logf("First category: %s (ID: %d)", result.Response.CategoryList[0].CategoryName, result.Response.CategoryList[0].CategoryID)
	}
}


