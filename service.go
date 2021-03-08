package kreditplus

type KreditplusService interface {
    GetCheckoutURL(param CheckoutRequest) (resp Response, err error)
}
