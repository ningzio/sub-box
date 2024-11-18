package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ningzio/sub-box/backend/internal/config"
	"github.com/ningzio/sub-box/backend/pkg/schema"

	"github.com/sagernet/serenity/option"
	box "github.com/sagernet/sing-box/option"
)

func main() {
	_ = box.LogOptions{}
	_ = option.Options{}

	// // Create generator with registry
	generator := schema.NewGenerator(schema.SchemaOptions{
		RequireAll:       false,
		IgnoreZeroValues: true,
		IncludeExamples:  true,
		Registry:         config.NewRegistry(),
	})

	// Generate schema for Options struct
	schema, err := generator.Generate(option.Options{})
	if err != nil {
		log.Fatalf("Failed to generate schema: %v", err)
	}

	// _ = schema

	// Output schema as JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(schema); err != nil {
		log.Fatalf("Failed to encode schema: %v", err)
	}
}
