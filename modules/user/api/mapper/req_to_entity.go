package mapper

import (
	"backend-golang/modules/user/api/model/req"
	"backend-golang/modules/user/domain/entity"
)

func ConvertCreateUserReqToUserEntity(req req.CreateUserReq) entity.UserEntity {
	return entity.UserEntity{
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
		Status:   entity.UNVERIFIED.Value(),
	}
}

func ConvertLoginUserReqToUserEntity(req req.LoginUserReq) entity.UserEntity {
	return entity.UserEntity{
		Email:    req.Email,
		Password: req.Password,
	}
}

func ConvertUpdateUserReqToUserEntity(req req.UpdateUserReq) entity.UserEntity {
	return entity.UserEntity{
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
	}
}
