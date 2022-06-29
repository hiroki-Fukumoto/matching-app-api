package service

import (
	"github.com/hiroki-Fukumoto/matching-app/api/response"
)

type InitialService interface {
	Initial() (res response.InitialResponse, err error)
}

type initialService struct{}

func NewInitialService() InitialService {
	return &initialService{}
}

func (is initialService) Initial() (res response.InitialResponse, err error) {
	// TODO
	res = response.ToInitialResponse()
	return
}
