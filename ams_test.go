package shopee

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOpenCampaignAddedProductUsesDocumentedParametersAndParsesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != PathAMSGetOpenCampaignAddedProduct {
			t.Fatalf("path = %q, want %q", r.URL.Path, PathAMSGetOpenCampaignAddedProduct)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"page_size":      "20",
			"cursor":         "1234,5678",
			"sort_by":        "-commission_rate",
			"search_type":    "ITEM_ID",
			"search_content": "123,456",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Errorf("%s = %q, want %q", key, got, want)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"error": "",
			"message": "",
			"request_id": "request-1",
			"response": {
				"item_list": [{
					"item_id": 123,
					"item_name": "test",
					"campaign_id": 456,
					"campaign_status": "Ongoing",
					"commission_rate": 1.11,
					"period_start_time": 1735660800,
					"period_end_time": 32503651199,
					"pending_terminated_time": 1735660900,
					"commission_protection_list": [{
						"commission_rate": 1.21,
						"protection_period_end_time": 1735661000
					}],
					"max_commission_rate_current_day": 1.21
				}],
				"total_count": 1000,
				"cursor": "next-cursor",
				"has_more": true
			}
		}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "token", 2, WithBaseURL(server.URL))
	result, err := sdk.AMS.GetOpenCampaignAddedProduct(GetOpenCampaignAddedProductRequest{
		PageSize:      20,
		Cursor:        "1234,5678",
		SortBy:        "-commission_rate",
		SearchType:    "ITEM_ID",
		SearchContent: "123,456",
	})
	if err != nil {
		t.Fatalf("GetOpenCampaignAddedProduct() error = %v", err)
	}
	if result.Response == nil {
		t.Fatal("response is nil")
	}
	if result.Response.TotalCount != 1000 || result.Response.Cursor != "next-cursor" || !result.Response.HasMore {
		t.Fatalf("pagination = %+v", result.Response)
	}
	item := result.Response.ItemList[0]
	if item.CampaignID != 456 || item.CommissionRate != 1.11 || item.PeriodEndTime != 32503651199 {
		t.Fatalf("item = %+v", item)
	}
	if len(item.CommissionProtectionList) != 1 ||
		item.CommissionProtectionList[0].ProtectionPeriodEndTime != 1735661000 {
		t.Fatalf("commission protection = %+v", item.CommissionProtectionList)
	}
}

func TestGetOpenCampaignAddedProductOmitsOptionalParameters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		for _, key := range []string{"cursor", "sort_by", "search_type", "search_content"} {
			if _, exists := query[key]; exists {
				t.Errorf("optional query %q should be omitted", key)
			}
		}
		_, _ = w.Write([]byte(`{"response":{"item_list":[],"total_count":0,"cursor":"","has_more":false}}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "token", 2, WithBaseURL(server.URL))
	_, err := sdk.AMS.GetOpenCampaignAddedProduct(GetOpenCampaignAddedProductRequest{PageSize: 100})
	if err != nil {
		t.Fatalf("GetOpenCampaignAddedProduct() error = %v", err)
	}
}
