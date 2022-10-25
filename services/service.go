package services

import "github.com/klover2/qbit-go-sdk/client"

// Service API v1 Go SDK 服务类型
type Service struct {
	Client      *client.Client
	baseUrl     string
	accessToken string
}
