package kreditplus

import (
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

const (
    BaseURL        = "URL"
    Authorization  = "Auth"
    DefaultTimeout = 10 * time.Second
)

func TestNew(t *testing.T) {
    _ = New(BaseURL, Authorization, DefaultTimeout)
    assert.NoError(t, nil)
}

func TestGenerateServiceResult(t *testing.T) {
    _ = GenerateServiceResult(nil, nil)
    assert.NoError(t, nil)
}

func TestKreditplus_GetCheckoutURL(t *testing.T) {

}
