package response

import (
	"github.com/hiroki-Fukumoto/matching-app-api/api/model"
)

type HobbyResponse struct {
	ID   string `json:"id" validate:"required"`   // ID
	Name string `json:"name" validate:"required"` // 名称
}

func (h *HobbyResponse) ToHobbyResponse(m *model.Hobby) HobbyResponse {
	h.ID = m.ID
	h.Name = m.Name

	return *h
}
