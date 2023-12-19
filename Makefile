
# make ent-new name=User
ent-new:
	@go run -mod=mod entgo.io/ent/cmd/ent new $(name) --target ./internal/ent/schema

ent-gen:
	@go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/ent/schema
