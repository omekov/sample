package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/omekov/sample/pkg/domain"
)

type UserRepo struct {
	conn *pgx.Conn
}

func NewUser(conn *pgx.Conn) *UserRepo {
	return &UserRepo{
		conn: conn,
	}
}

type Userer interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id uint) (domain.User, error)
	GetByName(ctx context.Context, name string) (domain.User, error)
	GetAll(ctx context.Context, filter domain.UserFilter) ([]domain.User, error)
}

func (r UserRepo) Create(ctx context.Context, user domain.User) error {
	userQuery, args, err := user.InsertScript()
	if err != nil {
		return err
	}

	return r.conn.QueryRow(ctx, userQuery, args...).Scan(&user.ID)
}

func (r UserRepo) GetByID(ctx context.Context, id uint) (domain.User, error) {
	user := domain.User{ID: id}
	sql, args, err := user.SelectOneScript()
	if err != nil {
		return user, err
	}

	if err := r.conn.QueryRow(ctx, sql, args...).Scan(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepo) GetByName(ctx context.Context, name string) (domain.User, error) {
	user := domain.User{Name: name}
	sql, args, err := user.SelectOneByNameScript()
	if err != nil {
		return user, err
	}

	if err := r.conn.QueryRow(ctx, sql, args...).Scan(&user); err != nil {
		return user, err
	}

	return user, nil
}

func (r UserRepo) GetAll(ctx context.Context, filter domain.UserFilter) ([]domain.User, error) {
	users := make([]domain.User, 0)
	user := domain.User{}

	sql, args, err := user.SelectScript(filter.Page)
	if err != nil {
		return users, err
	}
	userRows, err := r.conn.Query(ctx, sql, args...)
	if err != nil {
		return users, err
	}
	for userRows.Next() {
		user = domain.User{}
		if err := userRows.Scan(
			&user.ID,
			&user.Name,
			&user.FirstName,
			&user.BirthDate,
			&user.LastName,
			&user.CreatedAt,
		); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
