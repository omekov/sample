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
	repo := postgresql.NewRepositories(store.PostgresConn)
	jwt := jwt.NewJWT(*cfg.JWT)
	return &UseCase{
		Auth: NewAuth(repo, _jwt),
	}
}
