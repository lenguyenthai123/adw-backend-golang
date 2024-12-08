package repository

import (
	"backend-golang/modules/user/domain/entity"
	"context"
)

type UserReaderRepository interface {
	FindUserByCondition(ctx context.Context, condition map[string]interface{}) (*entity.UserEntity, error)
}

type UserWriterRepository interface {
	InsertUser(ctx context.Context, userEntity entity.UserEntity) error
	DeleteUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, userEntity entity.UserEntity) error
}
