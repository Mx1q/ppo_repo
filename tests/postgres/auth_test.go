package tests

import (
	"context"
	"errors"
	"github.com/Mx1q/ppo_repo/repository/postgres"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/stretchr/testify/require"
	"net/mail"
	"testing"
)

func TestAuthRepository_Register(t *testing.T) {
	repo := postgres.NewAuthRepository(testDbInstance)

	testCases := []struct {
		name     string
		authInfo *domain.User
		wantErr  bool
		errStr   error
	}{
		{
			name: "успешная регистрация",
			authInfo: &domain.User{
				Name:     "testingUser",
				Username: "testingUser",
				Password: "testingUser",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			wantErr: false,
		}, // успешная регистрация
		{
			name: "неуникальное имя пользователя",
			authInfo: &domain.User{
				Name:     "testingUser",
				Username: "testingUser",
				Password: "testingUser",
				Email: mail.Address{
					Name:    "",
					Address: "anotherTest@mail.ru",
				},
			},
			wantErr: true,
			errStr: errors.New("user registration: ERROR: duplicate key value " +
				"violates unique constraint \"user_login_key\" (SQLSTATE 23505)"),
		}, // неуникальное имя пользователя
		{
			name: "неуникальная почта",
			authInfo: &domain.User{
				Name:     "testingUser",
				Username: "anotherTestingUser",
				Password: "testingUser",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			wantErr: true,
			errStr: errors.New("user registration: ERROR: duplicate key value " +
				"violates unique constraint \"user_email_key\" (SQLSTATE 23505)"),
		}, // неуникальная почта
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Register(context.Background(), tt.authInfo)

			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestAuthRepository_GetByUsername(t *testing.T) {
	repo := postgres.NewAuthRepository(testDbInstance)

	testCases := []struct {
		name     string
		username string
		expected *domain.UserAuth
		wantErr  bool
		errStr   error
	}{
		{
			name:     "пользователь существует",
			username: "testingUser",
			expected: &domain.UserAuth{
				Username:   "testingUser",
				HashedPass: "testingUser",
			},
			wantErr: false,
		}, // пользователь существует
		{
			name:     "пользователь не найден",
			username: "testingUserNotFound",
			wantErr:  true,
			errStr:   errors.New("getting user by username: no rows in result set"),
		}, // пользователь не найден
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := repo.GetByUsername(context.Background(), tt.username)

			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected.HashedPass, res.HashedPass)
			}
		})
	}
}
