package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dasuma/blitz-utils-go/logs"
	"github.com/dasuma/blitz-utils-go/utils"
)

var user101 = utils.GetEnv("user_101", "CUE")
var password101 = utils.GetEnv("password_101", "123456")

//HTTPRequest request
func HTTPRequest(method string, url string, headers []Headers, body interface{}, token *string, setHeaders func(req *http.Request, userID string) *http.Request) (*http.Response, error) {
	req, _ := getReq(method, url, body)
	for _, value := range headers {
		req.Header.Set(value.Key, value.Value)
	}
	if token != nil {
		req = setHeaders(req, *token)
	}
	client := &http.Client{Timeout: 50 * time.Second}
	resp, err := client.Do(req)
	message := fmt.Sprintf("end request with url:%s and request %s and response %v", url, body, resp)
	logs.Logger.Info(message)
	return resp, err
}
func getReq(method string, url string, body interface{}) (*http.Request, error) {
	if method == http.MethodGet {
		return http.NewRequest(method, url, nil)
	} else {
		jsonReq, _ := json.Marshal(body)
		request := bytes.NewReader(jsonReq)
		return http.NewRequest(method, url, request)
	}
}

//SetHeaders101 set headers braze
func SetHeaders101(req *http.Request, token string) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AUTHORIZATION", token)
	return req
}
