package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"

	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type UserService interface {
	Create(req *request.CreateUserRequest) (res *response.LoginUserResponse, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(
	ur repository.UserRepository,
) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (uu userService) Create(req *request.CreateUserRequest) (res *response.LoginUserResponse, err error) {
	user, err := uu.userRepository.Create(req)
	if err != nil {
		return nil, err
	}

	apiToken, err := util.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	r := response.ToLoginUserResponse(user, apiToken)
	res = &r

	return res, nil
}
