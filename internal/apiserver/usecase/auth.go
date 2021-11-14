package usecase

import (
	"context"

	"github.com/omekov/sample/internal/apiserver/delivery/http"
	"github.com/omekov/sample/pkg/domain"
	"github.com/omekov/sample/pkg/jwt"
	"github.com/omekov/sample/pkg/repository/postgresql"
)

type Auth struct {
	userRepo postgresql.Userer
	jwt      *jwt.Config
}

func NewAuth(repo *postgresql.Repositories, jwt *jwt.Config) *Auth {
	return &Auth{
		userRepo: repo.User,
		jwt:      jwt,
	}
}

func (uc Auth) FindByEmail(ctx context.Context, credential *http.Credential) (domain.User, error) {
	// uc.jwt.NewAccessJWT()
	uc.jwt.RefreshTokenSecret
	return uc.userRepo.GetByName(ctx, credential.Username)
}
