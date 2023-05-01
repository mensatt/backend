//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{
		Package: "github.com/mensatt/backend/internal/database/ent",
		Target:  "./ent",
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	}); err != nil {
		log.Fatal("running database codegen:", err)
	}
}
