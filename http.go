package kreditplus

import (
    "encoding/json"
    "io"
    "net/http"
    "time"
)

// httpRequest data model
type kreditplusHttpClient struct {
    httpClient *http.Client
}

// newRequest function for intialize httpRequest object
// Paramter, timeout in time.Duration
func newRequest(timeout time.Duration) *kreditplusHttpClient {
    return &kreditplusHttpClient{
        httpClient: &http.Client{Timeout: time.Second * timeout},
    }
}

// newReq function for initalize http request,
// paramters, http method, uri path, body, and headers
func (kp *kreditplusHttpClient) newReq(method string, fullPath string, body io.Reader, headers map[string]string) (*http.Request, error) {
    req, err := http.NewRequest(method, fullPath, body)
    if err != nil {
        return nil, err
    }

    for key, value := range headers {
        req.Header.Set(key, value)
    }

    return req, nil
}

// exec private function for call http request
func (kp *kreditplusHttpClient) exec(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
    req, err := kp.newReq(method, path, body, headers)

    if err != nil {
        return err
    }

    res, errHttpClient := kp.httpClient.Do(req)
    if errHttpClient != nil {
        return err
    }

    defer res.Body.Close()

    if v != nil {
        return json.NewDecoder(res.Body).Decode(v)
    }

    return nil
}
