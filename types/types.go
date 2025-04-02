// package types defines the internal types.
package types

import "fmt"

// taggedRepoURL is the repo url of a tag of releases.
const taggedRepoURL string = "%s/releases/tag/%s"

// ImplementationType is the type of implementations.
type ImplementationType uint

const (
	// GoImplementationType is the type of a go implementation.
	GoImplementationType = iota + 1

	// RefImplementationType is the type of the graphql reference implementation.
	RefImplementationType
)

// Repository represents the code repository of a graphql implementation.
type Repository struct {
	// Name is the code repository name.
	Name string

	// URL is the code repository URL.
	URL string

	// ReferenceName is the code repository reference name, eg. GitHub a tag.
	ReferenceName string

	// Dir is the code repository directory path.
	Dir string
}

// String returns the string summary of the code repository.
func (r *Repository) String(prefix string) string {
	base := fmt.Sprintf("%s: %s\n", prefix, taggedRepoURL)

	return fmt.Sprintf(base, r.URL, r.ReferenceName)
}

// Introspection represents a graphql introspection.
type Introspection struct {
	// Query is the introspection query.
	Query string
}

// Implementation represents a graphql implementation.
type Implementation struct {
	// Repo is the code repository of the implementation.
	Repo Repository

	// Type is the implementation type.
	Type ImplementationType

	// TestNames is the list of test names of the implementation.
	TestNames []string

	// TestNamesFilePath is the file path of the test names.
	TestNamesFilePath string

	// Introspection is the introspection of the implementation.
	Introspection Introspection
}

// MapKey returns the map key of the implementation.
func (i *Implementation) MapKey(prefix string) string {
	return i.Repo.String(prefix)
}

// Specification represents a graphql specification.
type Specification struct {
	Repo Repository
}

// SpecificationIntrospection represents the introspection result of the graphql specification.
type SpecificationIntrospection struct {
	// QueryResult contains the result of the introspection query.
	QueryResult IntrospectionQueryResult
}

// ImplementationIntrospection represents the introspection result of a graphql implementation.
type ImplementationIntrospection struct {
	// QueryResult contains the result of the introspection query.
	QueryResult IntrospectionQueryResult
}
