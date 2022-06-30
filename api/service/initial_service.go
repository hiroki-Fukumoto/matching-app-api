package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type InitialService interface {
	Initial() (res *response.InitialResponse, err error)
}

type initialService struct{}

func NewInitialService() InitialService {
	return &initialService{}
}

func (is initialService) Initial() (res *response.InitialResponse, err error) {
	r := response.ToInitialResponse()
	res = &r
	return
}
