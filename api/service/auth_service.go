package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/repository"
	"github.com/hiroki-Fukumoto/matching-app-api/api/request"
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(req *request.LoginRequest) (res *response.LoginUserResponse, err error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(ur repository.UserRepository) AuthService {
	return &authService{
		userRepository: ur,
	}
}

func (as authService) Login(req *request.LoginRequest) (res *response.LoginUserResponse, err error) {
	// メールアドレスからユーザー取得
	r := as.userRepository.FindByEmailRequest()
	r.Email = req.Email
	user, err := as.userRepository.FindByEmail(r)
	if err != nil {
		return nil, err
	}
	// パスワードチェック
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
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
