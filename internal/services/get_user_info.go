package services

import (
	"context"
	"go-server-start/internal/repositories"
	api "go-server-start/internal/types"
	"go-server-start/pkg/errors"
)

// GetUserInfo retrieves user information with context
func GetUserInfo(ctx context.Context, req *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	// Validate request
	if req.Name == "" {
		return nil, errors.NewBadRequest("name is required", nil)
	}

	// Get user from repository
	user, err := repositories.GetUserByName(ctx, req.Name)
	if err != nil {
		return nil, err // Repository already wraps errors
	}

	// Map to response
	return &api.GetUserInfoResponse{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
