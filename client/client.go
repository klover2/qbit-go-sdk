package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*Client HTTP客户端*/
type Client struct {
	httpClient   *http.Client
	clientId     string
	clientSecret string
}

// APIResult  请求结果
type APIResult struct {
	StatusCode int
	Status     string
	Content    string
}

/*NewClient 新建http客户端*/
func NewClient(clientId string, clientSecret string) *Client {
	client := &Client{
		clientId:     clientId,
		clientSecret: clientSecret,
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	return client
}

// Get 发送一个 HTTP Get 请求
func (client *Client) Get(requestURL string, header http.Header) (*APIResult, error) {
	return client.doRequest(http.MethodGet, requestURL, nil, header)
}

// Post 发送一个 HTTP Post 请求
func (client *Client) Post(requestURL string, requestBody interface{}, header http.Header) (*APIResult, error) {
	return client.requestWithJSONBody(http.MethodPost, requestURL, requestBody, header)
}

// Patch 发送一个 HTTP Patch 请求
func (client *Client) Patch(requestURL string, requestBody interface{}, header http.Header) (*APIResult, error) {
	return client.requestWithJSONBody(http.MethodPatch, requestURL, requestBody, header)
}

// Put 发送一个 HTTP Put 请求
func (client *Client) Put(requestURL string, requestBody interface{}, header http.Header) (*APIResult, error) {
	return client.requestWithJSONBody(http.MethodPut, requestURL, requestBody, header)
}

// Delete 发送一个 HTTP Delete 请求
func (client *Client) Delete(requestURL string, requestBody interface{}, header http.Header) (*APIResult, error) {
	return client.requestWithJSONBody(http.MethodDelete, requestURL, requestBody, header)
}

// 处理 body 参数
func (client *Client) requestWithJSONBody(method, requestURL string, body interface{}, header http.Header) (
	*APIResult, error,
) {
	var (
		reqBody *strings.Reader
	)

	if body == nil {
		reqBody = nil
	} else {
		var (
			stringBody string
			ok         bool
		)
		if stringBody, ok = body.(string); ok == false {
			dataType, err := json.Marshal(body)

			if err == nil {
				return nil, err
			}
			stringBody = string(dataType)
		}
		reqBody = strings.NewReader(stringBody)
	}

	return client.doRequest(method, requestURL, reqBody, header)
}

// 发送请求
func (client *Client) doRequest(
	method string,
	requestURL string,
	reqBody io.Reader,
	header http.Header,
) (*APIResult, error) {
	var (
		err     error
		request *http.Request
	)

	if request, err = http.NewRequest(method, requestURL, reqBody); err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Content-Type", "application/json")

	// Add Request Header Parameters
	for key, values := range header {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// Send HTTP Request
	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	// deal Response
	res, err := dealResponse(response)

	return res, err
}

// 处理错误信息
func dealResponse(response *http.Response) (*APIResult, error) {
	slurp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid response, read body error: %w", err)
	}
	_ = response.Body.Close()

	response.Body = ioutil.NopCloser(bytes.NewBuffer(slurp))

	res := &APIResult{
		StatusCode: response.StatusCode,
		Status:     strings.TrimSpace(strings.Replace(response.Status, "200", "", -1)),
		Content:    string(slurp),
	}

	return res, nil
}
