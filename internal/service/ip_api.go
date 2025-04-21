package service

import (
	"context"
	"fmt"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/domain"
	"net/http"
	"time"

	httpchain "github.com/chenmingyong0423/go-http-chain"
)

func NewIpApiService() *IpApiService {
	return &IpApiService{
		host: "http://ip-api.com",
		client: httpchain.NewWithClient(&http.Client{
			Timeout: time.Second * 10,
		}),
	}
}

type IIpApiService interface {
	GetLocation(ctx context.Context, ip string) (*domain.IpApiResponse, error)
}

var _ IIpApiService = (*IpApiService)(nil)

type IpApiService struct {
	host   string
	client *httpchain.Client
}

func (s *IpApiService) GetLocation(ctx context.Context, ip string) (*domain.IpApiResponse, error) {
	var resp domain.IpApiResponse
	err := s.client.Get(fmt.Sprintf("%s/json/%s", s.host, ip)).DoAndParse(ctx, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
