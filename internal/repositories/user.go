package repositories

import (
	"context"
	"go-server-start/internal/models"
	"go-server-start/pkg/database"
	"go-server-start/pkg/errors"
	"go-server-start/pkg/logger"

	"gorm.io/gorm"
)

// GetUserByName retrieves a user by name
func GetUserByName(ctx context.Context, name string) (*models.User, error) {
	if name == "" {
		return nil, errors.NewBadRequest("name cannot be empty", nil)
	}

	var user models.User
	result := database.GetDB().WithContext(ctx).
		Select("id, name").
		Where("name = ?", name).
		First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.NewNotFound("user not found", nil)
		}
		logger.Sugar.Errorw("Failed to query user by name",
			"name", name,
			"error", result.Error.Error(),
		)
		return nil, errors.NewInternalServer("failed to query user", result.Error)
	}

	return &user, nil
}

// CreateUser creates a new user
func CreateUser(ctx context.Context, user *models.User) error {
	return database.WithTransaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(user).Error; err != nil {
			logger.Sugar.Errorw("Failed to create user",
				"user", user,
				"error", err.Error(),
			)
			return errors.NewInternalServer("failed to create user", err)
		}
		return nil
	})
}

// UpdateUser updates an existing user
func UpdateUser(ctx context.Context, user *models.User) error {
	return database.WithTransaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Save(user)
		if result.Error != nil {
			logger.Sugar.Errorw("Failed to update user",
				"user", user,
				"error", result.Error.Error(),
			)
			return errors.NewInternalServer("failed to update user", result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.NewNotFound("user not found", nil)
		}
		return nil
	})
}

// DeleteUser deletes a user by ID
func DeleteUser(ctx context.Context, id int64) error {
	return database.WithTransaction(func(tx *gorm.DB) error {
		result := tx.WithContext(ctx).Delete(&models.User{}, id)
		if result.Error != nil {
			logger.Sugar.Errorw("Failed to delete user",
				"id", id,
				"error", result.Error.Error(),
			)
			return errors.NewInternalServer("failed to delete user", result.Error)
		}
		if result.RowsAffected == 0 {
			return errors.NewNotFound("user not found", nil)
		}
		return nil
	})
}

// ListUsers retrieves a paginated list of users
func ListUsers(ctx context.Context, offset, limit int) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	// Get total count
	if err := database.GetDB().WithContext(ctx).Model(&models.User{}).Count(&total).Error; err != nil {
		logger.Sugar.Errorw("Failed to count users",
			"error", err.Error(),
		)
		return nil, 0, errors.NewInternalServer("failed to count users", err)
	}

	// Get paginated users
	if err := database.GetDB().WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		logger.Sugar.Errorw("Failed to list users",
			"offset", offset,
			"limit", limit,
			"error", err.Error(),
		)
		return nil, 0, errors.NewInternalServer("failed to list users", err)
	}

	return users, total, nil
}
