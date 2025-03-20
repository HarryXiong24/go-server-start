package api

// UserInfo is a common response type for user data
type UserInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GetUserInfoRequest defines the request for getting user info
type GetUserInfoRequest struct {
	Name string `form:"name" binding:"required"`
}

// GetUserInfoResponse defines the response for getting user info
type GetUserInfoResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CreateUserRequest defines the request for creating a user
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateUserResponse defines the response for creating a user
type CreateUserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// UpdateUserRequest defines the request for updating a user
type UpdateUserRequest struct {
	ID   int64  `uri:"id" binding:"required,gt=0"`
	Name string `json:"name" binding:"required"`
}

// UpdateUserResponse defines the response for updating a user
type UpdateUserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// DeleteUserRequest defines the request for deleting a user
type DeleteUserRequest struct {
	ID int64 `uri:"id" binding:"required,gt=0"`
}

// ListUsersRequest defines the request for listing users
type ListUsersRequest struct {
	Page     int `form:"page,default=1" binding:"min=1"`
	PageSize int `form:"page_size,default=10" binding:"min=1,max=100"`
}

// ListUsersResponse defines the response for listing users
type ListUsersResponse struct {
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Users []*UserInfo `json:"users"`
}
