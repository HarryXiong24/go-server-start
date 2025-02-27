package api

type GetUserInfoRequest struct {
	Name string `form:"name"`
}

type GetUserInfoResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
