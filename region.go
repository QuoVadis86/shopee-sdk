package shopee

import "strings"

type Region string

const (
	RegionGlobal    Region = "GLOBAL"
	RegionCN        Region = "CN"
	RegionBrazil    Region = "BRAZIL"
	RegionSandbox   Region = "SANDBOX"
	RegionSandboxCN Region = "SANDBOX_CN"
)

var BaseURLs = map[Region]string{
	RegionGlobal:    "https://partner.shopeemobile.com",
	RegionCN:        "https://openplatform.shopee.cn",
	RegionBrazil:    "https://openplatform.shopee.com.br",
	RegionSandbox:   "https://openplatform.sandbox.test-stable.shopee.sg",
	RegionSandboxCN: "https://openplatform.sandbox.test-stable.shopee.cn",
}

var AuthURLs = map[Region]string{
	RegionGlobal:    "https://partner.shopeemobile.com/api/v2/shop/auth_partner",
	RegionCN:        "https://openplatform.shopee.cn/api/v2/shop/auth_partner",
	RegionBrazil:    "https://openplatform.shopee.com.br/api/v2/shop/auth_partner",
	RegionSandbox:   "https://openplatform.sandbox.test-stable.shopee.sg/api/v2/shop/auth_partner",
	RegionSandboxCN: "https://openplatform.sandbox.test-stable.shopee.cn/api/v2/shop/auth_partner",
}

const apiPrefix = "/api/v2"

type apiPaths struct {
	prefix string
}

func newAPIPaths(service string) *apiPaths {
	return &apiPaths{prefix: apiPrefix + "/" + service}
}

func (p *apiPaths) P(endpoint string) string {
	return p.prefix + "/" + endpoint
}

// APIPath returns a full API path for a given service and endpoint.
// Example: APIPath("product", "get_category") -> "/api/v2/product/get_category"
// Some endpoint values may already contain dots; they are preserved as-is.
func APIPath(service, endpoint string) string {
	return apiPrefix + "/" + service + "/" + endpoint
}

// ---------------------------------------------------------------------------
// Product (Local)
// ---------------------------------------------------------------------------

var (
	PathProductGetCategory         = APIPath("product", "get_category")
	PathProductGetAttributeTree    = APIPath("product", "get_attribute_tree")
	PathProductGetBrandList        = APIPath("product", "get_brand_list")
	PathProductGetItemLimit        = APIPath("product", "get_item_limit")
	PathProductGetItemList         = APIPath("product", "get_item_list")
	PathProductGetItemBaseInfo     = APIPath("product", "get_item_base_info")
	PathProductGetItemExtraInfo    = APIPath("product", "get_item_extra_info")
	PathProductAddItem             = APIPath("product", "add_item")
	PathProductUpdateItem          = APIPath("product", "update_item")
	PathProductDeleteItem          = APIPath("product", "delete_item")
	PathProductUpdateTierVariation = APIPath("product", "update_tier_variation")
	PathProductGetModelList        = APIPath("product", "get_model_list")
	PathProductAddModel            = APIPath("product", "add_model")
	PathProductUpdateModel         = APIPath("product", "update_model")
	PathProductDeleteModel         = APIPath("product", "delete_model")
	PathProductUpdatePrice         = APIPath("product", "update_price")
	PathProductUpdateStock         = APIPath("product", "update_stock")
	PathProductGetBoostedList      = APIPath("product", "get_boosted_list")
	PathProductGetItemPromotion    = APIPath("product", "get_item_promotion")
	PathProductUpdateSIPItemPrice  = APIPath("product", "update_sip_item_price")
	PathProductGetComment          = APIPath("product", "get_comment")
	PathProductGetRecommendAttr    = APIPath("product", "get_recommend_attribute")
	PathProductGetWeightRec        = APIPath("product", "get_weight_recommendation")
	PathProductGetSizeChartList    = APIPath("product", "get_size_chart_list")
	PathProductGetSizeChartDetail  = APIPath("product", "get_size_chart_detail")
	PathProductGetVariations       = APIPath("product", "get_variations")
	PathProductGetAllVehicleList   = APIPath("product", "get_all_vehicle_list")
	PathProductGetVehicleCompList  = APIPath("product", "get_vehicle_list_by_compatibility_detail")
	PathProductGetContentDiagResult = APIPath("product", "get_item_content_diagnosis_result")
	PathProductGetContentDiagList  = APIPath("product", "get_item_list_by_content_diagnosis")
	PathProductGetKitLimit         = APIPath("product", "get_kit_item_limit")
	PathProductAddKitItem          = APIPath("product", "add_kit_item")
	PathProductUpdateKitItem       = APIPath("product", "update_kit_item")
	PathProductGetKitItemInfo      = APIPath("product", "get_kit_item_info")
	PathProductGetSSPList          = APIPath("product", "get_ssp_list")
	PathProductGetSSPInfo          = APIPath("product", "get_ssp_info")
	PathProductAddSSPItem          = APIPath("product", "add_ssp_item")
	PathProductGetAItemByPItemID   = APIPath("product", "get_aitem_by_pitem_id")
	PathProductSearchAttrValue     = APIPath("product", "search_attribute_value_list")
	PathProductGetMainItemList     = APIPath("product", "get_main_item_list")
	PathProductGetDirectItemList   = APIPath("product", "get_direct_item_list")
	PathProductGetDirectShopPrice  = APIPath("product", "get_direct_shop_recommended_price")
	PathProductGetCertRule         = APIPath("product", "get_product_certification_rule")
	PathProductPublishToOutlet     = APIPath("product", "publish_item_to_outlet_shop")
	PathProductGetMartItemMapping  = APIPath("product", "get_mart_item_mapping_by_id")
	PathProductSearchUnpackaged    = APIPath("product", "search_unpackaged_model_list")
	PathProductGetMartItemByOutlet = APIPath("product", "get_mart_item_by_outlet_item_id")
	PathProductGetItemViolation    = APIPath("product", "get_item_violation_info")
	PathProductBoostItem           = APIPath("product", "boost_item")
	PathProductUnlistItem          = APIPath("product", "unlist_item")
)

// ---------------------------------------------------------------------------
// Global Product (Cross-border)
// ---------------------------------------------------------------------------

var (
	PathGlobalProductGetCategory         = APIPath("global_product", "get_category")
	PathGlobalProductGetAttributeTree    = APIPath("global_product", "get_attribute_tree")
	PathGlobalProductGetBrandList        = APIPath("global_product", "get_brand_list")
	PathGlobalProductGetItemLimit        = APIPath("global_product", "get_global_item_limit")
	PathGlobalProductGetItemList         = APIPath("global_product", "get_global_item_list")
	PathGlobalProductGetItemInfo         = APIPath("global_product", "get_global_item_info")
	PathGlobalProductAddItem             = APIPath("global_product", "add_global_item")
	PathGlobalProductUpdateItem          = APIPath("global_product", "update_global_item")
	PathGlobalProductDeleteItem          = APIPath("global_product", "delete_global_item")
	PathGlobalProductUpdateTierVariation = APIPath("global_product", "update_tier_variation")
	PathGlobalProductAddModel            = APIPath("global_product", "add_global_model")
	PathGlobalProductUpdateModel         = APIPath("global_product", "update_global_model")
	PathGlobalProductDeleteModel         = APIPath("global_product", "delete_global_model")
	PathGlobalProductGetModelList        = APIPath("global_product", "get_global_model_list")
	PathGlobalProductUpdateSizeChart     = APIPath("global_product", "update_size_chart")
	PathGlobalProductCreatePublishTask   = APIPath("global_product", "create_publish_task")
	PathGlobalProductGetPublishableShop  = APIPath("global_product", "get_publishable_shop")
	PathGlobalProductGetPublishTaskResult = APIPath("global_product", "get_publish_task_result")
	PathGlobalProductGetPublishedList    = APIPath("global_product", "get_published_list")
	PathGlobalProductUpdatePrice         = APIPath("global_product", "update_price")
	PathGlobalProductUpdateStock         = APIPath("global_product", "update_stock")
	PathGlobalProductSetSyncField        = APIPath("global_product", "set_sync_field")
	PathGlobalProductGetGlobalItemID     = APIPath("global_product", "get_global_item_id")
	PathGlobalProductGetRecommendAttr    = APIPath("global_product", "get_recommend_attribute")
	PathGlobalProductGetShopPublishable  = APIPath("global_product", "get_shop_publishable_status")
	PathGlobalProductGetVariations       = APIPath("global_product", "get_variations")
	PathGlobalProductGetSizeChartList    = APIPath("global_product", "get_size_chart_list")
	PathGlobalProductGetSizeChartDetail  = APIPath("global_product", "get_size_chart_detail")
	PathGlobalProductSearchAttrValue     = APIPath("global_product", "search_global_attribute_value_list")
	PathGlobalProductGetLocalAdjRate     = APIPath("global_product", "get_local_adjustment_rate")
	PathGlobalProductUpdateLocalAdjRate  = APIPath("global_product", "update_local_adjustment_rate")
	PathGlobalProductGetVideoUploadResult = APIPath("media_space", "get_video_upload_result")
	PathGlobalProductCancelVideoUpload   = APIPath("media_space", "cancel_video_upload")
)

// ---------------------------------------------------------------------------
// Order
// ---------------------------------------------------------------------------

var (
	PathOrderGetOrderList            = APIPath("order", "get_order_list")
	PathOrderGetOrderDetail          = APIPath("order", "get_order_detail")
	PathOrderGetShipmentList         = APIPath("order", "get_shipment_list")
	PathOrderSearchPackageList       = APIPath("order", "search_package_list")
	PathOrderGetPackageDetail        = APIPath("order", "get_package_detail")
	PathOrderSplitOrder              = APIPath("order", "split_order")
	PathOrderUnsplitOrder            = APIPath("order", "unsplit_order")
	PathOrderCancelOrder             = APIPath("order", "cancel_order")
	PathOrderSetNote                 = APIPath("order", "set_note")
	PathOrderGetPendingInvoiceList   = APIPath("order", "get_pending_buyer_invoice_order_list")
	PathOrderGetBuyerInvoiceInfo     = APIPath("order", "get_buyer_invoice_info")
	PathOrderGetWarehouseFilterCfg   = APIPath("order", "get_warehouse_filter_config")
	PathOrderGetBookingList          = APIPath("order", "get_booking_list")
	PathOrderGetBookingDetail        = APIPath("order", "get_booking_detail")
	PathOrderGetFBSInvoicesResult    = APIPath("order", "get_fbs_invoices_result")
	PathOrderGetEstimateCancelValue  = APIPath("order", "get_estimate_cancel_value")
)

// ---------------------------------------------------------------------------
// Logistics
// ---------------------------------------------------------------------------

var (
	PathLogisticsGetShippingParam         = APIPath("logistics", "get_shipping_parameter")
	PathLogisticsGetMassShippingParam     = APIPath("logistics", "get_mass_shipping_parameter")
	PathLogisticsShipOrder                = APIPath("logistics", "ship_order")
	PathLogisticsMassShipOrder            = APIPath("logistics", "mass_ship_order")
	PathLogisticsUpdateShippingOrder      = APIPath("logistics", "update_shipping_order")
	PathLogisticsGetTrackingNumber        = APIPath("logistics", "get_tracking_number")
	PathLogisticsGetMassTrackingNumber    = APIPath("logistics", "get_mass_tracking_number")
	PathLogisticsGetShippingDocParam      = APIPath("logistics", "get_shipping_document_parameter")
	PathLogisticsCreateShippingDoc        = APIPath("logistics", "create_shipping_document")
	PathLogisticsGetShippingDocResult     = APIPath("logistics", "get_shipping_document_result")
	PathLogisticsGetShippingDocDataInfo   = APIPath("logistics", "get_shipping_document_data_info")
	PathLogisticsGetTrackingInfo          = APIPath("logistics", "get_tracking_info")
	PathLogisticsGetAddressList           = APIPath("logistics", "get_address_list")
	PathLogisticsSetAddressConfig         = APIPath("logistics", "set_address_config")
	PathLogisticsUpdateAddress            = APIPath("logistics", "update_address")
	PathLogisticsDeleteAddress            = APIPath("logistics", "delete_address")
	PathLogisticsGetChannelList           = APIPath("logistics", "get_channel_list")
	PathLogisticsUpdateChannel            = APIPath("logistics", "update_channel")
	PathLogisticsGetOperatingHours        = APIPath("logistics", "get_operating_hours")
	PathLogisticsGetOpHoursRestrictions   = APIPath("logistics", "get_operating_hour_restrictions")
	PathLogisticsUpdateOperatingHours     = APIPath("logistics", "update_operating_hours")
	PathLogisticsDeleteSpecialOpHour      = APIPath("logistics", "delete_special_operating_hour")
	PathLogisticsBatchUpdateTPFTracking   = APIPath("logistics", "batch_update_tpf_warehouse_tracking_status")
	PathLogisticsBatchShipOrder           = APIPath("logistics", "batch_ship_order")
	PathLogisticsUpdateTrackingStatus     = APIPath("logistics", "update_tracking_status")
	PathLogisticsGetBookingShippingParam  = APIPath("logistics", "get_booking_shipping_parameter")
	PathLogisticsShipBooking              = APIPath("logistics", "ship_booking")
	PathLogisticsGetBookingTrackingNum    = APIPath("logistics", "get_booking_tracking_number")
	PathLogisticsGetBookingShipDocParam   = APIPath("logistics", "get_booking_shipping_document_parameter")
	PathLogisticsCreateBookingShipDoc     = APIPath("logistics", "create_booking_shipping_document")
	PathLogisticsGetBookingShipDocResult  = APIPath("logistics", "get_booking_shipping_document_result")
	PathLogisticsGetBookingShipDocData    = APIPath("logistics", "get_booking_shipping_document_data_info")
	PathLogisticsGetBookingTrackingInfo   = APIPath("logistics", "get_booking_tracking_info")
	PathLogisticsCreateShipDocJob         = APIPath("logistics", "create_shipping_document_job")
	PathLogisticsGetShipDocJobStatus      = APIPath("logistics", "get_shipping_document_job_status")
	PathLogisticsUpdateSelfCollectOrder   = APIPath("logistics", "update_self_collection_order_logistics")
	PathLogisticsGetMartPackagingInfo     = APIPath("logistics", "get_mart_packaging_info")
	PathLogisticsSetMartPackagingInfo     = APIPath("logistics", "set_mart_packaging_info")
	PathLogisticsCheckPolygonUpdate       = APIPath("logistics", "check_polygon_update_status")
	PathLogisticsGetPauseStatus           = APIPath("logistics", "get_pause_status")
	PathLogisticsSetPauseStatus           = APIPath("logistics", "set_pause_status")
	PathLogisticsGetUnbindOrderList       = APIPath("first_mile", "get_unbind_order_list")
	PathLogisticsGetDetail                = APIPath("first_mile", "get_detail")
	PathLogisticsGetTrackingNumList       = APIPath("first_mile", "get_tracking_number_list")
	PathLogisticsGetWaybill               = APIPath("first_mile", "get_waybill")
	PathLogisticsGetCourierChannelList    = APIPath("first_mile", "get_courier_delivery_channel_list")
	PathLogisticsGetTransitWarehouseList  = APIPath("first_mile", "get_transit_warehouse_list")
	PathLogisticsGetCourierDeliveryDetail = APIPath("first_mile", "get_courier_delivery_detail")
	PathLogisticsGetCourierDeliveryWaybill = APIPath("first_mile", "get_courier_delivery_waybill")
	PathLogisticsGetCourierTrackingList   = APIPath("first_mile", "get_courier_delivery_tracking_number_list")
)

// ---------------------------------------------------------------------------
// Shop
// ---------------------------------------------------------------------------

var (
	PathShopGetInfo              = APIPath("shop", "get_shop_info")
	PathShopGetProfile           = APIPath("shop", "get_profile")
	PathShopUpdateProfile        = APIPath("shop", "update_profile")
	PathShopGetWarehouseDetail   = APIPath("shop", "get_warehouse_detail")
	PathShopGetNotification      = APIPath("shop", "get_shop_notification")
	PathShopGetAuthResellerBrand = APIPath("shop", "get_authorised_reseller_brand")
	PathShopGetBROnboardingInfo  = APIPath("shop", "get_br_shop_onboarding_info")
	PathShopGetHolidayMode       = APIPath("shop", "get_shop_holiday_mode")
	PathShopSetHolidayMode       = APIPath("shop", "set_shop_holiday_mode")
	PathShopGetMerchantInfo      = APIPath("merchant", "get_merchant_info")
	PathShopGetListByMerchant    = APIPath("merchant", "get_shop_list_by_merchant")
	PathShopGetMerchantWarehouseLocations = APIPath("merchant", "get_merchant_warehouse_location_list")
	PathShopGetMerchantWarehouseList      = APIPath("merchant", "get_merchant_warehouse_list")
	PathShopGetWarehouseEligibleShopList  = APIPath("merchant", "get_warehouse_eligible_shop_list")
	PathShopGetMerchantPrepaidAccountList = APIPath("merchant", "get_merchant_prepaid_account_list")
	PathShopGetPerformance        = APIPath("ams", "get_shop_performance")
	PathShopGetMetricSourceDetail = APIPath("account_health", "get_metric_source_detail")
	PathShopGetPenaltyHistory     = APIPath("account_health", "get_penalty_point_history")
	PathShopGetPunishmentHistory  = APIPath("account_health", "get_punishment_history")
	PathShopGetListingsWithIssues = APIPath("account_health", "get_listings_with_issues")
	PathShopGetLateOrders         = APIPath("account_health", "get_late_orders")
	// PathShopGetTotalBalance removed — Shopee API v2 no longer has this endpoint
	PathShopGetToggleInfo         = APIPath("ads", "get_shop_toggle_info")
)

// ---------------------------------------------------------------------------
// AMS (Sponsored Ads / Marketing)
// ---------------------------------------------------------------------------

var (
	PathAMSGetOpenCampaignAddedProduct        = APIPath("ams", "get_open_campaign_added_product")
	PathAMSGetOpenCampaignNotAddedProduct     = APIPath("ams", "get_open_campaign_not_added_product")
	PathAMSBatchAddProductsToOpenCampaign     = APIPath("ams", "batch_add_products_to_open_campaign")
	PathAMSAddAllProductsToOpenCampaign       = APIPath("ams", "add_all_products_to_open_campaign")
	PathAMSGetAutoAddToggleStatus             = APIPath("ams", "get_auto_add_new_product_toggle_status")
	PathAMSUpdateAutoAddNewProductSetting     = APIPath("ams", "update_auto_add_new_product_setting")
	PathAMSBatchEditProductsOCSetting         = APIPath("ams", "batch_edit_products_open_campaign_setting")
	PathAMSEditAllProductsOCSetting           = APIPath("ams", "edit_all_products_open_campaign_setting")
	PathAMSBatchRemoveProductsOCSetting       = APIPath("ams", "batch_remove_products_open_campaign_setting")
	PathAMSRemoveAllProductsOCSetting         = APIPath("ams", "remove_all_products_open_campaign_setting")
	PathAMSGetOpenCampaignBatchTaskResult     = APIPath("ams", "get_open_campaign_batch_task_result")
	PathAMSGetOptimizationSuggestionProduct   = APIPath("ams", "get_optimization_suggestion_product")
	PathAMSBatchGetProductsSuggestedRate      = APIPath("ams", "batch_get_products_suggested_rate")
	PathAMSGetShopSuggestedRate               = APIPath("ams", "get_shop_suggested_rate")
	PathAMSGetTargetedCampaignAddableProducts = APIPath("ams", "get_targeted_campaign_addable_product_list")
	PathAMSGetRecommendedAffiliateList        = APIPath("ams", "get_recommended_affiliate_list")
	PathAMSGetManagedAffiliateList            = APIPath("ams", "get_managed_affiliate_list")
	PathAMSQueryAffiliateList                 = APIPath("ams", "query_affiliate_list")
	PathAMSCreateTargetedCampaign             = APIPath("ams", "create_new_targeted_campaign")
	PathAMSGetTargetedCampaignList            = APIPath("ams", "get_targeted_campaign_list")
	PathAMSGetTargetedCampaignSettings        = APIPath("ams", "get_targeted_campaign_settings")
	PathAMSUpdateTargetedCampaignBasicInfo    = APIPath("ams", "update_basic_info_of_targeted_campaign")
	PathAMSEditTargetedCampaignProductList    = APIPath("ams", "edit_product_list_of_targeted_campaign")
	PathAMSEditTargetedCampaignAffiliateList  = APIPath("ams", "edit_affiliate_list_of_targeted_campaign")
	PathAMSGetPerfDataUpdateTime              = APIPath("ams", "get_performance_data_update_time")
	PathAMSGetShopPerformance                 = APIPath("ams", "get_shop_performance")
	PathAMSGetProductPerformance              = APIPath("ams", "get_product_performance")
	PathAMSGetAffiliatePerformance            = APIPath("ams", "get_affiliate_performance")
	PathAMSGetContentPerformance              = APIPath("ams", "get_content_performance")
	PathAMSGetCampaignKeyMetricsPerformance   = APIPath("ams", "get_campaign_key_metrics_performance")
	PathAMSGetOpenCampaignPerformance         = APIPath("ams", "get_open_campaign_performance")
	PathAMSGetTargetedCampaignPerformance     = APIPath("ams", "get_targeted_campaign_performance")
	PathAMSGetConversionReport                = APIPath("ams", "get_conversion_report")
	PathAMSGetValidationList                  = APIPath("ams", "get_validation_list")
	PathAMSGetValidationReport                = APIPath("ams", "get_validation_report")
	PathAMSGetCoverList                       = APIPath("video", "get_cover_list")
	PathAMSEditVideoInfo                      = APIPath("video", "edit_video_info")
	PathAMSGetVideoList                       = APIPath("video", "get_video_list")
	PathAMSGetVideoDetail                     = APIPath("video", "get_video_detail")
	PathAMSDeleteVideo                        = APIPath("video", "delete_video")
	PathAMSGetOverviewPerformance             = APIPath("video", "get_overview_performance")
	PathAMSGetMetricTrend                     = APIPath("video", "get_metric_trend")
	PathAMSGetUserDemographics                = APIPath("video", "get_user_demographics")
	PathAMSGetVideoPerfList                   = APIPath("video", "get_video_performance_list")
	PathAMSGetProductPerfList                 = APIPath("ams", "get_product_performance_list")
	PathAMSGetVideoDetailPerf                 = APIPath("video", "get_video_detail_performance")
	PathAMSGetVideoDetailMetricTrend          = APIPath("video", "get_video_detail_metric_trend")
	PathAMSGetVideoDetailAudienceDist         = APIPath("video", "get_video_detail_audience_distribution")
	PathAMSGetVideoDetailProductPerf          = APIPath("video", "get_video_detail_product_performance")
)

// ---------------------------------------------------------------------------
// Payment / Finance
// ---------------------------------------------------------------------------

var (
	PathPaymentGetEscrowDetail          = APIPath("payment", "get_escrow_detail")
	PathPaymentSetShopInstallment       = APIPath("payment", "set_shop_installment_status")
	PathPaymentGetShopInstallment       = APIPath("payment", "get_shop_installment_status")
	PathPaymentGetPayoutDetail          = APIPath("payment", "get_payout_detail")
	PathPaymentSetItemInstallment       = APIPath("payment", "set_item_installment_status")
	PathPaymentGetItemInstallment       = APIPath("payment", "get_item_installment_status")
	PathPaymentGetPaymentMethodList     = APIPath("payment", "get_payment_method_list")
	PathPaymentGetWalletTransactionList = APIPath("payment", "get_wallet_transaction_list")
	PathPaymentGetEscrowList            = APIPath("payment", "get_escrow_list")
	PathPaymentGetPayoutInfo            = APIPath("payment", "get_payout_info")
	PathPaymentGetBillingTransaction    = APIPath("payment", "get_billing_transaction_info")
	PathPaymentGetEscrowDetailBatch     = APIPath("payment", "get_escrow_detail_batch")
	PathPaymentGetIncomeStatement       = APIPath("payment", "get_income_statement")
	PathPaymentGetIncomeReport          = APIPath("payment", "get_income_report")
	PathPaymentGetIncomeOverview        = APIPath("payment", "get_income_overview")
	PathPaymentGetIncomeDetail          = APIPath("payment", "get_income_detail")
)

// ---------------------------------------------------------------------------
// Promotion (Discounts, Bundle Deals, Add-on Deals, Vouchers, Flash Sales)
// ---------------------------------------------------------------------------

var (
	PathPromotionAddDiscount          = APIPath("discount", "add_discount")
	PathPromotionAddDiscountItem      = APIPath("discount", "add_discount_item")
	PathPromotionDeleteDiscount       = APIPath("discount", "delete_discount")
	PathPromotionDeleteDiscountItem   = APIPath("discount", "delete_discount_item")
	PathPromotionGetDiscount          = APIPath("discount", "get_discount")
	PathPromotionGetDiscountList      = APIPath("discount", "get_discount_list")
	PathPromotionUpdateDiscount       = APIPath("discount", "update_discount")
	PathPromotionUpdateDiscountItem   = APIPath("discount", "update_discount_item")
	PathPromotionGetSIPDiscounts      = APIPath("discount", "get_sip_discounts")
	PathPromotionSetSIPDiscount       = APIPath("discount", "set_sip_discount")
	PathPromotionDeleteSIPDiscount    = APIPath("discount", "delete_sip_discount")
	PathPromotionAddBundleDeal        = APIPath("bundle_deal", "add_bundle_deal")
	PathPromotionAddBundleDealItem    = APIPath("bundle_deal", "add_bundle_deal_item")
	PathPromotionGetBundleDealList    = APIPath("bundle_deal", "get_bundle_deal_list")
	PathPromotionGetBundleDeal        = APIPath("bundle_deal", "get_bundle_deal")
	PathPromotionGetBundleDealItem    = APIPath("bundle_deal", "get_bundle_deal_item")
	PathPromotionUpdateBundleDeal     = APIPath("bundle_deal", "update_bundle_deal")
	PathPromotionUpdateBundleDealItem = APIPath("bundle_deal", "update_bundle_deal_item")
	PathPromotionDeleteBundleDeal     = APIPath("bundle_deal", "delete_bundle_deal")
	PathPromotionDeleteBundleDealItem = APIPath("bundle_deal", "delete_bundle_deal_item")
	PathPromotionAddAddOnDeal         = APIPath("add_on_deal", "add_add_on_deal")
	PathPromotionAddAddOnDealMainItem = APIPath("add_on_deal", "add_add_on_deal_main_item")
	PathPromotionAddAddOnDealSubItem  = APIPath("add_on_deal", "add_add_on_deal_sub_item")
	PathPromotionDeleteAddOnDeal      = APIPath("add_on_deal", "delete_add_on_deal")
	PathPromotionDeleteAddOnDealMain  = APIPath("add_on_deal", "delete_add_on_deal_main_item")
	PathPromotionDeleteAddOnDealSub   = APIPath("add_on_deal", "delete_add_on_deal_sub_item")
	PathPromotionGetAddOnDealList     = APIPath("add_on_deal", "get_add_on_deal_list")
	PathPromotionGetAddOnDeal         = APIPath("add_on_deal", "get_add_on_deal")
	PathPromotionGetAddOnDealMainItem = APIPath("add_on_deal", "get_add_on_deal_main_item")
	PathPromotionGetAddOnDealSubItem  = APIPath("add_on_deal", "get_add_on_deal_sub_item")
	PathPromotionUpdateAddOnDeal      = APIPath("add_on_deal", "update_add_on_deal")
	PathPromotionUpdateAddOnDealMain  = APIPath("add_on_deal", "update_add_on_deal_main_item")
	PathPromotionUpdateAddOnDealSub   = APIPath("add_on_deal", "update_add_on_deal_sub_item")
	PathPromotionEndAddOnDeal         = APIPath("add_on_deal", "end_add_on_deal")
	PathPromotionAddVoucher           = APIPath("voucher", "add_voucher")
	PathPromotionDeleteVoucher        = APIPath("voucher", "delete_voucher")
	PathPromotionUpdateVoucher        = APIPath("voucher", "update_voucher")
	PathPromotionGetVoucher           = APIPath("voucher", "get_voucher")
	PathPromotionGetVoucherList       = APIPath("voucher", "get_voucher_list")
	PathPromotionGetTimeSlotID        = APIPath("shop_flash_sale", "get_time_slot_id")
	PathPromotionCreateFlashSale        = APIPath("shop_flash_sale", "create_shop_flash_sale")
	PathPromotionGetItemCriteria        = APIPath("shop_flash_sale", "get_item_criteria")
	PathPromotionAddFlashSaleItems      = APIPath("shop_flash_sale", "add_shop_flash_sale_items")
	PathPromotionGetFlashSaleList       = APIPath("shop_flash_sale", "get_shop_flash_sale_list")
	PathPromotionGetFlashSale           = APIPath("shop_flash_sale", "get_shop_flash_sale")
	PathPromotionGetFlashSaleItems      = APIPath("shop_flash_sale", "get_shop_flash_sale_items")
	PathPromotionUpdateFlashSale        = APIPath("shop_flash_sale", "update_shop_flash_sale")
	PathPromotionUpdateFlashSaleItems   = APIPath("shop_flash_sale", "update_shop_flash_sale_items")
	PathPromotionDeleteFlashSale        = APIPath("shop_flash_sale", "delete_shop_flash_sale")
	PathPromotionDeleteFlashSaleItems   = APIPath("shop_flash_sale", "delete_shop_flash_sale_items")
	PathPromotionAddFollowPrize       = APIPath("follow_prize", "add_follow_prize")
	PathPromotionDeleteFollowPrize    = APIPath("follow_prize", "delete_follow_prize")
	PathPromotionUpdateFollowPrize    = APIPath("follow_prize", "update_follow_prize")
	PathPromotionGetFollowPrizeDetail = APIPath("follow_prize", "get_follow_prize_detail")
	PathPromotionGetFollowPrizeList   = APIPath("follow_prize", "get_follow_prize_list")
	PathPromotionGetTopPicksList      = APIPath("top_picks", "get_top_picks_list")
	PathPromotionAddTopPicks          = APIPath("top_picks", "add_top_picks")
	PathPromotionUpdateTopPicks       = APIPath("top_picks", "update_top_picks")
	PathPromotionDeleteTopPicks       = APIPath("top_picks", "delete_top_picks")
	PathPromotionAddShopCategory      = APIPath("shop_category", "add_shop_category")
	PathPromotionGetShopCategoryList  = APIPath("shop_category", "get_shop_category_list")
	PathPromotionDeleteShopCategory   = APIPath("shop_category", "delete_shop_category")
	PathPromotionUpdateShopCategory   = APIPath("shop_category", "update_shop_category")
	PathPromotionAddShopCategoryItemList   = APIPath("shop_category", "add_item_list")
	PathPromotionGetShopCategoryItemList   = APIPath("product", "get_item_list")
	PathPromotionDeleteShopCategoryItemList = APIPath("shop_category", "delete_item_list")
)

// ---------------------------------------------------------------------------
// Returns
// ---------------------------------------------------------------------------

var (
	PathReturnGetList            = APIPath("returns", "get_return_list")
	PathReturnGetDetail          = APIPath("returns", "get_return_detail")
	PathReturnGetAvailableSolutions = APIPath("returns", "get_available_solutions")
	PathReturnQueryProof         = APIPath("returns", "query_proof")
	PathReturnGetDisputeReason   = APIPath("returns", "get_return_dispute_reason")
	PathReturnCancelDispute      = APIPath("returns", "cancel_dispute")
	PathReturnGetShippingCarrier = APIPath("returns", "get_shipping_carrier")
	PathReturnGetReverseTracking = APIPath("returns", "get_reverse_tracking_info")
)

// ---------------------------------------------------------------------------
// Ads (CPC / Product Ads / GMS)
// ---------------------------------------------------------------------------

var (
	PathAdsGetRecommendedKeywords    = APIPath("ads", "get_recommended_keyword_list")
	PathAdsGetRecommendedItems       = APIPath("ads", "get_recommended_item_list")
	PathAdsGetAllCPCHourlyPerf       = APIPath("ads", "get_all_cpc_ads_hourly_performance")
	PathAdsGetAllCPCDailyPerf        = APIPath("ads", "get_all_cpc_ads_daily_performance")
	PathAdsGetProductCampaignDailyPerf  = APIPath("ads", "get_product_campaign_daily_performance")
	PathAdsGetProductCampaignHourlyPerf = APIPath("ads", "get_product_campaign_hourly_performance")
	PathAdsGetProductLevelCampaignIDs   = APIPath("ads", "get_product_level_campaign_id_list")
	PathAdsGetProductLevelCampaignSetting = APIPath("ads", "get_product_level_campaign_setting_info")
	PathAdsCreateManualProductAds       = APIPath("ads", "create_manual_product_ads")
	PathAdsEditManualProductAdKeywords  = APIPath("ads", "edit_manual_product_ad_keywords")
	PathAdsEditManualProductAds         = APIPath("ads", "edit_manual_product_ads")
	PathAdsGetCreateAdBudgetSuggestion  = APIPath("ads", "get_create_product_ad_budget_suggestion")
	PathAdsGetRecommendedROITarget      = APIPath("ads", "get_product_recommended_roi_target")
	PathAdsGetFacilShopRate             = APIPath("ads", "get_ads_facil_shop_rate")
	PathAdsCheckCreateGMSCampaignElig   = APIPath("ads", "check_create_gms_product_campaign_eligibility")
	PathAdsCreateGMSCampaign            = APIPath("ads", "create_gms_product_campaign")
	PathAdsEditGMSCampaign              = APIPath("ads", "edit_gms_product_campaign")
	PathAdsEditGMSItemCampaign          = APIPath("ads", "edit_gms_item_product_campaign")
	PathAdsGetGMSCampaignPerf           = APIPath("ads", "get_gms_campaign_performance")
	PathAdsGetGMSItemPerf               = APIPath("ads", "get_gms_item_performance")
)

// ---------------------------------------------------------------------------
// Partner / Auth
// ---------------------------------------------------------------------------

var (
	PathPartnerGetShopsByPartner      = APIPath("public", "get_shops_by_partner")
	PathPartnerGetMerchantsByPartner  = APIPath("public", "get_merchants_by_partner")
	PathPartnerGetAccessToken         = APIPath("auth", "token/get")
	PathPartnerRefreshAccessToken     = APIPath("auth", "access_token/get")
	PathPartnerGetTokenByResendCode   = APIPath("public", "get_token_by_resend_code")
	PathPartnerGetShopeeIPRanges      = APIPath("public", "get_shopee_ip_ranges")
	PathPartnerSetAppPushConfig       = APIPath("push", "set_app_push_config")
	PathPartnerGetAppPushConfig       = APIPath("push", "get_app_push_config")
	PathPartnerGetLostPushMessage     = APIPath("push", "get_lost_push_message")
	PathPartnerGetBoundWHSInfo        = APIPath("sbs", "get_bound_whs_info")
)

// ---------------------------------------------------------------------------
// Brazil (BR) Specific
// ---------------------------------------------------------------------------

var (
	PathBRQueryShopEnrollmentStatus = APIPath("fbs", "query_br_shop_enrollment_status")
	PathBRQueryShopInvoiceError     = APIPath("fbs", "query_br_shop_invoice_error")
	PathBRQueryShopBlockStatus      = APIPath("fbs", "query_br_shop_block_status")
	PathBRQuerySKUBlockStatus       = APIPath("fbs", "query_br_sku_block_status")
)

// ---------------------------------------------------------------------------
// WHS (Warehouse)
// ---------------------------------------------------------------------------

var (
	PathWHSGetCurrentInventory = APIPath("sbs", "get_current_inventory")
	PathWHSGetExpiryReport     = APIPath("sbs", "get_expiry_report")
	PathWHSGetStockAging       = APIPath("sbs", "get_stock_aging")
	PathWHSGetStockMovement    = APIPath("sbs", "get_stock_movement")
)

// ---------------------------------------------------------------------------
// Live Streaming / Session
// ---------------------------------------------------------------------------

var (
	PathLiveCreateSession         = APIPath("livestream", "create_session")
	PathLiveUpdateSession         = APIPath("livestream", "update_session")
	PathLiveGetSessionDetail      = APIPath("livestream", "get_session_detail")
	PathLiveAddItemList           = APIPath("shop_category", "add_item_list")
	PathLiveDeleteItemList        = APIPath("shop_category", "delete_item_list")
	PathLiveUpdateItemList        = APIPath("livestream", "update_item_list")
	PathLiveGetItemCount          = APIPath("livestream", "get_item_count")
	PathLiveGetItemList           = APIPath("product", "get_item_list")
	PathLiveUpdateShowItem        = APIPath("livestream", "update_show_item")
	PathLiveDeleteShowItem        = APIPath("livestream", "delete_show_item")
	PathLiveGetShowItem           = APIPath("livestream", "get_show_item")
	PathLiveGetLikeItemList       = APIPath("livestream", "get_like_item_list")
	PathLiveGetRecentItemList     = APIPath("livestream", "get_recent_item_list")
	PathLiveGetItemSetList        = APIPath("livestream", "get_item_set_list")
	PathLiveGetItemSetItemList    = APIPath("livestream", "get_item_set_item_list")
	PathLiveGetSessionMetric      = APIPath("livestream", "get_session_metric")
	PathLiveGetSessionItemMetric  = APIPath("livestream", "get_session_item_metric")
	PathLiveGetLatestCommentList  = APIPath("livestream", "get_latest_comment_list")
)

// Service prefixes for path construction.
var (
	ServiceProduct       = "product"
	ServiceGlobalProduct = "global_product"
	ServiceOrder         = "order"
	ServiceLogistics     = "logistics"
	ServiceShop          = "shop"
	ServiceAMS           = "ams"
	ServicePayment       = "payment"
	ServicePromotion     = "promotion"
	ServiceReturns       = "returns"
	ServiceAds           = "ads"
	ServicePartner       = "partner"
	ServiceAuth          = "auth"
	ServiceBR            = "br"
	ServiceWHS           = "warehouse"
	ServiceLive          = "media"
)

func init() {
	// Validate that APIPath produces consistent results.
	if !strings.HasPrefix(PathProductGetCategory, apiPrefix) {
		panic("programming error: PathProductGetCategory must start with " + apiPrefix)
	}
}
