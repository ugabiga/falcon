package service_test

import (
	"context"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/config"
	"github.com/ugabiga/falcon/internal/service"
	"testing"
)

func initAuthenticationService() *service.AuthenticationService {
	cfg := &config.Config{
		DBDriverName: "sqlite3",
		DBSource:     "file:ent?mode=memory&cache=shared&_fk=1",
	}
	entClient := client.NewEntClient(cfg)

	return service.NewAuthenticationService(entClient, cfg)
}
func TestAuthenticationService_SignUp(t *testing.T) {
	srv := initAuthenticationService()

	t.Run("should create a new user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		authentication, err := srv.SignUp(
			ctx,
			"google",
			"",
			"",
			"test",
		)
		if err != nil {
			t.Fatal(err)
		}

		if authentication == nil {
			t.Fatal("authentication is nil")
		}

		if authentication.ID == 0 {
			t.Fatal("authentication.ID is 0")
		}

		if authentication.Edges.User == nil {
			t.Fatal("authentication.Edges.User is nil")
		}
	})
}

func TestAuthenticationService_SignInOrSignUp(t *testing.T) {
	srv := initAuthenticationService()

	t.Run("should create a new user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		authentication, err := srv.SignInOrSignUp(
			ctx,
			"google",
			"",
			"",
			"test",
		)
		if err != nil {
			t.Fatal(err)
		}

		if authentication == nil {
			t.Fatal("authentication is nil")
		}

		if authentication.ID == 0 {
			t.Fatal("authentication.ID is 0")
		}

		if authentication.Edges.User == nil {
			t.Fatal("authentication.Edges.User is nil")
		}
	})
}
