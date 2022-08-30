package kreditplus

type KreditplusService interface {
	GetCheckoutURL(param CheckoutRequest) (resp Response, err error)
	UpdateTrackingStatus(param UpdateTrackingRequest) (resp Response, err error)
	CancelOrder(param CancelRequest) (resp Response, err error)
}
