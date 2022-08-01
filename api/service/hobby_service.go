package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
)

type HobbyService interface {
	FindAll() (res []*response.HobbyResponse, err error)
}

type hobbyService struct {
	hobbyRepository repository.HobbyRepository
}

func NewHobbyService(hr repository.HobbyRepository) HobbyService {
	return &hobbyService{
		hobbyRepository: hr,
	}
}

func (hs hobbyService) FindAll() (res []*response.HobbyResponse, err error) {
	hobbies, err := hs.hobbyRepository.FindAll()
	if err != nil {
		return nil, err
	}
	res = []*response.HobbyResponse{}
	for _, h := range hobbies {
		r := &response.HobbyResponse{}
		r.ToHobbyResponse(h)
		res = append(res, r)
	}

	return res, nil
}
