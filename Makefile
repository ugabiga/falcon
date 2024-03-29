
# make ent-new name=User
ent-new:
	@go run -mod=mod entgo.io/ent/cmd/ent new $(name) --target ./internal/ent/schema

ent-gen:
	@go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema
	@go generate ./internal/ent/

ent-visual:
	@atlas schema inspect -u ent://internal/ent/schema --dev-url "sqlite://demo?mode=memory&_fk=1" --visualize

dev:
	@air server | docker-compose up -d

db:
	@air server

df:
	@cd web && yarn dev & cd web && yarn openapi

and-reverse:
	@/Users/sanghwa/Library/Android/sdk/platform-tools/adb reverse tcp:3000 tcp:3000
	@/Users/sanghwa/Library/Android/sdk/platform-tools/adb reverse tcp:8080 tcp:8080

gql-gen:
	@go run github.com/99designs/gqlgen generate
	@cd web && yarn codegen

build:
	@docker build -t "falcon:latest" -f ./lambda.Dockerfile .

deploy:
	@git checkout prod && git merge main && git push && git checkout main
