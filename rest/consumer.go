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
func HTTPRequest(method string, url string, body interface{}, token *string, setHeaders func(req *http.Request, userID string) *http.Request) (*http.Response, error) {
	headers := []Headers{{Key: "Content-Type", Value: "application/json"}}
	req, _ := getReq(method, url, body)
	for _, value := range headers {
		req.Header.Set(value.Key, value.Value)
	}
	if token != nil {
		req = setHeaders(req, *token)
	}
	client := &http.Client{Timeout: 50 * time.Second}
	resp, err := client.Do(req)
	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    newStr := buf.String()
	message := fmt.Sprintf("end request with url:%s and request %s and response %v", url, body, newStr)
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

//SetHeadersToken set headers braze
func SetHeadersToken(req *http.Request, token string) *http.Request {
	req.Header.Set("AUTHORIZATION", token)
	return req
}
