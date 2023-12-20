package resolvers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"github.com/ugabiga/falcon/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userSrv *service.UserService
}

func NewResolver(
	userSrv *service.UserService,
) *handler.Server {
	resolver := &Resolver{
		userSrv: userSrv,
	}

	graphSrv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers:  resolver,
				Directives: generated.DirectiveRoot{},
				Complexity: generated.ComplexityRoot{},
			},
		),
	)
	return graphSrv
}
