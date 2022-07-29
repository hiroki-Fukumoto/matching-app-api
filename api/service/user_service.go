package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
	"golang.org/x/crypto/bcrypt"

	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type UserService interface {
	Create(req *request.CreateUserRequest) (res *response.LoginUserResponse, err error)
	PickupToday(targetSex string) (res []*response.UserResponse, err error)
	FindAll(req *request.SearchUserRequest, loginUserID string) (res []*response.UserResponse, err error)
	FindByID(id string, loginUserID string) (res *response.UserResponse, err error)
}

type userService struct {
	userRepository     repository.UserRepository
	favoriteRepository repository.FavoriteRepository
}

func NewUserService(
	ur repository.UserRepository,
	fr repository.FavoriteRepository,
) UserService {
	return &userService{
		userRepository:     ur,
		favoriteRepository: fr,
	}
}

func (us userService) Create(req *request.CreateUserRequest) (res *response.LoginUserResponse, err error) {
	r := us.userRepository.CreateRequest()
	r.Name = req.Name
	r.Email = req.Email
	r.Sex = req.Sex
	r.Birthday = util.ParseDate(req.Birthday)
	r.Prefecture = uint16(req.Prefecture)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	r.Password = passwordHash

	user, err := us.userRepository.Create(r)
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
func (us userService) PickupToday(targetSex string) (res []*response.UserResponse, err error) {
	r := us.userRepository.FindPickUpTodayRequest()
	r.Sex = targetSex
	users, err := us.userRepository.FindPickUpToday(r)
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

func (us userService) FindAll(req *request.SearchUserRequest, loginUserID string) (res []*response.UserResponse, err error) {
	r := us.userRepository.FindAllRequest()
	r.Page = req.Page
	r.FromAge = req.FromAge
	r.ToAge = req.ToAge
	r.Prefectures = req.Prefectures
	r.Sort = req.Sort
	users, err := us.userRepository.FindAll(r)
	if err != nil {
		return nil, err
	}

	sreq := us.favoriteRepository.FindSendLikesRequest()
	sreq.SenderID = loginUserID
	sendLikes, err := us.favoriteRepository.FindSendLikes(sreq)
	if err != nil {
		return nil, err
	}

	res = []*response.UserResponse{}
	for _, u := range users {
		r := &response.UserResponse{}
		r.ToUserResponse(u)
		for _, s := range sendLikes {
			if s.ReceiverID == u.ID {
				r.IsLiked = true
				break
			}
		}
		res = append(res, r)
	}
	return res, nil
}

func (us userService) FindByID(id string, loginUserID string) (res *response.UserResponse, err error) {
	r := us.userRepository.FindByIDRequest()
	r.ID = id
	user, err := us.userRepository.FindByID(r)
	if err != nil {
		return nil, err
	}

	sreq := us.favoriteRepository.FindSendLikesRequest()
	sreq.SenderID = loginUserID
	sendLikes, err := us.favoriteRepository.FindSendLikes(sreq)
	if err != nil {
		return nil, err
	}

	res = &response.UserResponse{}
	res.ToUserResponse(user)

	for _, s := range sendLikes {
		if s.ReceiverID == user.ID {
			res.IsLiked = true
			break
		}
	}

	if id == loginUserID {
		res.IsMySelf = true
	}

	return res, nil
}
