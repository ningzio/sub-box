package schema

import (
	"fmt"
	"reflect"
	"strings"
)

// Generator handles JSON Schema generation from Go types
type Generator struct {
	options     SchemaOptions
	processing  map[reflect.Type]bool
	definitions map[string]*JSONSchema
}

// NewGenerator creates a new schema generator with the given options
func NewGenerator(options SchemaOptions) *Generator {
	if options.Registry == nil {
		options.Registry = NewRegistry()
	}
	return &Generator{
		options:     options,
		processing:  make(map[reflect.Type]bool),
		definitions: make(map[string]*JSONSchema),
	}
}

// Generate creates a JSON Schema from a Go type
func (g *Generator) Generate(v interface{}) (*JSONSchema, error) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Create the root schema
	rootSchema := &JSONSchema{
		Schema: "http://json-schema.org/draft-07/schema#",
		Type:   "object",
	}

	// Generate schema for the type
	schema, err := g.generateSchema(t)
	if err != nil {
		return nil, err
	}

	// Add the main type to definitions
	if len(g.definitions) > 0 {
		rootSchema.Definitions = g.definitions
	}

	// Copy all properties from the generated schema to the root
	rootSchema.Properties = schema.Properties
	rootSchema.Required = schema.Required
	rootSchema.Description = schema.Description
	rootSchema.Title = schema.Title
	rootSchema.Default = schema.Default
	rootSchema.Examples = schema.Examples
	rootSchema.OneOf = schema.OneOf
	rootSchema.AnyOf = schema.AnyOf
	rootSchema.Items = schema.Items
	rootSchema.AdditionalProperties = schema.AdditionalProperties

	return rootSchema, nil
}

// generateSchema generates a schema for the given reflect.Type
func (g *Generator) generateSchema(t reflect.Type) (*JSONSchema, error) {
	// If we're already processing this type, return a reference
	if g.processing[t] {
		return g.createReference(t), nil
	}

	// For struct types, check if we've already generated a schema
	if t.Kind() == reflect.Struct {
		key := getTypeKey(t)
		if _, exists := g.definitions[key]; exists {
			return g.createReference(t), nil
		}
	}

	var schema *JSONSchema
	var err error

	switch t.Kind() {
	case reflect.Ptr:
		return g.generateSchema(t.Elem())
	case reflect.Struct:
		// Mark as processing to detect cycles
		g.processing[t] = true
		schema, err = g.generateStructSchema(t)
		delete(g.processing, t) // Clear processing flag
		if err != nil {
			return nil, err
		}
		// Store in definitions for reuse
		g.definitions[getTypeKey(t)] = schema
		return schema, nil
	case reflect.Slice, reflect.Array:
		return g.generateArraySchema(t)
	case reflect.Map:
		return g.generateMapSchema(t)
	case reflect.String:
		return g.generatePrimitiveSchema("string")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return g.generatePrimitiveSchema("integer")
	case reflect.Float32, reflect.Float64:
		return g.generatePrimitiveSchema("number")
	case reflect.Bool:
		return g.generatePrimitiveSchema("boolean")
	default:
		return nil, fmt.Errorf("unsupported type: %v", t.Kind())
	}
}

// generatePrimitiveSchema generates a schema for primitive types
func (g *Generator) generatePrimitiveSchema(typeName string) (*JSONSchema, error) {
	schema := &JSONSchema{Type: typeName}
	return schema, nil
}

// generateStructSchema generates a schema for a struct type
func (g *Generator) generateStructSchema(t reflect.Type) (*JSONSchema, error) {
	schema := &JSONSchema{
		Type:       "object",
		Properties: make(map[string]*JSONSchema),
	}

	// Get struct type information
	structInfo, structInfoExist := g.options.Registry.GetStructInfoFromType(t)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get the field name from json tag or field name
		name := getFieldName(field)
		if name == "-" {
			fmt.Printf("Skipping field %s with tag %s\n", field.Name, t.PkgPath())
			continue
		}

		// Generate schema for the field
		fieldSchema, err := g.generateSchema(field.Type)
		if err != nil {
			return nil, fmt.Errorf("failed to generate schema for field %s: %v", name, err)
		}

		// Apply field metadata if available
		if structInfoExist {
			if fieldInfo, ok := structInfo.GetProperty(name); ok {
				if err := g.patch(fieldSchema, fieldInfo); err != nil {
					return nil, fmt.Errorf("failed to apply field metadata for field %s: %v", name, err)
				}
			}
		}

		schema.Properties[name] = fieldSchema
	}

	if structInfoExist {
		// Apply struct-level metadata
		if err := g.patch(schema, structInfo); err != nil {
			return nil, fmt.Errorf("failed to apply struct metadata: %v", err)
		}
	}

	return schema, nil
}

// generateArraySchema generates a schema for an array type
func (g *Generator) generateArraySchema(t reflect.Type) (*JSONSchema, error) {
	items, err := g.generateSchema(t.Elem())
	if err != nil {
		return nil, err
	}

	schema := &JSONSchema{
		Type:  "array",
		Items: items,
	}

	return schema, nil
}

// generateMapSchema generates a schema for a map type
func (g *Generator) generateMapSchema(t reflect.Type) (*JSONSchema, error) {
	if t.Key().Kind() != reflect.String {
		return nil, fmt.Errorf("only string keys are supported for maps")
	}

	valueSchema, err := g.generateSchema(t.Elem())
	if err != nil {
		return nil, err
	}

	schema := &JSONSchema{
		Type:                 "object",
		AdditionalProperties: valueSchema,
	}

	return schema, nil
}

// getFieldName gets the field name from struct field, using json tag if available
func getFieldName(field reflect.StructField) string {
	name := field.Name
	if tag := field.Tag.Get("json"); tag != "" {
		parts := strings.Split(tag, ",")
		if parts[0] != "" {
			name = parts[0]
		}
	}
	return name
}

// createReference creates a reference to a type
func (g *Generator) createReference(t reflect.Type) *JSONSchema {
	return &JSONSchema{
		Ref: fmt.Sprintf("#/definitions/%s", getTypeKey(t)),
	}
}

// getTypeKey gets a unique key for a type
func getTypeKey(t reflect.Type) string {
	if t.PkgPath() == "" {
		return t.Name()
	}
	// Remove package path prefix to make the schema cleaner
	return t.Name()
}

// patch applies type information to a schema
func (g *Generator) patch(schema *JSONSchema, patch JSONSchemaPatch) error {
	if patch.Description != "" {
		schema.Description = patch.Description
	}
	if len(patch.Enums) > 0 {
		schema.Enum = patch.Enums
	}
	if patch.Type != "" {
		schema.Type = patch.Type
	}

	for k, p := range patch.Properties {
		if _, ok := schema.Properties[k]; !ok {
			schema.Properties[k] = &JSONSchema{}
		}
		g.patch(schema.Properties[k], p)
	}

	for _, o := range patch.OneOf {
		s, err := g.generateSchema(reflect.TypeOf(o))
		if err != nil {
			return err
		}
		schema.OneOf = append(schema.OneOf, s)
	}

	for _, a := range patch.AnyOf {
		s, err := g.generateSchema(reflect.TypeOf(a))
		if err != nil {
			return err
		}
		schema.AnyOf = append(schema.AnyOf, s)
	}

	return nil
}
