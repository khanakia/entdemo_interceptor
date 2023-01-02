package main

import (
	"log"

	_ "github.com/lib/pq"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {

	err := entc.Generate("./ent/schema",
		&gen.Config{},
		// entc.TemplateDir("./tmpl"),
		entc.FeatureNames("sql/execquery", "intercept", "schema/snapshot"),
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
