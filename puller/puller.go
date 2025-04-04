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

// repositories returns the parameters as a repositories slice.
func (p *PullParams) repositories() []*types.Repository {
	repos := []*types.Repository{}

	if p.Specification != nil {
		repos = append(repos, p.Specification)
	}

	if p.Implementation != nil {
		repos = append(repos, p.Implementation)
	}

	return repos
}

// PullResult represents the result of the pull method.
type PullResult struct {
}

// Pull pulls a set of code repositories and returns the result.
func (p *Puller) Pull(params *PullParams) (*PullResult, error) {
	repos := params.repositories()

	if err := p.createReposDir(); err != nil {
		return nil, err
	}

	if err := p.gitCloneRepos(repos); err != nil {
		return nil, err
	}

	return &PullResult{}, nil
}

// createReposDir creates the `repos` directory and returns whether it succeeded or not.
func (p *Puller) createReposDir() error {
	if _, err := os.Stat(reposDirName); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(reposDirName, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create a directory: %w", err)
			}
		} else {
			return fmt.Errorf("failed to check if directory exist: %w", err)
		}
	}

	return nil
}

// gitCloneRepos clones the given repositories and returns whether or not it succeeded.
func (p *Puller) gitCloneRepos(repos []*types.Repository) error {
	for _, r := range repos {
		name := filepath.Join(reposDirName, r.Name)

		if err := p.createRepoDir(r.Name); err != nil {
			return err
		}

		if _, err := git.PlainClone(name, false, &git.CloneOptions{
			URL:      r.URL,
			Progress: os.Stdout,
		}); err != nil {
			if strings.Contains(err.Error(), "repository already exists") {
				return nil
			}

			return fmt.Errorf("failed to clone a git repository: %w", err)
		}
	}

	return nil
}

// createRepoDir creates the `repo` directory and returns whether it succeeded or not.
func (p *Puller) createRepoDir(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create a directory: %w", err)
		}
	}

	return nil
}
