package shopee

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestGenerateSignature(t *testing.T) {
	partnerKey := "test-key"
	partnerID := int64(12345)
	apiPath := "/api/v2/product/get_item_list"
	timestamp := int64(1700000000)
	accessToken := "test-token"
	shopID := int64(67890)

	sig := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, shopID, 0)

	// Recompute manually to verify
	parts := make([]byte, 0, 256)
	parts = strconv.AppendInt(parts, partnerID, 10)
	parts = append(parts, apiPath...)
	parts = strconv.AppendInt(parts, timestamp, 10)
	parts = append(parts, accessToken...)
	parts = strconv.AppendInt(parts, shopID, 10)
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write(parts)
	expected := hex.EncodeToString(mac.Sum(nil))

	if sig != expected {
		t.Fatalf("signature mismatch:\n  got:  %s\n  want: %s", sig, expected)
	}
}

func TestGenerateSignatureNoAccessToken(t *testing.T) {
	partnerKey := "test-key"
	partnerID := int64(12345)
	apiPath := "/api/v2/partner/get_shopee_ip_ranges"
	timestamp := int64(1700000000)

	sig := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, "", 0, 0)

	parts := make([]byte, 0, 256)
	parts = strconv.AppendInt(parts, partnerID, 10)
	parts = append(parts, apiPath...)
	parts = strconv.AppendInt(parts, timestamp, 10)
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write(parts)
	expected := hex.EncodeToString(mac.Sum(nil))

	if sig != expected {
		t.Fatalf("signature mismatch:\n  got:  %s\n  want: %s", sig, expected)
	}
}

func TestGenerateSignatureNoShopID(t *testing.T) {
	partnerKey := "test-key"
	partnerID := int64(12345)
	apiPath := "/api/v2/auth/token/get"
	timestamp := int64(1700000000)
	accessToken := "test-token"

	sig := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, 0, 0)

	parts := make([]byte, 0, 256)
	parts = strconv.AppendInt(parts, partnerID, 10)
	parts = append(parts, apiPath...)
	parts = strconv.AppendInt(parts, timestamp, 10)
	parts = append(parts, accessToken...)
	mac := hmac.New(sha256.New, []byte(partnerKey))
	mac.Write(parts)
	expected := hex.EncodeToString(mac.Sum(nil))

	if sig != expected {
		t.Fatalf("signature mismatch:\n  got:  %s\n  want: %s", sig, expected)
	}
}

func TestGenerateSignatureConsistency(t *testing.T) {
	partnerKey := "shpk4372466c416b4b726a59676d794a74437649416b5759614351675948764b"
	partnerID := int64(1235051)
	apiPath := "/api/v2/shop/get_shop_info"
	timestamp := int64(1700000000)
	accessToken := "test-access-token"
	shopID := int64(12345)

	sig1 := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, shopID, 0)
	sig2 := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, shopID, 0)

	if sig1 != sig2 {
		t.Fatal("signature is not deterministic")
	}
}

func TestIsMerchantAPI(t *testing.T) {
	merchant := []string{
		"/api/v2/merchant/get_merchant_info",
		"/api/v2/global_product/get_category",
		"/api/v2/first_mile/get_courier_delivery_detail",
	}
	shop := []string{
		"/api/v2/shop/get_shop_info",
		"/api/v2/product/get_item_list",
		"/api/v2/order/get_order_list",
		"/api/v2/media/upload_image",
		"/api/v2/auth/token/get",
	}
	for _, p := range merchant {
		if !IsMerchantAPI(p) {
			t.Errorf("isMerchantAPI(%q) = false, want true", p)
		}
	}
	for _, p := range shop {
		if IsMerchantAPI(p) {
			t.Errorf("isMerchantAPI(%q) = true, want false", p)
		}
	}
}

func TestBaseQueryDomain(t *testing.T) {
	c := &Client{
		PartnerID:  12345,
		PartnerKey: "test-key",
		AccessToken: "token",
		ShopID:     67890,
		MerchantID: 11111,
	}

	shopQ := c.baseQuery("/api/v2/shop/get_shop_info", 1700000000)
	if shopQ.Get("shop_id") != "67890" {
		t.Errorf("shop query missing shop_id: %v", shopQ)
	}
	if shopQ.Get("merchant_id") != "" {
		t.Errorf("shop query must not carry merchant_id: %v", shopQ)
	}

	merQ := c.baseQuery("/api/v2/global_product/get_category", 1700000000)
	if merQ.Get("merchant_id") != "11111" {
		t.Errorf("merchant query missing merchant_id: %v", merQ)
	}
	if merQ.Get("shop_id") != "" {
		t.Errorf("merchant query must not carry shop_id: %v", merQ)
	}
}
