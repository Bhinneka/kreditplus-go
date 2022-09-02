package kreditplus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Kreditplus struct {
	BaseURL       string
	Authorization string
	client        *kreditplusHttpClient
	*logger
}

func New(baseUrl string, authorization string, timeout time.Duration) *Kreditplus {
	httpRequest := newRequest(timeout)
	return &Kreditplus{
		BaseURL:       baseUrl,
		Authorization: authorization,
		client:        httpRequest,
		logger:        newLogger(),
	}
}

func (kp *Kreditplus) call(method string, path string, body io.Reader, v interface{}, headers map[string]string) error {
	kp.info().Println("Starting http call..")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = fmt.Sprintf("%s%s", kp.BaseURL, path)
	return kp.client.exec(method, path, body, v, headers)
}

func (kp *Kreditplus) GetCheckoutURL(request CheckoutRequest) (resp Response, err error) {
	kp.info().Println("Starting Get Checkout URL KreditPlus")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			kp.error().Println(err.Error())
		}
	}()
	var response Response

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = kp.Authorization

	pathURL := "/api/checkout/order"
	//Marshal Order
	payload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	err = kp.call("POST", pathURL, bytes.NewBuffer(payload), &response, headers)
	if err != nil {
		return response, err
	}

	if response.Code != http.StatusOK {
		err = errors.New(response.Messages)
		return response, err
	}

	return response, nil
}

func (kp *Kreditplus) UpdateTrackingStatus(request UpdateTrackingRequest) (resp Response, err error) {
	kp.info().Println("Starting Update Tracking Status KreditPlus")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			kp.error().Println(err.Error())
		}
	}()
	var response Response

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = kp.Authorization

	pathURL := "/api/order/callback_tracking_online"
	//Marshal Payload
	payload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	err = kp.call("POST", pathURL, bytes.NewBuffer(payload), &response, headers)
	if err != nil {
		return response, err
	}

	if response.Code != http.StatusOK {
		err = errors.New(response.Messages)
		return response, err
	}

	return response, nil
}

func (kp *Kreditplus) CancelOrder(request CancelRequest) (resp Response, err error) {
	kp.info().Println("Starting Cancel Order KreditPlusLimit")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			kp.error().Println(err.Error())
		}
	}()
	var response Response

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	headers["Authorization"] = kp.Authorization

	pathURL := "/api/order/callback_cancel_merchant"
	//Marshal Payload
	payload, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	err = kp.call("POST", pathURL, bytes.NewBuffer(payload), &response, headers)
	if err != nil {
		return response, err
	}

	if response.Code != http.StatusOK {
		err = errors.New(response.Messages)
		return response, err
	}

	return response, nil
}

func GenerateServiceResult(data interface{}, err error) ServiceResult {
	var output ServiceResult
	output = ServiceResult{Result: data, Error: err}
	return output
}
