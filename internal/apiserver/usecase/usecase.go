package usecase

import (
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/config"
	"github.com/omekov/sample/pkg/jwt"
	"github.com/omekov/sample/pkg/repository/postgresql"
)

type UseCase struct {
	Auth *Auth
}

func NewUseCase(cfg *config.ENV, store *stores.Store) *UseCase {
	userRepo := postgresql.NewUser(store.PostgresConn)
	_jwt := jwt.NewJWT(*cfg.JWT)
	return &UseCase{
		Auth: NewAuth(userRepo, _jwt),
	}
}
