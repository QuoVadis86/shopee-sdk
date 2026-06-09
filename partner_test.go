package shopee

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestPartnerRefreshAccessTokenUsesPublicAuthSignature(t *testing.T) {
	const (
		partnerID    = int64(123456)
		partnerKey   = "partner-key"
		shopID       = int64(987654)
		refreshToken = "old-refresh-token"
	)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != PathPartnerRefreshAccessToken {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("access_token"); got != "" {
			t.Fatalf("refresh request must not include access_token: %s", got)
		}
		if got := r.URL.Query().Get("shop_id"); got != "" {
			t.Fatalf("refresh request must not include shop_id in query: %s", got)
		}

		timestamp, err := strconv.ParseInt(r.URL.Query().Get("timestamp"), 10, 64)
		if err != nil {
			t.Fatalf("invalid timestamp: %v", err)
		}
		wantSign := GenerateSignature(partnerKey, partnerID, PathPartnerRefreshAccessToken, timestamp, "", 0)
		if got := r.URL.Query().Get("sign"); got != wantSign {
			t.Fatalf("unexpected signature: got %s want %s", got, wantSign)
		}

		var body RefreshAccessTokenParams
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.PartnerID != partnerID || body.ShopID != shopID || body.RefreshToken != refreshToken {
			t.Fatalf("unexpected body: %+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"new-access","refresh_token":"new-refresh","expire_in":14400,"shop_id":987654}`))
	}))
	defer server.Close()

	sdk := NewSDK(
		partnerID,
		partnerKey,
		"expired-access-token",
		shopID,
		WithBaseURL(server.URL),
		WithHTTPClient(server.Client()),
	)
	resp, err := sdk.Partner.RefreshAccessToken(&RefreshAccessTokenParams{
		RefreshToken: refreshToken,
		PartnerID:    partnerID,
		ShopID:       shopID,
	})
	if err != nil {
		t.Fatalf("RefreshAccessToken failed: %v", err)
	}
	if resp.AccessToken != "new-access" || resp.RefreshToken != "new-refresh" || resp.ExpireIn != 14400 {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestPartnerRefreshAccessTokenSupportsMerchantAndNestedResponse(t *testing.T) {
	const merchantID = int64(42)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body RefreshAccessTokenParams
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("decode body: %v", err)
		}
		if body.MerchantID != merchantID || body.ShopID != 0 {
			t.Fatalf("unexpected merchant refresh body: %+v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"response":{"access_token":"nested-access","refresh_token":"nested-refresh","expire_in":7200,"merchant_id":42}}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "", 0, WithBaseURL(server.URL), WithHTTPClient(server.Client()))
	resp, err := sdk.Partner.RefreshAccessToken(&RefreshAccessTokenParams{
		RefreshToken: "merchant-refresh",
		PartnerID:    1,
		MerchantID:   merchantID,
	})
	if err != nil {
		t.Fatalf("RefreshAccessToken failed: %v", err)
	}
	if resp.AccessToken != "nested-access" || resp.MerchantID != merchantID {
		t.Fatalf("unexpected response: %+v", resp)
	}
}

func TestPartnerRefreshAccessTokenReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"error":"error_auth","message":"Invalid refresh_token.","request_id":"request-1"}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "", 0, WithBaseURL(server.URL), WithHTTPClient(server.Client()))
	_, err := sdk.Partner.RefreshAccessToken(&RefreshAccessTokenParams{
		RefreshToken: "invalid",
		PartnerID:    1,
		ShopID:       2,
	})
	if err == nil {
		t.Fatal("expected API error")
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected APIError, got %T: %v", err, err)
	}
	if apiErr.ErrorCode != "error_auth" || apiErr.RequestID != "request-1" {
		t.Fatalf("unexpected API error: %+v", apiErr)
	}
}
