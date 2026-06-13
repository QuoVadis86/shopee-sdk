package shopee

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEscrowDetailParsesProfitFees(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("order_sn"); got != "ORDER-1" {
			t.Errorf("order_sn = %q", got)
		}
		_, _ = w.Write([]byte(`{
			"response":{
				"order_sn":"ORDER-1",
				"buyer_user_name":"buyer",
				"order_income":{
					"escrow_amount":90,
					"buyer_total_amount":100,
					"seller_discount":5,
					"shopee_discount":3,
					"voucher_from_seller":2,
					"commission_fee":4,
					"service_fee":1.5,
					"seller_transaction_fee":2.5,
					"order_ams_commission_fee":1.25,
					"actual_shipping_fee":6,
					"items":[{
						"item_id":10,
						"model_id":20,
						"quantity_purchased":2,
						"ams_commission_fee":0.75
					}]
				}
			}
		}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "token", 2, WithBaseURL(server.URL))
	result, err := sdk.Payment.GetEscrowDetail("ORDER-1")
	if err != nil {
		t.Fatalf("GetEscrowDetail() error = %v", err)
	}
	income := result.Response.OrderIncome
	if income.CommissionFee != 4 || income.SellerTransactionFee != 2.5 ||
		income.OrderAMSCommissionFee != 1.25 || income.Items[0].AMSCommissionFee != 0.75 {
		t.Fatalf("income = %+v", income)
	}
}

func TestGetEscrowDetailBatchUsesTypedResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{
			"response":[{
				"escrow_detail":{
					"order_sn":"ORDER-1",
					"order_income":{"commission_fee":4}
				}
			}]
		}`))
	}))
	defer server.Close()

	sdk := NewSDK(1, "key", "token", 2, WithBaseURL(server.URL))
	result, err := sdk.Payment.GetEscrowDetailBatch([]string{"ORDER-1"})
	if err != nil {
		t.Fatalf("GetEscrowDetailBatch() error = %v", err)
	}
	if result.Response[0].EscrowDetail.OrderIncome.CommissionFee != 4 {
		t.Fatalf("response = %+v", result.Response)
	}
}
