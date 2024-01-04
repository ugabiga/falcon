package service_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/client"
	"github.com/ugabiga/falcon/internal/migration"
	"github.com/ugabiga/falcon/internal/repository"
	"github.com/ugabiga/falcon/internal/service"
	"github.com/ugabiga/falcon/pkg/config"
	"testing"
)

func initAuthenticationService(t *testing.T) *service.AuthenticationService {
	cfg := &config.Config{
		DBDriverName:  "sqlite3",
		DBSource:      "file:ent?mode=memory&cache=shared&_fk=1",
		DynamoIsLocal: true,
	}
	entClient := client.NewEntClient(cfg)
	dynamoClient, err := client.NewDynamoClient(cfg)
	if err != nil {
		t.Fatal(err)
	}

	m := migration.NewMigration(dynamoClient)
	if err := m.Migrate(true); err != nil {
		t.Fatal(err)
	}

	authDynamoRepo := repository.NewAuthenticationDynamoRepository(dynamoClient)
	userDynamoRepo := repository.NewUserDynamoRepository(dynamoClient)

	return service.NewAuthenticationService(entClient, cfg, authDynamoRepo, userDynamoRepo)
}

func TestAuthenticationService_SignUp(t *testing.T) {
	srv := initAuthenticationService(t)

	t.Run("should create a new user", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		authentication, user, err := srv.SignUp(
			ctx,
			"google",
			uuid.New().String(),
			uuid.New().String(),
			"userUser",
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
	//srv := initAuthenticationService(t)

	t.Run("should create a new user", func(t *testing.T) {
		t.Parallel()

		//ctx := context.Background()
		//_, err := srv.SignInOrSignUp(
		//	ctx,
		//	"google",
		//	"",
		//	"",
		//	"test",
		//)
		//if err != nil {
		//	t.Fatal(err)
		//}

	})
}
