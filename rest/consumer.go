package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dasuma/blitz-utils-go/utils"
	"github.com/labstack/gommon/log"
)

var user101 = utils.GetEnv("user_101", "CUE")
var password101 = utils.GetEnv("password_101", "123456")
var errorCode = []int{
	400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 425, 426,
	428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}

//HTTPRequest request
func HTTPRequest(method string, url string, body interface{}, result interface{}, headers map[string]string) (interface{}, error) {
	req, _ := getReq(method, url, body)
	req = setHeaders(req, headers)
	client := &http.Client{Timeout: time.Duration(1) * time.Minute}
	resp, err := client.Do(req)
	if resp != nil {
		if ContainsInt(errorCode, resp.StatusCode) {
			responseData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf(err.Error())
			}
			message := fmt.Sprintf("error request with url:%s and request %s and code_rescode:%v response %v", url, body, resp.StatusCode, string(responseData))
			log.Info(message)
			return nil, errors.New(message)
		}
	} else {
		message := fmt.Sprintf("error request with url:%s and request %s and code_rescode:%v response %v", url, body, nil, nil)
		log.Info(message)
		return nil, errors.New(message)
	}
	json.NewDecoder(resp.Body).Decode(&result)
	message := fmt.Sprintf("end request with url:%s and request %s and code_rescode:%v response %v", url, body, resp.StatusCode, result)
	log.Info(message)
	return result, err
}
func getReq(method string, url string, body interface{}) (*http.Request, error) {
	if method == http.MethodGet {
		return http.NewRequest(method, url, nil)
	}
	jsonReq, _ := json.Marshal(body)
	request := bytes.NewReader(jsonReq)
	return http.NewRequest(method, url, request)

}

//SetHeaders set headers
func setHeaders(req *http.Request, headers map[string]string) *http.Request {
	for key, element := range headers {
		req.Header.Set(key, element)
	}
	return req
}
