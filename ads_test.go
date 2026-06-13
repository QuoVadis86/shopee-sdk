package shopee

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCPCDailyPerformanceUsesDocumentedDatesAndParsesDailyMetrics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if got := query.Get("start_date"); got != "09-06-2026" {
			t.Errorf("start_date = %q", got)
		}
		if got := query.Get("end_date"); got != "09-06-2026" {
			t.Errorf("end_date = %q", got)
		}
		if query.Get("date_from") != "" || query.Get("date_to") != "" {
			t.Error("legacy date_from/date_to should not be sent")
		}
		_, _ = w.Write([]byte(`{
			"request_id":"request-1",
			"response":[{
				"date":"09-06-2026",
				"impression":123,
				"clicks":12,
				"ctr":9.76,
				"direct_order":2,
				"broad_order":3,
				"direct_gmv":100.5,
				"broad_gmv":120.5,
				"expense":10.25,
				"direct_roas":9.8,
				"broad_roas":11.75
			}]
		}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "token", 2, WithBaseURL(server.URL))
	result, err := sdk.Ads.GetAllCPCDailyPerformance("09-06-2026", "09-06-2026")
	if err != nil {
		t.Fatalf("GetAllCPCDailyPerformance() error = %v", err)
	}
	if len(result.Response) != 1 {
		t.Fatalf("response length = %d", len(result.Response))
	}
	metric := result.Response[0]
	if metric.Clicks != 12 || metric.Expense != 10.25 || metric.BroadGMV != 120.5 {
		t.Fatalf("metric = %+v", metric)
	}
}

func TestGetProductCampaignDailyPerformanceUsesCampaignListAndParsesMetrics(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if got := query.Get("campaign_id_list"); got != "11,22" {
			t.Errorf("campaign_id_list = %q", got)
		}
		_, _ = w.Write([]byte(`{
			"response":[{
				"shop_id":2,
				"region":"TH",
				"campaign_list":[{
					"campaign_id":11,
					"ad_type":"auto",
					"campaign_placement":"search",
					"ad_name":"Campaign",
					"metrics_list":[{
						"date":"09-06-2026",
						"impression":100,
						"clicks":10,
						"expense":5.5,
						"broad_gmv":50,
						"broad_order":2
					}]
				}]
			}]
		}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "token", 2, WithBaseURL(server.URL))
	result, err := sdk.Ads.GetProductCampaignDailyPerformance(
		[]int64{11, 22},
		"09-06-2026",
		"09-06-2026",
	)
	if err != nil {
		t.Fatalf("GetProductCampaignDailyPerformance() error = %v", err)
	}
	campaign := result.Response[0].CampaignList[0]
	if campaign.CampaignID != 11 || campaign.MetricsList[0].Expense != 5.5 {
		t.Fatalf("campaign = %+v", campaign)
	}
}
