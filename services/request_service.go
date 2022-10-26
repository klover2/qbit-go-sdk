package services

import (
	"encoding/json"
	"net/http"

	"github.com/klover2/qbit-go-sdk/client"
)

type RequestService struct {
	client      *client.Client
	accessToken string
}

type ContentOutput struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Output struct {
	Status  int           `json:"status"`
	Reason  string        `json:"reason"`
	Content ContentOutput `json:"content"`
}

// NewRequestService 初始化service
func NewRequestService() *RequestService {
	return &RequestService{client: client.NewClient()}
}

func (s *RequestService) SetAccessToken(accessToken string) {
	s.accessToken = accessToken
}

func (s *RequestService) PostRequest(url string, body interface{}) (Output, error) {
	var content ContentOutput
	res, err := s.client.Post(url, body, http.Header{"x-qbit-access-token": []string{s.accessToken}})
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, nil
	}

	err = json.Unmarshal([]byte(res.Content), &content)
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, err
	}
	return Output{
		Status:  res.Status,
		Reason:  res.Reason,
		Content: content,
	}, nil
}

func (s *RequestService) GetRequest(url string, query map[string]interface{}) (Output, error) {
	var content ContentOutput
	res, err := s.client.Get(url, query, http.Header{"x-qbit-access-token": []string{s.accessToken}})
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, nil
	}

	err = json.Unmarshal([]byte(res.Content), &content)
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, err
	}
	return Output{
		Status:  res.Status,
		Reason:  res.Reason,
		Content: content,
	}, nil
}

func (s *RequestService) DeleteRequest(url string, body interface{}) (Output, error) {
	var content ContentOutput
	res, err := s.client.Delete(url, body, http.Header{"x-qbit-access-token": []string{s.accessToken}})
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, nil
	}

	err = json.Unmarshal([]byte(res.Content), &content)
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, err
	}
	return Output{
		Status:  res.Status,
		Reason:  res.Reason,
		Content: content,
	}, nil
}

func (s *RequestService) PutRequest(url string, body interface{}) (Output, error) {
	var content ContentOutput
	res, err := s.client.Put(url, body, http.Header{"x-qbit-access-token": []string{s.accessToken}})
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, nil
	}

	err = json.Unmarshal([]byte(res.Content), &content)
	if err != nil {
		return Output{
			Status:  res.Status,
			Reason:  res.Reason,
			Content: ContentOutput{},
		}, err
	}
	return Output{
		Status:  res.Status,
		Reason:  res.Reason,
		Content: content,
	}, nil
}
