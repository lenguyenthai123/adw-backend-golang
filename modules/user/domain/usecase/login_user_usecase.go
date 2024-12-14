package usecase

import (
	"backend-golang/api/middlewares/constant"
	"backend-golang/modules/user/domain/entity"
	"backend-golang/pkgs/log"
	"context"
	"log/slog"
)

type LoginUserUseCase interface {
	ExecLoginUser(ctx context.Context, userEntity entity.UserEntity) (*entity.UserEntity, error)
}

type loginUserUseCaseImpl struct {
	tokenProvider    TokenProvider
	accessTokenTime  int
	refreshTokenTime int
	hashAlgo         HashAlgo
	readerRepo       userReaderRepository
}

var _ LoginUserUseCase = (*loginUserUseCaseImpl)(nil)

func NewLoginUserUseCase(
	tokenProvider TokenProvider,
	accessTokenTime int,
	refreshTokenTime int,
	hashAlgo HashAlgo,
	readerRepo userReaderRepository) LoginUserUseCase {
	return &loginUserUseCaseImpl{
		tokenProvider:    tokenProvider,
		accessTokenTime:  accessTokenTime,
		refreshTokenTime: refreshTokenTime,
		hashAlgo:         hashAlgo,
		readerRepo:       readerRepo,
	}
}

func (useCase loginUserUseCaseImpl) ExecLoginUser(ctx context.Context,
	userEntity entity.UserEntity) (*entity.UserEntity, error) {
	// Check if user exists
	user, err := useCase.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"email": userEntity.Email,
	})
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.email_not_found",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrEmailNotFound(err)
	}

	// Check if password is correct
	if err := useCase.hashAlgo.ComparePasswords(user.Password, []byte(userEntity.Password)); err != nil {
		log.JsonLogger.Error("ExecLoginUser.password_not_match",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrEmailOrPasswordInvalid(err)
	}

	// Generate access token
	accessToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{
			"user_id": user.ID,
			"role":    "USER",
		},
		useCase.accessTokenTime,
	)
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.generate_access_token",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrInternal(err)
	}

	// Generate refresh token
	refreshToken, err := useCase.tokenProvider.Generate(
		map[string]interface{}{
			"user_id": user.ID,
			"role":    "USER",
		},
		useCase.refreshTokenTime,
	)
	if err != nil {
		log.JsonLogger.Error("ExecLoginUser.generate_refresh_token",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return nil, constant.ErrInternal(err)
	}

	user.AccessToken = accessToken["token"].(string)
	user.RefreshToken = refreshToken["token"].(string)

	return user, nil
}
