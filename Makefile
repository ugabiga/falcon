
# make ent-new name=User
ent-new:
	@go run -mod=mod entgo.io/ent/cmd/ent new $(name) --target ./internal/ent/schema

ent-gen:
	@go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema
	@go generate ./internal/ent/

ent-visual:
	@atlas schema inspect -u ent://internal/ent/schema --dev-url "sqlite://demo?mode=memory&_fk=1" --visualize

dev:
	@docker-compose up -d --build && \
		air

gql-gen:
	@go run github.com/99designs/gqlgen generate
	@cd web && yarn codegen

build:
	@docker build -t "falcon:latest" -f ./lambda.Dockerfile .
