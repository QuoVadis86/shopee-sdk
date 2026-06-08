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

	sig := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, shopID)

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

	sig := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, "", 0)

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
	apiPath := "/api/v2/auth/get_access_token"
	timestamp := int64(1700000000)
	accessToken := "test-token"

	sig := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, 0)

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

	sig1 := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, shopID)
	sig2 := GenerateSignature(partnerKey, partnerID, apiPath, timestamp, accessToken, shopID)

	if sig1 != sig2 {
		t.Fatal("signature is not deterministic")
	}
}
