package utils

import (
	"crypto/tls"
	"net/http"
)

func MakeRequest(url string) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return http.Get(url)
}
