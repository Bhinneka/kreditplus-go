package main

import (
    "encoding/json"
    "fmt"
    "github.com/Bhinneka/kreditplus-go"
    "time"
)

const (
    BaseURL        = "YOUR_BASE_URL"
    Authorization  = "YOUR_API_KEY"
    DefaultTimeout = 10 * time.Second
)

func main() {
    newKP := kreditplus.New(BaseURL, Authorization, DefaultTimeout)
    getCheckoutURL(newKP)
}

func getCheckoutURL(kp kreditplus.KreditplusService) {
    var checkoutReq kreditplus.CheckoutRequest
    payload := `{"prospect_id": "ORDER_TEST_456", "shipping_cost": 7000, "merchant_url": "https://bhinneka.com", "mobile_phone": "081212341234", "assets": [{"product_id": "NA", "asset_type": "WG", "asset_code": "PHABLET", "category_code": "TEST", "category_name": "PHONE TABLET", "discount_amount": 20000, "dp_amount": 0, "insurance_amount": 0, "item_description": "test", "otr": 1500000, "quantity": 1}]}`
    _ = json.Unmarshal([]byte(payload), &checkoutReq)


    checkoutResponse, errGetCheckoutURL := kp.GetCheckoutURL(checkoutReq)
    if errGetCheckoutURL != nil {
        fmt.Println(errGetCheckoutURL.Error())
        return
    }

    fmt.Println(checkoutResponse)
}



