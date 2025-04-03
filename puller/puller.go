package puller

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/graphql-go/compatibility-base/types"
)

// reposDirName is the code repository root directory name.
const reposDirName = "repos"

// Puller represents the puller component.
type Puller struct {
}

// New returns a pointer to a Puller struct.
func New() *Puller {
	return &Puller{}
}

// PullParams represents the parameters of the pull method.
type PullParams struct {
	// Specification is the code repository of the graphql specification.
	Specification *types.Repository

	// Implementation is the code repository of the graphql implementation.
	Implementation *types.Repository
}

// PullResult represents the result of the pull method.
type PullResult struct {
}

// Pull pulls a set of code repositories and returns the result.
func (p *Puller) Pull(params *PullParams) (*PullResult, error) {
	repos := []*types.Repository{}

	if params.Specification != nil {
		repos = append(repos, params.Specification)
	}

	if params.Implementation != nil {
		repos = append(repos, params.Implementation)
	}

	if _, err := os.Stat(reposDirName); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(reposDirName, os.ModePerm); err != nil {
				return nil, fmt.Errorf("failed to create a directory: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to check if directory exist: %w", err)
		}
	}

	for _, r := range repos {
		name := filepath.Join(reposDirName, r.Name)

		if _, err := os.Stat(name); os.IsNotExist(err) {
			if err := os.Mkdir(name, os.ModePerm); err != nil {
				return nil, fmt.Errorf("failed to create a directory: %w", err)
			}
		}

		if _, err := git.PlainClone(name, false, &git.CloneOptions{
			URL:      r.URL,
			Progress: os.Stdout,
		}); err != nil {
			if strings.Contains(err.Error(), "repository already exists") {
				return nil, nil
			}

			return nil, fmt.Errorf("failed to clone a git repository: %w", err)
		}
	}

	return &PullResult{}, nil
}
