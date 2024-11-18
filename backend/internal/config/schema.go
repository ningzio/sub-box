package config

import (
	"github.com/ningzio/sub-box/backend/pkg/schema"

	"github.com/sagernet/serenity/option"
	box "github.com/sagernet/sing-box/option"
)

// for testing
type Foo struct {
	Foo string
}

type Bar struct {
	Bar string
}

// NewRegistry creates a new schema registry with predefined type information
func NewRegistry() *schema.Patcher {
	_ = box.LogOptions{}
	_ = option.Options{}
	registry := schema.NewRegistry()

	registry.RegisterStruct(box.LogOptions{}, schema.JSONSchemaPatch{
		Description: "Log options",
		Type:        "object",
		Properties: map[string]schema.JSONSchemaPatch{
			"level": {
				Description: "Log level",
				Type:        "string",
				Enums: []interface{}{
					"debug",
					"info",
					"warn",
					"error",
				},
			},
		},
	}).AddProperty("newFields", schema.JSONSchemaPatch{
		Type:  "object",
		OneOf: []any{Foo{}, Bar{}},
	})

	return registry
}

func ptr[T any](v T) *T {
	return &v
}
