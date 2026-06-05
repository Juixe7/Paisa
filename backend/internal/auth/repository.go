package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrSessionNotFound   = errors.New("session not found")
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	
	CreateSession(ctx context.Context, session *Session) error
	GetSession(ctx context.Context, token string) (*Session, error)
	DeleteSession(ctx context.Context, token string) error

	CreateRefreshToken(ctx context.Context, token *RefreshToken) error
	GetRefreshToken(ctx context.Context, tokenStr string) (*RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, tokenStr string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}

type pgRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &pgRepository{pool: pool}
}

func (r *pgRepository) CreateUser(ctx context.Context, user *User) error {
	// SQL placeholder
	return nil
}

func (r *pgRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, ErrUserNotFound
}

func (r *pgRepository) GetUserByID(ctx context.Context, id string) (*User, error) {
	return nil, ErrUserNotFound
}

func (r *pgRepository) UpdateUser(ctx context.Context, user *User) error {
	return nil
}

func (r *pgRepository) CreateSession(ctx context.Context, session *Session) error {
	return nil
}

func (r *pgRepository) GetSession(ctx context.Context, token string) (*Session, error) {
	return nil, ErrSessionNotFound
}

func (r *pgRepository) DeleteSession(ctx context.Context, token string) error {
	return nil
}

func (r *pgRepository) CreateRefreshToken(ctx context.Context, token *RefreshToken) error {
	return nil
}

func (r *pgRepository) GetRefreshToken(ctx context.Context, tokenStr string) (*RefreshToken, error) {
	return nil, errors.New("token not found")
}

func (r *pgRepository) DeleteRefreshToken(ctx context.Context, tokenStr string) error {
	return nil
}

func (r *pgRepository) DeleteUserRefreshTokens(ctx context.Context, userID string) error {
	return nil
}
