package services

import (
	"go-server-start/api"
	"go-server-start/internal/repositories"
)

func GetUserInfo(req *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	var res api.GetUserInfoResponse

	if req.Name != "" {
		userInfo, err := repositories.GetUserInfo(req.Name)

		if err != nil {
			return nil, err
		}

		res.ID = userInfo.ID
		res.Name = userInfo.Name
	}

	return &res, nil
}
