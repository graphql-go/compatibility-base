package types

type IntrospectionType interface {
}

type IntrospectionScalarType struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	SpecifiedByURL string `json:"specifiedByURL"`
}

// TypeRef represents a reference to a type in the schema
type TypeRef struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type IntrospectionObjectType struct {
	Kind        string                `json:"kind"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Fields      []IntrospectionField  `json:"fields"`
	Interfaces  []TypeRef             `json:"interfaces"`
}

type IntrospectionInterfaceType struct {
	Kind          string                `json:"kind"`
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	Fields        []IntrospectionField  `json:"fields"`
	Interfaces    []TypeRef             `json:"interfaces"`
	PossibleTypes []TypeRef             `json:"possibleTypes"`
}

type IntrospectionUnionType struct {
	Kind          string    `json:"kind"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	PossibleTypes []TypeRef `json:"possibleTypes"`
}

type IntrospectionEnumValue struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

type IntrospectionEnumType struct {
	Kind        string                   `json:"kind"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	EnumValues  []IntrospectionEnumValue `json:"enumValues"`
}

type IntrospectionInputObjectType struct {
	Kind        string                    `json:"kind"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	InputFields []IntrospectionInputValue `json:"inputFields"`
	IsOneOf     bool                      `json:"isOneOf"`
}
