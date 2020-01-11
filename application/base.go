package application

import (
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
)

//BaseResponse is the basic struct of weibo api response object
type BaseResponse struct {
	Error string `json:"error"`
	ErrorCode int `json:"error_code"`
	Request string `json:"request"`
	OriginalBytes []byte
	NoError bool
}

func handleResponse(resp *http.Response)(*BaseResponse, error) {
	defer resp.Body.Close()
	resBytes, _ := ioutil.ReadAll(resp.Body)
	response := &BaseResponse{OriginalBytes: resBytes}
	if strings.Index(string(resBytes), "error") < 0 {
		response.NoError = true
		return response, nil
	}
	response.NoError = false
	if err := json.Unmarshal(resBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

//SimpleGet send a simple http-get request
func SimpleGet(uri string, params ...string) (*BaseResponse, error ) {
	for _, param := range params {
			uri += param
	}
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return handleResponse(resp)
}

//SimplePost send a simple http-post request
func SimplePost(uri string, params map[string]string) (*BaseResponse, error) {
	data := make(map[string][]string)
	for key, value := range params {
		data[key] = []string{value}
	}
	resp, err := http.PostForm(uri, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return handleResponse(resp)
}