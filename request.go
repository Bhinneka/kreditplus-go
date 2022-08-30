package kreditplus

type CheckoutRequest struct {
	ProspectID      string          `json:"prospect_id"`
	ShippingCost    int             `json:"shipping_cost"`
	MerchantURL     string          `json:"merchant_url"`
	MobilePhone     string          `json:"mobile_phone"`
	Assets          []Asset         `json:"assets"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type Asset struct {
	AssetCode       string `json:"asset_code"`
	AssetType       string `json:"asset_type"`
	CategoryCode    string `json:"category_code"`
	CategoryName    string `json:"category_name"`
	DiscountAmount  int    `json:"discount_amount"`
	DPAmount        int    `json:"dp_amount"`
	InsuranceAmount int    `json:"insurance_amount"`
	ItemDescription string `json:"item_description"`
	OTR             int    `json:"otr"`
	ProductID       string `json:"product_id"`
	Quantity        int    `json:"quantity"`
}

type ShippingAddress struct {
	ReceiverName string `json:"receiver_name"`
	Address      string `json:"address"`
	MobilePhone  string `json:"mobile_phone"`
}

type UpdateTrackingRequest struct {
	AWBNumber     string `json:"awb_number"`
	LogisticName  string `json:"logistic_name"`
	ProspectID    string `json:"prospect_id"`
	TracingStatus string `json:"tracking_status"`
}

type CancelRequest struct {
	CancelReason string `json:"cancel_reason"`
	ProspectID   string `json:"prospect_id"`
}
