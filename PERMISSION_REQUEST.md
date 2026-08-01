# Shopee App 权限申请清单

以下工具已正确实现（路径/参数与官方文档一致），但 Shopee 返回 `error_api_permission`。
需要在 Shopee 开放平台（open.shopee.com）为 app 开通对应 API 权限后即可使用。

| 服务 | API 数量 | 需要开通的权限 |
|---|---|---|
| ams | 35 | AMS (Shopee 营销/广告服务) 权限 |
| product | 2 | Kit/SSP 等特殊商品权限 |
| ads | 1 | Shopee Ads 广告服务权限 |
| chat | 1 | Chat API 权限 |

## 明细

- `shopee_ads_get_ads_facil_shop_rate`
- `shopee_ams_add_all_products_to_open_campaign`
- `shopee_ams_batch_add_products_to_open_campaign`
- `shopee_ams_batch_edit_products_open_campaign_setting`
- `shopee_ams_batch_get_products_suggested_rate`
- `shopee_ams_batch_remove_products_open_campaign_setting`
- `shopee_ams_create_new_targeted_campaign`
- `shopee_ams_edit_affiliate_list_of_targeted_campaign`
- `shopee_ams_edit_all_products_open_campaign_setting`
- `shopee_ams_edit_product_list_of_targeted_campaign`
- `shopee_ams_get_affiliate_performance`
- `shopee_ams_get_auto_add_new_product_toggle_status`
- `shopee_ams_get_campaign_key_metrics_performance`
- `shopee_ams_get_content_performance`
- `shopee_ams_get_conversion_report`
- `shopee_ams_get_managed_affiliate_list`
- `shopee_ams_get_open_campaign_added_product`
- `shopee_ams_get_open_campaign_batch_task_result`
- `shopee_ams_get_open_campaign_not_added_product`
- `shopee_ams_get_open_campaign_performance`
- `shopee_ams_get_optimization_suggestion_product`
- `shopee_ams_get_performance_data_update_time`
- `shopee_ams_get_product_performance`
- `shopee_ams_get_recommended_affiliate_list`
- `shopee_ams_get_shop_performance`
- `shopee_ams_get_shop_suggested_rate`
- `shopee_ams_get_targeted_campaign_addable_product_list`
- `shopee_ams_get_targeted_campaign_list`
- `shopee_ams_get_targeted_campaign_performance`
- `shopee_ams_get_targeted_campaign_settings`
- `shopee_ams_get_validation_list`
- `shopee_ams_get_validation_report`
- `shopee_ams_query_affiliate_list`
- `shopee_ams_remove_all_products_open_campaign_setting`
- `shopee_ams_update_auto_add_new_product_setting`
- `shopee_ams_update_basic_info_of_targeted_campaign`
- `shopee_chat_send_message`
- `shopee_product_add_ssp_item`
- `shopee_product_get_attributes`