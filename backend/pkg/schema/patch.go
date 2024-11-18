package schema

import (
	"reflect"
)

// JSONSchemaPatch contains metadata about a type
type JSONSchemaPatch struct {
	// Enums contains enum values for the type
	Enums []interface{} `json:"enum,omitempty"`
	// Description of the type
	Description string `json:"description,omitempty"`
	// Properties for object types
	Properties map[string]JSONSchemaPatch `json:"properties,omitempty"`
	// Type of the value
	Type string `json:"type,omitempty"`

	OneOf []any `json:"oneOf,omitempty"`
	AnyOf []any `json:"anyOf,omitempty"`
}

func (t *JSONSchemaPatch) AddProperty(name string, info JSONSchemaPatch) *JSONSchemaPatch {
	if t.Properties == nil {
		t.Properties = make(map[string]JSONSchemaPatch)
	}
	t.Properties[name] = info
	return t
}

func (t *JSONSchemaPatch) AddProperties(properties map[string]JSONSchemaPatch) *JSONSchemaPatch {
	if t.Properties == nil {
		t.Properties = make(map[string]JSONSchemaPatch)
	}
	for k, v := range properties {
		t.Properties[k] = v
	}
	return t
}

func (t *JSONSchemaPatch) AddOneOf(v ...any) *JSONSchemaPatch {
	t.OneOf = append(t.OneOf, v...)
	return t
}

func (t *JSONSchemaPatch) AddAnyOf(v ...any) *JSONSchemaPatch {
	t.AnyOf = append(t.AnyOf, v...)
	return t
}

func (t *JSONSchemaPatch) GetProperty(name string) (JSONSchemaPatch, bool) {
	if t.Properties == nil {
		return JSONSchemaPatch{}, false
	}
	p, ok := t.Properties[name]
	return p, ok
}

// Patcher manages type metadata
type Patcher struct {
	types map[string]JSONSchemaPatch
}

// getTypeKey1 returns a unique key for a type and its field path
func getTypeKey1(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var key string
	if t.PkgPath() == "" {
		key = t.Name()
	} else {
		key = t.PkgPath() + "." + t.Name()
	}
	return key
}

// NewRegistry creates a new type registry
func NewRegistry() *Patcher {
	return &Patcher{
		types: make(map[string]JSONSchemaPatch),
	}
}

// RegisterStruct registers a new struct type with its type information
func (r *Patcher) RegisterStruct(t interface{}, info JSONSchemaPatch, properties ...map[string]JSONSchemaPatch) *JSONSchemaPatch {
	typ := reflect.TypeOf(t)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	key := getTypeKey1(typ)
	for _, p := range properties {
		info.AddProperties(p)
	}
	r.types[key] = info
	return &info
}

// GetStructInfo retrieves struct type information
func (r *Patcher) GetStructInfo(t interface{}) (JSONSchemaPatch, bool) {
	return r.GetStructInfoFromType(reflect.TypeOf(t))
}

// GetStructInfoFromType retrieves struct type information using reflect.Type
func (r *Patcher) GetStructInfoFromType(typ reflect.Type) (JSONSchemaPatch, bool) {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	key := getTypeKey1(typ)
	si, ok := r.types[key]
	return si, ok
}
