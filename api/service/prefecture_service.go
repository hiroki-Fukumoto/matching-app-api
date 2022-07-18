package service

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/response"
	"github.com/hiroki-Fukumoto/matching-app-api/api/util"
)

type PrefectureService interface {
	FindPrefectures() (res []*response.PrefectureResponse)
}

type prefectureService struct{}

func NewPrefectureService() PrefectureService {
	return &prefectureService{}
}

func (ps prefectureService) FindPrefectures() (res []*response.PrefectureResponse) {
	prefes := util.GetPrefectureCodeAndNameList()

	res = []*response.PrefectureResponse{}
	for _, p := range prefes {
		r := &response.PrefectureResponse{
			Code: int(p.Code),
			Name: p.Name,
		}
		res = append(res, r)
	}

	return res
}
