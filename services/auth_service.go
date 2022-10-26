package services

import (
	"encoding/json"
	"fmt"

	"github.com/klover2/qbit-go-sdk/client"
)

type AuthService struct {
	client       *client.Client
	baseUrl      string
	clientId     string
	clientSecret string
}

type ErrOutput struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GetCodeOutput struct {
	Timestamp int    `json:"timestamp"`
	State     string `json:"state"`
	Code      string `json:"code"`
}

type GetAccessTokenOutput struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	Timestamp    int    `json:"timestamp"`
}

type RefreshAccessTokenOutput struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
	Timestamp   int    `json:"timestamp"`
}

// NewRequestService 初始化service
func NewAuthService(clientId string, clientSecret string, baseUrl string) *AuthService {
	svc := &AuthService{client: client.NewClient(), clientId: clientId, clientSecret: clientSecret, baseUrl: baseUrl}
	return svc
}

// 获取code
func (s *AuthService) GetCode(state string, redirectUri string) (GetCodeOutput, error) {
	url := s.baseUrl + "/open-api/oauth/authorize"
	query := map[string]interface{}{
		"clientId":    s.clientId,
		"state":       state,
		"redirectUri": redirectUri,
	}
	var data GetCodeOutput

	res, err := s.client.Get(url, query, nil)
	if err != nil {
		return data, err
	}

	if res.Status >= 200 && res.Status < 300 {

		err := json.Unmarshal([]byte(res.Content), &data)
		if err != nil {
			return data, err
		}
		return data, nil
	} else {
		var errMessage ErrOutput
		err := json.Unmarshal([]byte(res.Content), &errMessage)
		if err != nil {
			return data, err
		}
		return data, fmt.Errorf(errMessage.Message)
	}
}

// 获取access token
func (s *AuthService) GetAccessToken(code string) (GetAccessTokenOutput, error) {
	url := s.baseUrl + "/open-api/oauth/access-token"
	body := map[string]interface{}{
		"clientId":     s.clientId,
		"clientSecret": s.clientSecret,
		"code":         code,
	}
	var data GetAccessTokenOutput
	res, err := s.client.Post(url, body, nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(res.Content), &data)
	if err != nil {
		return data, err
	}
	return data, nil

}

// 刷新access token
func (s *AuthService) RefreshToken(refreshToken string) (RefreshAccessTokenOutput, error) {
	url := s.baseUrl + "/open-api/oauth/refresh-token"
	body := map[string]interface{}{
		"clientId":     s.clientId,
		"refreshToken": refreshToken,
	}
	var data RefreshAccessTokenOutput
	res, err := s.client.Post(url, body, nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(res.Content), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
