package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/app"
	"testing"
)

func Initialize(t *testing.T) app.Tester {
	tester := app.InitializeTestApplication()
	tester.ResetTables(t)
	return tester
}

func TestAuthenticationService_SignUp(t *testing.T) {
	tester := Initialize(t)
	defer tester.CleanUp(t)
	srv := tester.AuthenticationSrv

	t.Run("should create a new user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		authentication, user, err := srv.SignUp(
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
	})
}

func TestAuthenticationService_SignInOrSignUp(t *testing.T) {
	tester := Initialize(t)
	defer tester.CleanUp(t)
	srv := tester.AuthenticationSrv

	t.Run("should create a new authentication and user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		authentication, user, err := srv.SignInOrSignUp(
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

		retryAuthentication, retryUser, err := srv.SignInOrSignUp(
			ctx,
			"google",
			authentication.Identifier,
			uuid.New().String(),
			"new_user",
		)
		if err != nil {
			t.Fatal(err)
		}

		if retryAuthentication == nil {
			t.Fatal("retryAuthentication is nil")
		}

		if retryUser == nil {
			t.Fatal("retryUser is nil")
		}

		if retryAuthentication.ID != authentication.ID {
			t.Fatalf("retryAuthentication.ID: %s != authentication.ID: %s", retryAuthentication.ID, authentication.ID)
		}

		if retryUser.ID != user.ID {
			t.Fatalf("retryUser.ID: %s != user.ID: %s", retryUser.ID, user.ID)
		}
	})
}
