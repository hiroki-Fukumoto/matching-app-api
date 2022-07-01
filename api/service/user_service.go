package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"

	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type UserService interface {
	Create(req *request.CreateUserRequest) (res *response.LoginUserResponse, err error)
	PickupToday(targetSex string) (res []*response.UserResponse, err error)
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

	res = &response.LoginUserResponse{}
	res.ToLoginUserResponse(user, apiToken)

	return res, nil
}

// TODO: ピックアップ方法
func (uu userService) PickupToday(targetSex string) (res []*response.UserResponse, err error) {
	users, err := uu.userRepository.FindPickUpToday(targetSex)
	if err != nil {
		return nil, err
	}

	res = []*response.UserResponse{}
	for _, u := range users {
		r := &response.UserResponse{}
		r.ToUserResponse(u)
		res = append(res, r)
	}

	return res, nil
}
