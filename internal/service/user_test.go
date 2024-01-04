package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/service"
	"testing"
)

func prepareUser(t *testing.T, authenticationSrv *service.AuthenticationService) *model.User {
	ctx := context.Background()
	authentication, user, err := authenticationSrv.SignUp(
		ctx,
		"google",
		uuid.New().String(),
		uuid.New().String(),
		"new_user",
	)
	if err != nil {
		t.Fatal(err)
	}

	if authentication == nil {
		t.Fatal("authentication is nil")
	}

	if user == nil {
		t.Fatal("user is nil")
	}

	return user
}

func TestUserService_GetUser(t *testing.T) {
	tester := Initialize(t)
	srv := tester.UserSrv
	targetUser := prepareUser(t, tester.AuthenticationSrv)

	t.Run("should return a user", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		u, err := srv.GetByID(
			ctx,
			targetUser.ID,
		)
		if err != nil {
			t.Fatal(err)
		}

		if u == nil {
			t.Fatal("u is nil")
		}

		if u.ID == "" {
			t.Fatal("u.ID is empty")
		}
	})
}

func TestUserService_EditUser(t *testing.T) {
	tester := Initialize(t)
	srv := tester.UserSrv
	targetUser := prepareUser(t, tester.AuthenticationSrv)
	t.Run("should update a user", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()

		u, err := srv.Update(
			ctx,
			targetUser.ID,
			&model.User{
				Name:     "new name",
				Timezone: "new timezone",
			},
		)
		if err != nil {
			t.Fatal(err)
		}

		if u == nil {
			t.Fatal("u is nil")
		}

		if u.ID == "" {
			t.Fatal("u.ID is empty")
		}

		if u.Name != "new name" {
			t.Fatal("u.Name is not new name")
		}

		if u.Timezone != "new timezone" {
			t.Fatal("u.Timezone is not new timezone")
		}
	})
}
