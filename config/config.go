package config

import (
	"os"

	"github.com/graphql-go/compatibility-base/implementation"
	"github.com/graphql-go/compatibility-base/types"
)

// Config represents a set of configuration values.
type Config struct {
	// IsDebug represents whether or not the configuration is in debug mode.
	IsDebug bool

	// GraphqlGoImplementation represents the graphql-go implementation.
	GraphqlGoImplementation types.Implementation

	// GraphqlJSImplementation represents the graphql-js implementation.
	GraphqlJSImplementation types.Implementation

	// GraphqlSpecification represents the graphql specification.
	GraphqlSpecification types.Specification

	// RefImplementation is the default reference implementation.
	RefImplementation types.Implementation

	// Implementations is the default list of graphql implementations.
	Implementations []types.Implementation

	// GraphqlSpecificationWithPrefix represents the graphql specification repository link with a prefix.
	GraphqlSpecificationWithPrefix string

	// AvailableImplementations represents a list of available implementations.
	AvailableImplementations []string
}

// New returns a pointer to a Config struct.
func New() *Config {
	isDebug := isDebug()
	graphqlGoImplementation := graphqlGoImplementation()
	graphqlJSImplementation := graphqlJSImplementation()
	graphqlSpecification := graphqlSpecification()
	implementations := []types.Implementation{graphqlGoImplementation}
	graphqlSpecificationWithPrefix := graphqlSpecificationWithPrefix(graphqlSpecification)
	availableImplementations := availableImplementations(implementations)

	return &Config{
		IsDebug:                        isDebug,
		GraphqlGoImplementation:        graphqlGoImplementation,
		GraphqlJSImplementation:        graphqlJSImplementation,
		GraphqlSpecification:           graphqlSpecification,
		RefImplementation:              graphqlJSImplementation,
		Implementations:                implementations,
		GraphqlSpecificationWithPrefix: graphqlSpecificationWithPrefix,
		AvailableImplementations:       availableImplementations,
	}
}

// availableImplementations returns a list of available implementations.
func availableImplementations(implementations []types.Implementation) []string {
	var result = []string{}

	for _, i := range implementations {
		result = append(result, i.Repo.String(implementation.ImplementationPrefix))
	}

	return result
}

// graphqlGoImplementation returns the graphql-go implementation.
func graphqlGoImplementation() types.Implementation {
	return types.Implementation{
		Repo: types.Repository{
			Name:          "graphql-go-graphql",
			URL:           "https://github.com/graphql-go/graphql",
			ReferenceName: "v0.8.1",
			Dir:           "./repos/graphql-go-graphql/",
		},
		Type: types.GoImplementationType,
	}
}

// graphqlJSImplementation returns the graphql-js implementation.
func graphqlJSImplementation() types.Implementation {
	return types.Implementation{
		Repo: types.Repository{
			Name:          "graphql-graphql-js",
			URL:           "https://github.com/graphql/graphql-js",
			ReferenceName: "v0.6.0",
			Dir:           "./repos/graphql-graphql-js/",
		},
		Type:              types.RefImplementationType,
		TestNamesFilePath: "./puller-js/unit-tests.txt",
	}
}

// graphqlSpecification returns the graphql specification.
func graphqlSpecification() types.Specification {
	return types.Specification{
		Repo: types.Repository{
			Name:          "graphql-specification",
			URL:           "https://github.com/graphql/graphql-spec",
			ReferenceName: "October2021",
			Dir:           "./repos/graphql-specification/",
		},
	}
}

// isDebug returns the current debug value.
func isDebug() bool {
	return os.Getenv("DEBUG") == "true"
}

// graphqlSpecificationWithPrefix returns the graphql specification repository link with a prefix.
func graphqlSpecificationWithPrefix(graphqlSpecification types.Specification) string {
	return graphqlSpecification.Repo.String(implementation.SpecificationPrefix)
}
