package types

type IntrospectionSchema struct {
	Description      string                   `json:"description"`
	QueryType        IntrospectionObjectType  `json:"queryType"`
	MutationType     IntrospectionObjectType  `json:"mutationType"`
	SubscriptionType IntrospectionObjectType  `json:"subscriptionType"`
	Types            []IntrospectionType      `json:"types"`
	Directives       []IntrospectionDirective `json:"directives"`
}
