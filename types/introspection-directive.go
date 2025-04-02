package types

type IntrospectionDirective struct {
	Name         string                    `json:"name"`
	Description  string                    `json:"description"`
	IsRepeatable bool                      `json:"isRepeatable"`
	Locations    []DirectiveLocation       `json:"locations"`
	Args         []IntrospectionInputValue `json:"args"`
}

type DirectiveLocation string

const (
	Query                DirectiveLocation = "QUERY"
	Mutation             DirectiveLocation = "MUTATION"
	Subscription         DirectiveLocation = "SUBSCRIPTION"
	Field                DirectiveLocation = "FIELD"
	FragmentDefinition   DirectiveLocation = "FRAGMENT_DEFINITION"
	FragmentSpread       DirectiveLocation = "FRAGMENT_SPREAD"
	InlineFragment       DirectiveLocation = "INLINE_FRAGMENT"
	VariableDefinition   DirectiveLocation = "VARIABLE_DEFINITION"
	Schema               DirectiveLocation = "SCHEMA"
	Scalar               DirectiveLocation = "SCALAR"
	Object               DirectiveLocation = "OBJECT"
	FieldDefinition      DirectiveLocation = "FIELD_DEFINITION"
	ArgumentDefinition   DirectiveLocation = "ARGUMENT_DEFINITION"
	Interface            DirectiveLocation = "INTERFACE"
	Union                DirectiveLocation = "UNION"
	Enum                 DirectiveLocation = "ENUM"
	EnumValue            DirectiveLocation = "ENUM_VALUE"
	InputObject          DirectiveLocation = "INPUT_OBJECT"
	InputFieldDefinition DirectiveLocation = "INPUT_FIELD_DEFINITION"
)

type IntrospectionInputValue struct {
	Name              string                    `json:"name"`
	Description       string                    `json:"description"`
	TypeRef           IntrospectionInputTypeRef `json:"typeRef"`
	DefaultValue      string                    `json:"defaultValue"`
	IsDeprecated      bool                      `json:"isDeprecated"`
	DeprecationReason string                    `json:"deprecationReason"`
}

type IntrospectionInputType interface {
}

type IntrospectionInputTypeRef interface {
}

type IntrospectionNamedTypeRef interface {
	IntrospectionInputType
}

type IntrospectionListTypeRef interface {
	IntrospectionInputTypeRef
}

type IntrospectionNonNullTypeRef interface {
	IntrospectionNamedTypeRef
	IntrospectionListTypeRef
}
