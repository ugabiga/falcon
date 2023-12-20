//go:build ignore
// +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"log"
)

func main() {
	log.Println("Running ent code generation...")

	extensions := entgqlExtensions()

	if err := entc.Generate("./schema",
		&gen.Config{
			IDType: &field.TypeInfo{Type: field.TypeString},
		},
		entc.Extensions(extensions...),
	); err != nil {
		log.Fatalf("ruuing ent code generator: %v", err)
	}

	log.Println("Finished ent code generation")
}

func entgqlExtensions() []entc.Extension {
	configDir := "../.."
	graphqlDir := "../../api/graph"

	ex, err := entgql.NewExtension(
		entgql.WithConfigPath(configDir+"/gqlgen.yml"),
		entgql.WithSchemaPath(graphqlDir+"/ent.graphql.bak"),
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	return []entc.Extension{
		ex,
	}
}
