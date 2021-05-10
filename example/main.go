package main

import (
    "encoding/json"
    "fmt"
    "github.com/Bhinneka/kreditplus-go"
    "time"
)

const (
    BaseURL        = "https://sandbox.kreditplus.com/hub-api"
    Authorization  = "SEdtxXLQSDRCeTaIHXWISotSUmdFxzUE"
    DefaultTimeout = 10 * time.Second
)

func main() {
    newKP := kreditplus.New(BaseURL, Authorization, DefaultTimeout)
    GetCheckoutURL(newKP)
    //UpdateOrderTracking(newKP)
}

func GetCheckoutURL(kp kreditplus.KreditplusService) {
    var checkoutReq kreditplus.CheckoutRequest


    payload := `{"prospect_id": "ORDER_TEST_45679", "shipping_cost": 7000, "merchant_url": "https://bhinneka.com", "mobile_phone": "081212341234", "assets": [{"product_id": "NA", "asset_type": "WG", "asset_code": "PHABLET", "category_code": "TEST", "category_name": "PHONE TABLET", "discount_amount": 20000, "dp_amount": 0, "insurance_amount": 0, "item_description": "test", "otr": 1500000, "quantity": 1}],"shipping_address": { "receiver_name":"test", "address":"test ini khusus alamat buat test aja", "mobile_phone":"081212345678"}}`
    _ = json.Unmarshal([]byte(payload), &checkoutReq)


    checkoutResponse, errGetCheckoutURL := kp.GetCheckoutURL(checkoutReq)
    if errGetCheckoutURL != nil {
        fmt.Println(errGetCheckoutURL.Error())
        return
    }

    fmt.Println(checkoutResponse)
}
func UpdateOrderTracking(kp kreditplus.KreditplusService) {
    var UpdateOrderTracking kreditplus.UpdateTrackingRequest
    payload := `{"prospect_id": "ORDER_TEST_45678","awb_number":"test1234551244", "logistic_name":"JNE", "tracking_status":"Barang Dalam Pengiriman"}`
    _ = json.Unmarshal([]byte(payload), &UpdateOrderTracking)


    updateTrackingResp, errUpdateTracking := kp.UpdateTrackingStatus(UpdateOrderTracking)
    if errUpdateTracking != nil {
        fmt.Println(errUpdateTracking.Error())
        return
    }

    fmt.Println(updateTrackingResp)
}


