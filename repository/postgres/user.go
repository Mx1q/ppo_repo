package postgres

import (
	"context"
	"fmt"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `insert into saladRecipes.user(name, email, login, password)
		values ($1, $2, $3, $4)`

	_, err := r.db.Exec(
		ctx,
		query,
		user.Name,
		user.Email.Address,
		user.Username,
		user.Password,
	)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	return nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `select id, name, email, login, password, role
		from saladRecipes.user
		where id = $1`

	user := new(domain.User)
	//user.Email = mail.Address{}
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email.Address,
		&user.Username,
		&user.Password,
		&user.Role)

	if err != nil {
		return nil, fmt.Errorf("getting user by id: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context, page int) ([]*domain.User, error) {
	query := `select id, name, email, login, password, role
		from saladRecipes.user
		offset $1
		limit $2`

	rows, err := r.db.Query(
		ctx,
		query,
		PageSize*(page-1),
		PageSize,
	)
	if err != nil {
		return nil, fmt.Errorf("getting users: %w", err)
	}

	users := make([]*domain.User, 0)
	for rows.Next() {
		tmp := new(domain.User)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Email.Address,
			&tmp.Username,
			&tmp.Password,
			&tmp.Role,
		)
		users = append(users, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning users: %w", err)
		}
	}

	return users, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	query := `update saladRecipes.user
		set
			name = $1,
			email = $2,
			password = $3,
			role = $5
		where id = $4`

	_, err := r.db.Exec(
		ctx,
		query,
		user.Name,
		user.Email.Address,
		user.Password,
		user.Role,
	)
	if err != nil {
		return fmt.Errorf("updating user: %w", err)
	}
	return nil
}

func (r *userRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.user
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting user by id: %w", err)
	}
	return nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `select id, name, email, login, role
		from saladRecipes.user
		where login = $1`

	user := new(domain.User)
	err := r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email.Address,
		&user.Username,
		&user.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("getting user by username: %w", err)
	}
	return user, nil
}
