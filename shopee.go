package shopee

// SDK is the main entry point for the Shopee Open Platform SDK.
// It aggregates all service clients for different API domains.
type SDK struct {
	Client        *Client
	Product       *ProductService
	GlobalProduct *GlobalProductService
	Order         *OrderService
	Logistics     *LogisticsService
	Shop          *ShopService
	AMS           *AMSService
	Payment       *PaymentService
	Promotion     *PromotionService
	Returns       *ReturnsService
	Ads           *AdsService
	Partner       *PartnerService
	BR            *BRService
	WHS           *WHSService
	Live          *LiveService
}

// NewSDK creates a new Shopee SDK with all service clients.
func NewSDK(partnerID int64, partnerKey, accessToken string, shopID int64, opts ...ClientOption) *SDK {
	c := NewClient(partnerID, partnerKey, accessToken, shopID, opts...)
	return &SDK{
		Client:        c,
		Product:       NewProductService(c),
		GlobalProduct: NewGlobalProductService(c),
		Order:         NewOrderService(c),
		Logistics:     NewLogisticsService(c),
		Shop:          NewShopService(c),
		AMS:           NewAMSService(c),
		Payment:       NewPaymentService(c),
		Promotion:     NewPromotionService(c),
		Returns:       NewReturnsService(c),
		Ads:           NewAdsService(c),
		Partner:       NewPartnerService(c),
		BR:            NewBRService(c),
		WHS:           NewWHSService(c),
		Live:          NewLiveService(c),
	}
}

// SetAccessToken updates the access token used by all services.
func (s *SDK) SetAccessToken(token string) {
	s.Client.AccessToken = token
}
