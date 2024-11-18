package schema

// JSONSchema represents a JSON Schema
type JSONSchema struct {
	// Core schema metadata
	ID          string                 `json:"$id,omitempty"`
	Schema      string                 `json:"$schema,omitempty"`
	Ref         string                 `json:"$ref,omitempty"`
	Definitions map[string]*JSONSchema `json:"definitions,omitempty"`

	// Schema annotations
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Default     any    `json:"default,omitempty"`
	Examples    []any  `json:"examples,omitempty"`

	// Basic validation
	Type string        `json:"type,omitempty"`
	Enum []interface{} `json:"enum,omitempty"`

	// Composition
	OneOf []*JSONSchema `json:"oneOf,omitempty"`
	AnyOf []*JSONSchema `json:"anyOf,omitempty"`

	// Array and Object validation
	Items                *JSONSchema            `json:"items,omitempty"`
	AdditionalProperties *JSONSchema            `json:"additionalProperties,omitempty"`
	Properties           map[string]*JSONSchema `json:"properties,omitempty"`
	Required             []string               `json:"required,omitempty"`
}

// SchemaOptions contains options for schema generation
type SchemaOptions struct {
	// RequireAll if true, marks all fields as required unless explicitly tagged otherwise
	RequireAll bool
	// IgnoreZeroValues if true, omits zero values from examples
	IgnoreZeroValues bool
	// IncludeExamples if true, includes example values in the schema
	IncludeExamples bool
	// Registry used to retrieve type information
	Registry *Patcher
}
