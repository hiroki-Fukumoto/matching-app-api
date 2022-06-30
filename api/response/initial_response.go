package response

type InitialResponse struct {
	MinVersion     string `json:"min_version" validate:"required"`     // min version
	CurrentVersion string `json:"current_version" validate:"required"` // current version
}

// TODO
func (i *InitialResponse) ToInitialResponse() InitialResponse {
	i.MinVersion = "1.0.0"
	i.CurrentVersion = "1.0.1"

	return *i
}
