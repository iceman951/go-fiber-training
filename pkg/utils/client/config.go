package client

import (
	"time"
)

type ClientConfig struct {
	SkipSSLVerify bool
	Timeout       time.Duration
}

var DefaultClientConfig = ClientConfig{
	SkipSSLVerify: true,
	Timeout:       30 * time.Second,
}

type HttpHeaders map[string]string
