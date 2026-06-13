package main

import (
	"fmt"
	"log"
	"os"

	"github.com/QuoVadis86/shopee-sdk"
)

func main() {
	partnerID := int64(123456)
	partnerKey := "your-partner-key"
	accessToken := "your-access-token"
	shopID := int64(12345)

	sdk := shopee.NewSDK(partnerID, partnerKey, accessToken, shopID, shopee.WithRegion(shopee.RegionSandbox))

	shopInfo, err := sdk.Shop.GetShopInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetShopInfo error: %v\n", err)
	} else if shopInfo.HasError() {
		fmt.Fprintf(os.Stderr, "API error: %s - %s\n", shopInfo.Error, shopInfo.Message)
	} else {
		log.Printf("Shop: %s (Region: %s)", shopInfo.ShopName, shopInfo.Region)
	}

	categories, err := sdk.Product.GetCategory("en")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetCategory error: %v\n", err)
	} else {
		log.Printf("Found %d categories", len(categories.Response.CategoryList))
	}

	items, err := sdk.Product.GetItemList(0, 10, 0, 0, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetItemList error: %v\n", err)
	} else {
		log.Printf("Items: %d total, hasNextPage: %v", items.Response.TotalCount, items.Response.HasNextPage)
	}

	orders, err := sdk.Order.GetOrderList("create_time", 0, 0, 10, "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetOrderList error: %v\n", err)
	} else {
		log.Printf("Orders: %d, more: %v", len(orders.Response.OrderList), orders.Response.More)
	}

	channels, err := sdk.Logistics.GetChannelList()
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetChannelList error: %v\n", err)
	} else {
		log.Printf("Logistics channels: %d", len(channels.Response.LogisticsChannelList))
	}

	addresses, err := sdk.Logistics.GetAddressList()
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetAddressList error: %v\n", err)
	} else {
		log.Printf("Addresses: %d", len(addresses.Response.AddressList))
	}

	amsProducts, err := sdk.AMS.GetOpenCampaignAddedProduct(10, "", "", "", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetOpenCampaignAddedProduct error: %v\n", err)
	} else if amsProducts.Response != nil {
		log.Printf("AMS products: %d total, has_more=%v", amsProducts.Response.TotalCount, amsProducts.Response.HasMore)
	}

	paymentMethods, err := sdk.Payment.GetPaymentMethodList()
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetPaymentMethodList error: %v\n", err)
	} else {
		log.Printf("Payment methods: %d", len(paymentMethods.Response.PaymentMethodList))
	}

	discounts, err := sdk.Promotion.GetDiscountList(10, 1, "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "GetDiscountList error: %v\n", err)
	} else {
		log.Printf("Discounts: %d", len(discounts.Response.DiscountList))
	}

	sdk.SetAccessToken("new-access-token")
	log.Println("Access token updated")
}
