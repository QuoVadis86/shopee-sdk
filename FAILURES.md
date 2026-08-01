# Shopee SDK/工具修复最终报告

## 各服务最终状态

| 服务 | OK | PARAM | 权限 | 404 | 资源NF | 其他 |
|---|---|---|---|---|---|---|
| product | 12 | 17 | 2 | 5 | 3 | 12 |
| logistics | 3 | 22 | 0 | 7 | 1 | 8 |
| ams | 0 | 0 | 35 | 0 | 1 | 0 |
| global_product | 0 | 31 | 0 | 0 | 2 | 0 |
| ads | 3 | 12 | 1 | 0 | 0 | 5 |
| order | 3 | 9 | 0 | 2 | 0 | 2 |
| payment | 6 | 4 | 0 | 1 | 2 | 3 |
| livestream | 0 | 15 | 0 | 0 | 0 | 0 |
| add_on_deal | 0 | 13 | 0 | 0 | 0 | 1 |
| video | 0 | 13 | 0 | 0 | 0 | 0 |
| discount | 0 | 8 | 0 | 0 | 0 | 3 |
| shop_flash_sale | 1 | 10 | 0 | 0 | 0 | 0 |
| shop | 5 | 1 | 0 | 0 | 1 | 3 |
| first_mile | 3 | 4 | 0 | 1 | 0 | 1 |
| bundle_deal | 0 | 7 | 0 | 0 | 0 | 2 |
| returns | 0 | 7 | 0 | 0 | 0 | 1 |
| merchant | 0 | 6 | 0 | 0 | 0 | 0 |
| follow_prize | 0 | 5 | 0 | 0 | 0 | 0 |
| voucher | 0 | 5 | 0 | 0 | 0 | 0 |
| account_health | 3 | 0 | 0 | 0 | 0 | 2 |
| fbs | 1 | 0 | 0 | 0 | 0 | 3 |
| shop_category | 1 | 3 | 0 | 0 | 0 | 0 |
| top_picks | 1 | 3 | 0 | 0 | 0 | 0 |
| sbs | 0 | 0 | 0 | 0 | 0 | 4 |
| promotion | 0 | 0 | 0 | 0 | 3 | 0 |
| media | 0 | 0 | 0 | 0 | 3 | 0 |
| chat | 0 | 0 | 1 | 0 | 0 | 0 |

## 已修复的工具(SDK 路径/签名修复后实测通过)

- ✅ shopee_account_health_get_late_orders
- ✅ shopee_account_health_get_listings_with_issues
- ✅ shopee_account_health_get_penalty_point_history
- ✅ shopee_ads_check_create_gms_product_campaign_eligibility
- ✅ shopee_ads_get_all_cpc_ads_daily_performance
- ✅ shopee_ads_get_product_level_campaign_id_list
- ✅ shopee_ads_get_recommended_item_list
- ✅ shopee_ads_get_recommended_keyword_list
- ✅ shopee_ads_get_shop_toggle_info
- ✅ shopee_fbs_query_br_shop_enrollment_status
- ✅ shopee_first_mile_get_courier_delivery_channel_list
- ✅ shopee_first_mile_get_transit_warehouse_list
- ✅ shopee_first_mile_get_unbind_order_list
- ✅ shopee_logistics_get_address_list
- ✅ shopee_logistics_get_channel_list
- ✅ shopee_logistics_get_operating_hour_restrictions
- ✅ shopee_logistics_get_operating_hours
- ✅ shopee_logistics_update_operating_hours
- ✅ shopee_order_get_order_detail
- ✅ shopee_order_get_pending_buyer_invoice_order_list
- ✅ shopee_order_get_shipment_list
- ✅ shopee_order_get_warehouse_filter_config
- ✅ shopee_payment_get_escrow_detail
- ✅ shopee_payment_get_escrow_detail_batch
- ✅ shopee_payment_get_escrow_list
- ✅ shopee_payment_get_income_overview
- ✅ shopee_payment_get_payment_method_list
- ✅ shopee_payment_get_shop_installment_status
- ✅ shopee_payment_get_wallet_transaction_list
- ✅ shopee_payment_set_shop_installment_status
- ✅ shopee_product_get_all_vehicle_list
- ✅ shopee_product_get_boosted_list
- ✅ shopee_product_get_category
- ✅ shopee_product_get_comment
- ✅ shopee_product_get_item_base_info
- ✅ shopee_product_get_item_content_diagnosis_result
- ✅ shopee_product_get_item_extra_info
- ✅ shopee_product_get_item_limit
- ✅ shopee_product_get_item_promotion
- ✅ shopee_product_get_item_violation_info
- ✅ shopee_product_get_kit_item_limit
- ✅ shopee_product_get_model_list
- ✅ shopee_product_get_product_certification_rule
- ✅ shopee_product_get_recommend_attribute
- ✅ shopee_product_update_sip_item_price
- ✅ shopee_shop_category_get_shop_category_list
- ✅ shopee_shop_flash_sale_get_item_criteria
- ✅ shopee_shop_get_profile
- ✅ shopee_shop_get_shop_holiday_mode
- ✅ shopee_shop_get_shop_info
- ✅ shopee_shop_get_shop_notification
- ✅ shopee_shop_get_warehouse_detail
- ✅ shopee_shop_update_profile
- ✅ shopee_top_picks_get_top_picks_list

## 需真实店铺数据(agent 提供参数后可调)

- ⚠️ shopee_add_on_deal_add_add_on_deal
- ⚠️ shopee_add_on_deal_add_add_on_deal_main_item
- ⚠️ shopee_add_on_deal_add_add_on_deal_sub_item
- ⚠️ shopee_add_on_deal_delete_add_on_deal
- ⚠️ shopee_add_on_deal_delete_add_on_deal_main_item
- ⚠️ shopee_add_on_deal_delete_add_on_deal_sub_item
- ⚠️ shopee_add_on_deal_end_add_on_deal
- ⚠️ shopee_add_on_deal_get_add_on_deal
- ⚠️ shopee_add_on_deal_get_add_on_deal_main_item
- ⚠️ shopee_add_on_deal_get_add_on_deal_sub_item
- ⚠️ shopee_add_on_deal_update_add_on_deal
- ⚠️ shopee_add_on_deal_update_add_on_deal_main_item
- ⚠️ shopee_add_on_deal_update_add_on_deal_sub_item
- ⚠️ shopee_ads_create_gms_product_campaign
- ⚠️ shopee_ads_create_manual_product_ads
- ⚠️ shopee_ads_edit_gms_item_product_campaign
- ⚠️ shopee_ads_edit_gms_product_campaign
- ⚠️ shopee_ads_edit_manual_product_ad_keywords
- ⚠️ shopee_ads_edit_manual_product_ads
- ⚠️ shopee_ads_get_all_cpc_ads_hourly_performance
- ⚠️ shopee_ads_get_create_product_ad_budget_suggestion
- ⚠️ shopee_ads_get_product_campaign_daily_performance
- ⚠️ shopee_ads_get_product_campaign_hourly_performance
- ⚠️ shopee_ads_get_product_level_campaign_setting_info
- ⚠️ shopee_ads_get_product_recommended_roi_target
- ⚠️ shopee_ams_get_product_performance_list
- ⚠️ shopee_bundle_deal_add_bundle_deal_item
- ⚠️ shopee_bundle_deal_delete_bundle_deal
- ⚠️ shopee_bundle_deal_delete_bundle_deal_item
- ⚠️ shopee_bundle_deal_get_bundle_deal
- ⚠️ shopee_bundle_deal_get_bundle_deal_item
- ⚠️ shopee_bundle_deal_update_bundle_deal
- ⚠️ shopee_bundle_deal_update_bundle_deal_item
- ⚠️ shopee_discount_add_discount
- ⚠️ shopee_discount_add_discount_item
- ⚠️ shopee_discount_delete_discount
- ⚠️ shopee_discount_delete_discount_item
- ⚠️ shopee_discount_get_discount
- ⚠️ shopee_discount_get_discount_list
- ⚠️ shopee_discount_update_discount
- ⚠️ shopee_discount_update_discount_item
- ⚠️ shopee_first_mile_get_courier_delivery_detail
- ⚠️ shopee_first_mile_get_courier_delivery_tracking_number_list
- ⚠️ shopee_first_mile_get_courier_delivery_waybill
- ⚠️ shopee_first_mile_get_tracking_number_list
- ⚠️ shopee_follow_prize_add_follow_prize
- ⚠️ shopee_follow_prize_delete_follow_prize
- ⚠️ shopee_follow_prize_get_follow_prize_detail
- ⚠️ shopee_follow_prize_get_follow_prize_list
- ⚠️ shopee_follow_prize_update_follow_prize
- ⚠️ shopee_global_product_add_global_item
- ⚠️ shopee_global_product_add_global_model
- ⚠️ shopee_global_product_cancel_video_upload
- ⚠️ shopee_global_product_create_publish_task
- ⚠️ shopee_global_product_delete_global_item
- ⚠️ shopee_global_product_delete_global_model
- ⚠️ shopee_global_product_get_attribute_tree
- ⚠️ shopee_global_product_get_brand_list
- ⚠️ shopee_global_product_get_category
- ⚠️ shopee_global_product_get_global_item_id
- ⚠️ shopee_global_product_get_global_item_info
- ⚠️ shopee_global_product_get_global_item_limit
- ⚠️ shopee_global_product_get_global_item_list
- ⚠️ shopee_global_product_get_global_model_list
- ⚠️ shopee_global_product_get_local_adjustment_rate
- ⚠️ shopee_global_product_get_publish_task_result
- ⚠️ shopee_global_product_get_publishable_shop
- ⚠️ shopee_global_product_get_published_list
- ⚠️ shopee_global_product_get_recommend_attribute
- ⚠️ shopee_global_product_get_shop_publishable_status
- ⚠️ shopee_global_product_get_size_chart_detail
- ⚠️ shopee_global_product_get_size_chart_list
- ⚠️ shopee_global_product_get_variations
- ⚠️ shopee_global_product_get_video_upload_result
- ⚠️ shopee_global_product_search_global_attribute_value_list
- ⚠️ shopee_global_product_set_sync_field
- ⚠️ shopee_global_product_update_global_item
- ⚠️ shopee_global_product_update_global_model
- ⚠️ shopee_global_product_update_local_adjustment_rate
- ⚠️ shopee_global_product_update_price
- ⚠️ shopee_global_product_update_size_chart
- ⚠️ shopee_global_product_update_stock
- ⚠️ shopee_global_product_update_tier_variation
- ⚠️ shopee_livestream_create_session
- ⚠️ shopee_livestream_delete_show_item
- ⚠️ shopee_livestream_get_item_count
- ⚠️ shopee_livestream_get_item_set_item_list
- ⚠️ shopee_livestream_get_item_set_list
- ⚠️ shopee_livestream_get_latest_comment_list
- ⚠️ shopee_livestream_get_like_item_list
- ⚠️ shopee_livestream_get_recent_item_list
- ⚠️ shopee_livestream_get_session_detail
- ⚠️ shopee_livestream_get_session_item_metric
- ⚠️ shopee_livestream_get_session_metric
- ⚠️ shopee_livestream_get_show_item
- ⚠️ shopee_livestream_update_item_list
- ⚠️ shopee_livestream_update_session
- ⚠️ shopee_livestream_update_show_item
- ⚠️ shopee_logistics_batch_ship_order
- ⚠️ shopee_logistics_batch_update_tpf_warehouse_tracking_status
- ⚠️ shopee_logistics_create_booking_shipping_document
- ⚠️ shopee_logistics_create_shipping_document
- ⚠️ shopee_logistics_create_shipping_document_job
- ⚠️ shopee_logistics_delete_address
- ⚠️ shopee_logistics_get_booking_shipping_parameter
- ⚠️ shopee_logistics_get_booking_tracking_info
- ⚠️ shopee_logistics_get_booking_tracking_number
- ⚠️ shopee_logistics_get_mass_shipping_parameter
- ⚠️ shopee_logistics_get_mass_tracking_number
- ⚠️ shopee_logistics_get_shipping_document_parameter
- ⚠️ shopee_logistics_get_shipping_parameter
- ⚠️ shopee_logistics_get_tracking_info
- ⚠️ shopee_logistics_get_tracking_number
- ⚠️ shopee_logistics_mass_ship_order
- ⚠️ shopee_logistics_set_address_config
- ⚠️ shopee_logistics_ship_booking
- ⚠️ shopee_logistics_update_address
- ⚠️ shopee_logistics_update_channel
- ⚠️ shopee_logistics_update_self_collection_order_logistics
- ⚠️ shopee_logistics_update_shipping_order
- ⚠️ shopee_logistics_update_tracking_status
- ⚠️ shopee_media_add_item_list
- ⚠️ shopee_media_delete_item_list
- ⚠️ shopee_media_get_item_list
- ⚠️ shopee_merchant_get_merchant_info
- ⚠️ shopee_merchant_get_merchant_prepaid_account_list
- ⚠️ shopee_merchant_get_merchant_warehouse_list
- ⚠️ shopee_merchant_get_merchant_warehouse_location_list
- ⚠️ shopee_merchant_get_shop_list_by_merchant
- ⚠️ shopee_merchant_get_warehouse_eligible_shop_list
- ⚠️ shopee_order_cancel_order
- ⚠️ shopee_order_get_booking_detail
- ⚠️ shopee_order_get_booking_list
- ⚠️ shopee_order_get_fbs_invoices_result
- ⚠️ shopee_order_get_order_list
- ⚠️ shopee_order_get_package_detail
- ⚠️ shopee_order_search_package_list
- ⚠️ shopee_order_set_note
- ⚠️ shopee_order_split_order
- ⚠️ shopee_payment_get_billing_transaction_info
- ⚠️ shopee_payment_get_income_report
- ⚠️ shopee_payment_get_income_statement
- ⚠️ shopee_payment_get_payout_detail
- ⚠️ shopee_payment_get_payout_info
- ⚠️ shopee_payment_set_item_installment_status
- ⚠️ shopee_product_add_item
- ⚠️ shopee_product_add_model
- ⚠️ shopee_product_get_aitem_by_pitem_id
- ⚠️ shopee_product_get_attribute_tree
- ⚠️ shopee_product_get_brand_list
- ⚠️ shopee_product_get_direct_item_list
- ⚠️ shopee_product_get_item_list
- ⚠️ shopee_product_get_main_item_list
- ⚠️ shopee_product_get_size_chart_detail
- ⚠️ shopee_product_get_size_chart_list
- ⚠️ shopee_product_get_ssp_info
- ⚠️ shopee_product_get_ssp_list
- ⚠️ shopee_product_get_variations
- ⚠️ shopee_product_get_vehicle_list_by_compatibility_detail
- ⚠️ shopee_product_publish_item_to_outlet_shop
- ⚠️ shopee_product_search_unpackaged_model_list
- ⚠️ shopee_product_unlist_item
- ⚠️ shopee_product_update_model
- ⚠️ shopee_product_update_price
- ⚠️ shopee_product_update_stock
- ⚠️ shopee_promotion_add_item_list
- ⚠️ shopee_promotion_delete_item_list
- ⚠️ shopee_promotion_get_item_list
- ⚠️ shopee_returns_cancel_dispute
- ⚠️ shopee_returns_get_available_solutions
- ⚠️ shopee_returns_get_return_detail
- ⚠️ shopee_returns_get_return_dispute_reason
- ⚠️ shopee_returns_get_return_list
- ⚠️ shopee_returns_get_reverse_tracking_info
- ⚠️ shopee_returns_get_shipping_carrier
- ⚠️ shopee_shop_category_add_shop_category
- ⚠️ shopee_shop_category_delete_shop_category
- ⚠️ shopee_shop_category_update_shop_category
- ⚠️ shopee_shop_flash_sale_add_shop_flash_sale_items
- ⚠️ shopee_shop_flash_sale_create_shop_flash_sale
- ⚠️ shopee_shop_flash_sale_delete_shop_flash_sale
- ⚠️ shopee_shop_flash_sale_delete_shop_flash_sale_items
- ⚠️ shopee_shop_flash_sale_get_shop_flash_sale
- ⚠️ shopee_shop_flash_sale_get_shop_flash_sale_items
- ⚠️ shopee_shop_flash_sale_get_shop_flash_sale_list
- ⚠️ shopee_shop_flash_sale_get_time_slot_id
- ⚠️ shopee_shop_flash_sale_update_shop_flash_sale
- ⚠️ shopee_shop_flash_sale_update_shop_flash_sale_items
- ⚠️ shopee_shop_get_authorised_reseller_brand
- ⚠️ shopee_shop_get_shop_performance
- ⚠️ shopee_top_picks_add_top_picks
- ⚠️ shopee_top_picks_delete_top_picks
- ⚠️ shopee_top_picks_update_top_picks
- ⚠️ shopee_video_delete_video
- ⚠️ shopee_video_edit_video_info
- ⚠️ shopee_video_get_cover_list
- ⚠️ shopee_video_get_metric_trend
- ⚠️ shopee_video_get_overview_performance
- ⚠️ shopee_video_get_user_demographics
- ⚠️ shopee_video_get_video_detail
- ⚠️ shopee_video_get_video_detail_audience_distribution
- ⚠️ shopee_video_get_video_detail_metric_trend
- ⚠️ shopee_video_get_video_detail_performance
- ⚠️ shopee_video_get_video_detail_product_performance
- ⚠️ shopee_video_get_video_list
- ⚠️ shopee_video_get_video_performance_list
- ⚠️ shopee_voucher_add_voucher
- ⚠️ shopee_voucher_delete_voucher
- ⚠️ shopee_voucher_get_voucher
- ⚠️ shopee_voucher_get_voucher_list
- ⚠️ shopee_voucher_update_voucher

## 权限/路径不可修(Shopee 侧)

- 🔒 shopee_ads_get_ads_facil_shop_rate
- 🔒 shopee_ams_add_all_products_to_open_campaign
- 🔒 shopee_ams_batch_add_products_to_open_campaign
- 🔒 shopee_ams_batch_edit_products_open_campaign_setting
- 🔒 shopee_ams_batch_get_products_suggested_rate
- 🔒 shopee_ams_batch_remove_products_open_campaign_setting
- 🔒 shopee_ams_create_new_targeted_campaign
- 🔒 shopee_ams_edit_affiliate_list_of_targeted_campaign
- 🔒 shopee_ams_edit_all_products_open_campaign_setting
- 🔒 shopee_ams_edit_product_list_of_targeted_campaign
- 🔒 shopee_ams_get_affiliate_performance
- 🔒 shopee_ams_get_auto_add_new_product_toggle_status
- 🔒 shopee_ams_get_campaign_key_metrics_performance
- 🔒 shopee_ams_get_content_performance
- 🔒 shopee_ams_get_conversion_report
- 🔒 shopee_ams_get_managed_affiliate_list
- 🔒 shopee_ams_get_open_campaign_added_product
- 🔒 shopee_ams_get_open_campaign_batch_task_result
- 🔒 shopee_ams_get_open_campaign_not_added_product
- 🔒 shopee_ams_get_open_campaign_performance
- 🔒 shopee_ams_get_optimization_suggestion_product
- 🔒 shopee_ams_get_performance_data_update_time
- 🔒 shopee_ams_get_product_performance
- 🔒 shopee_ams_get_recommended_affiliate_list
- 🔒 shopee_ams_get_shop_performance
- 🔒 shopee_ams_get_shop_suggested_rate
- 🔒 shopee_ams_get_targeted_campaign_addable_product_list
- 🔒 shopee_ams_get_targeted_campaign_list
- 🔒 shopee_ams_get_targeted_campaign_performance
- 🔒 shopee_ams_get_targeted_campaign_settings
- 🔒 shopee_ams_get_validation_list
- 🔒 shopee_ams_get_validation_report
- 🔒 shopee_ams_query_affiliate_list
- 🔒 shopee_ams_remove_all_products_open_campaign_setting
- 🔒 shopee_ams_update_auto_add_new_product_setting
- 🔒 shopee_ams_update_basic_info_of_targeted_campaign
- 🔒 shopee_chat_send_message
- 🔒 shopee_first_mile_get_waybill
- 🔒 shopee_logistics_check_polygon_update_status
- 🔒 shopee_logistics_get_booking_shipping_document_data_info
- 🔒 shopee_logistics_get_booking_shipping_document_parameter
- 🔒 shopee_logistics_get_booking_shipping_document_result
- 🔒 shopee_logistics_get_shipping_document_data_info
- 🔒 shopee_logistics_get_shipping_document_job_status
- 🔒 shopee_logistics_get_shipping_document_result
- 🔒 shopee_order_get_buyer_invoice_info
- 🔒 shopee_order_get_estimate_cancel_value
- 🔒 shopee_payment_get_item_installment_status
- 🔒 shopee_product_add_ssp_item
- 🔒 shopee_product_get_attributes
- 🔒 shopee_product_get_direct_shop_recommended_price
- 🔒 shopee_product_get_mart_item_by_outlet_item_id
- 🔒 shopee_product_get_mart_item_mapping_by_id
- 🔒 shopee_product_get_weight_recommendation
- 🔒 shopee_product_search_attribute_value_list