# Shopee Open Platform Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/meta001/shopee-sdk.svg)](https://pkg.go.dev/github.com/meta001/shopee-sdk)
[![Go Version](https://img.shields.io/github/go-mod/go-version/meta001/shopee-sdk)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/meta001/shopee-sdk)](LICENSE)

A comprehensive Go SDK for the [Shopee Open Platform v2 API](https://open.shopee.com). Covers **380+ API endpoints** across 15 service domains — product management, orders, logistics, promotions, advertising, finance, and more.

## Features

- **Full API Coverage** — Every v2 API endpoint from the official Shopee documentation
- **Idiomatic Go** — Clean, flat package design with no unnecessary abstractions
- **Automatic Signing** — HMAC-SHA256 signature generation for every request
- **Type Safety** — Strongly typed request/response structures for all endpoints
- **Region Support** — Built-in endpoints for Global, China, Brazil, Sandbox
- **No External Dependencies** — Pure standard library HTTP client
- **Context Ready** — Easy to integrate with any Go application

## Installation

```bash
go get github.com/meta001/shopee-sdk
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/meta001/shopee-sdk"
)

func main() {
    sdk := shopee.NewSDK(
        123456,                     // partner_id
        "your-partner-key",         // partner_key
        "your-access-token",        // access_token
        12345,                      // shop_id
        shopee.WithRegion(shopee.RegionSandbox),
    )

    // Shop info
    shop, err := sdk.Shop.GetShopInfo()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Shop: %s (%d)\n", shop.Response.ShopName, shop.Response.ShopID)

    // Product list
    items, err := sdk.Product.GetItemList(0, 10, 0, 0, "")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Total items: %d\n", items.Response.TotalCount)

    // Orders
    orders, err := sdk.Order.GetOrderList("create_time", 0, 0, 10, "", "")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Orders: %d\n", len(orders.Response.OrderList))
}
```

## Services

| Service | File | Endpoints | Description |
|---------|------|-----------|-------------|
| `sdk.Product` | `product.go` | 51 | Local product CRUD, categories, brands, models, stock, price |
| `sdk.GlobalProduct` | `global_product.go` | 34 | Cross-border product management, publish tasks |
| `sdk.Order` | `order.go` | 17 | Order list/detail, split, cancel, notes, invoices |
| `sdk.Logistics` | `logistics.go` | 51 | Shipping, tracking, addresses, channels, documents |
| `sdk.Shop` | `shop.go` | 24 | Shop info, profile, warehouse, performance, penalties |
| `sdk.AMS` | `ams.go` | 50 | Open/ Targeted campaigns, affiliate, performance reports |
| `sdk.Payment` | `payment.go` | 17 | Escrow, payouts, installment, income statements |
| `sdk.Promotion` | `promotion.go` | 67 | Discounts, bundle deals, add-on deals, vouchers, flash sales |
| `sdk.Returns` | `returns.go` | 9 | Return list/detail, disputes, reverse tracking |
| `sdk.Ads` | `ads.go` | 21 | CPC ads, product campaigns, GMS campaigns |
| `sdk.Partner` | `partner.go` | 10 | Partner shops, access tokens, IP ranges, push config |
| `sdk.BR` | `br.go` | 5 | Brazil-specific enrollment, invoice, block status |
| `sdk.WHS` | `whs.go` | 5 | Warehouse inventory, expiry, stock aging/movement |
| `sdk.Live` | `live.go` | 19 | Live session management, items, metrics, comments |

## Configuration

### Regions

```go
shopee.WithRegion(shopee.RegionGlobal)      // https://partner.shopeemobile.com
shopee.WithRegion(shopee.RegionCN)          // https://openplatform.shopee.cn
shopee.WithRegion(shopee.RegionBrazil)      // https://openplatform.shopee.com.br
shopee.WithRegion(shopee.RegionSandbox)     // test environment
shopee.WithRegion(shopee.RegionSandboxCN)   // China sandbox
```

### Custom HTTP Client

```go
sdk := shopee.NewSDK(pid, key, token, sid,
    shopee.WithHTTPClient(&http.Client{
        Timeout: 60 * time.Second,
    }),
)
```

### Updating Access Token

```go
sdk.SetAccessToken("new-token")
```

## Authentication

1. **Get authorization URL:** Use the pre-defined `AuthURLs` map by region
2. **User authorizes** in browser, Shopee redirects with a `code`
3. **Exchange code for token:** `sdk.Partner.GetAccessToken(&shopee.GetAccessTokenParams{...})`
4. **Use the token:** Pass it to `NewSDK` or update via `SetAccessToken`

## Error Handling

Every API method returns `(*ResponseType, error)`. Two error types:

- **`*shopee.APIError`** — Shopee returned an error response (check `ErrorCode`, `Message`, `RequestID`)
- **Network/parsing errors** — Standard Go errors for connectivity, JSON unmarshal, etc.

```go
result, err := sdk.Product.GetItemList(0, 10, 0, 0, "")
if err != nil {
    var apiErr *shopee.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error [%s]: %s\n", apiErr.ErrorCode, apiErr.Message)
    } else {
        fmt.Printf("Network error: %v\n", err)
    }
    return
}
```

Or check the response directly:

```go
if result.HasError() {
    fmt.Printf("Warning: %s - %s\n", result.Error, result.Message)
}
```

## Signing

The SDK automatically signs every request using HMAC-SHA256. The base string is:

```
partner_id + api_path + timestamp + [access_token] + [shop_id]
```

## Development

```bash
git clone https://github.com/meta001/shopee-sdk.git
cd shopee-sdk
go vet ./...
go build ./...
```

## License

MIT
