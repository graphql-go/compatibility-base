package types

type IntrospectionQueryResult struct {
	Schema IntrospectionSchema `json:"__schema"`
}
