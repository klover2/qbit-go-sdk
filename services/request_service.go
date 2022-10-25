package services

import (
	"github.com/klover2/qbit-go-sdk/client"
)

type RequestService Service

// NewRequestService 初始化service
func (s RequestService) NewRequestService(clientId string, clientSecret string) *RequestService {
	return &RequestService{Client: client.NewClient(clientId, clientSecret)}
}

// GetBaseUrl 获取客户端baseUrl
//func (s RequestService) GetBaseUrl() string {
//	return s.baseUrl
//}

// SetBaseUrl 设置客户端baseUrl
//func (s RequestService) SetBaseUrl(baseUrl string) {
//	s.baseUrl = baseUrl
//}

func (s RequestService) PostRequest(url string, body interface{}) (interface{}, error) {
	res, err := s.Client.Post(url, body, nil)
	return res, err
}

func (s RequestService) GetRequest(url string, query map[string]string) (interface{}, error) {
	var queryStr string = ""

	for k, v := range query {
		if queryStr == "" {
			queryStr = queryStr + k + "=" + v
		} else {
			queryStr = queryStr + "&" + k + "=" + v
		}
	}

	if queryStr != "" {
		queryStr = "?" + queryStr
	}
	
	res, err := s.Client.Get(url+queryStr, nil)
	return res, err
}

func (s RequestService) DeleteRequest(url string, body interface{}) (interface{}, error) {
	res, err := s.Client.Delete(url, body, nil)
	return res, err
}

func (s RequestService) PutRequest(url string, body interface{}) (interface{}, error) {
	res, err := s.Client.Put(url, body, nil)
	return res, err
}
