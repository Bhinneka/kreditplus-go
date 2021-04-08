package kreditplus

//ServiceResult
type ServiceResult struct {
    Result interface{}
    Error  error
}

type Response struct {
    Code     int      `json:"code"`
    Errors   *[]Error `json:"errors"`
    Messages string   `json:"messages"`
    Data     *Data    `json:"data"`
}

type Data struct {
    CheckoutURL string `json:"checkout_url"`
    ProspectID  string `json:"prospect_id"`
    ExpiredAt   string `json:"expired_at"`
}

type Error struct {
    Parameter string `json:"parameter"`
    Message   string `json:"message"`
}
