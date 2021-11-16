package usecase

import (
	"context"

	"github.com/omekov/sample/internal/model"
	"github.com/omekov/sample/pkg/jwt"
	"github.com/omekov/sample/pkg/repository/postgresql"
)

type Auth struct {
	userRepo postgresql.Userer
	jwt      *jwt.Config
}

func NewAuth(userRepo *postgresql.UserRepo, jwt *jwt.Config) *Auth {
	return &Auth{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (uc Auth) SignIn(ctx context.Context, credential *model.Credential) (model.Token, error) {
	token := model.Token{}
	if err := uc.jwt.Validate(*credential); err != nil {
		return token, err
	}

	user, err := uc.userRepo.GetByName(ctx, credential.Username)
	if err != nil {
		return token, err
	}

	if err = uc.jwt.ComparePassword(user.Password, credential.Password); err != nil {
		return token, err
	}

	return uc.jwt.GetToken(&user)
}
