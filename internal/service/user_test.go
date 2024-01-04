package service_test

import (
	"github.com/ugabiga/falcon/internal/ent"
	"testing"
)

func prepareUser(t *testing.T) *ent.User {
	//authenticationSrv := initAuthenticationService(t)
	//id := uuid.New().String()

	//a, err := authenticationSrv.SignUp(
	//	context.Background(),
	//	"google",
	//	id,
	//	"",
	//	"test",
	//)
	//if err != nil {
	//	t.Fatal(err)
	//}

	return &ent.User{}
}

func TestUserService_GetUser(t *testing.T) {
	//ctx := context.Background()
	//targetUser := prepareUser(t)

	//tester := app.InitializeTestApplication()
	//srv := tester.UserSrv

	t.Run("should return a user", func(t *testing.T) {
		t.Parallel()

		//u, err := srv.GetByID(
		//	ctx,
		//	targetUser.ID,
		//)
		//if err != nil {
		//	t.Fatal(err)
		//}
		//
		//if u == nil {
		//	t.Fatal("u is nil")
		//}
		//
		//if u.ID == 0 {
		//	t.Fatal("u.ID is 0")
		//}
	})
}

func TestUserService_EditUser(t *testing.T) {
	//ctx := context.Background()
	//targetUser := prepareUser(t)
	//
	//srv := initUserService()
	//
	//t.Run("should update a user", func(t *testing.T) {
	//	t.Parallel()
	//
	//	u, err := srv.Update(
	//		ctx,
	//		targetUser.ID,
	//		&ent.User{
	//			Name:     "new name",
	//			Timezone: "new timezone",
	//		},
	//	)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	if u == nil {
	//		t.Fatal("u is nil")
	//	}
	//
	//	if u.ID == 0 {
	//		t.Fatal("u.ID is 0")
	//	}
	//
	//	if u.Name != "new name" {
	//		t.Fatal("u.Name is not new name")
	//	}
	//
	//	if u.Timezone != "new timezone" {
	//		t.Fatal("u.Timezone is not new timezone")
	//	}
	//})
}
