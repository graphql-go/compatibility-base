package types

type IntrospectionType interface {
}

type IntrospectionScalarType struct {
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	SpecifiedByURL string `json:"specifiedByURL"`
}

type IntrospectionObjectType struct {
	Kind        string             `json:"kind"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Fields      IntrospectionField `json:"fields"`
	Interfaces  interface{}        `json:"interfaces"`
}

type IntrospectionInterfaceType struct {
	Kind        string             `json:"kind"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Fields      IntrospectionField `json:"fields"`
	Interfaces  interface{}        `json:"interfaces"`
	// possibleTypes IntrospectionObjectType    `json:"possibleTypes"`
}

type IntrospectionUnionType struct {
	Kind          string                  `json:"kind"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	PossibleTypes IntrospectionObjectType `json:"possibleTypes"`
}

type IntrospectionEnumType struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// enumValues  IntrospectionEnumValue `json:"enumValues"`
}

type IntrospectionInputObjectType struct {
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// inputFields IntrospectionInputValue `json:"inputFields"`
	IsOneOf bool `json:"isOneOf"`
}
