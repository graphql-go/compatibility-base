package cmd

import (
	"errors"
	"fmt"

	"github.com/graphql-go/compatibility-base/bubbletea"
)

// CLI represents the command line interface component.
type CLI struct {
	bubbletea *bubbletea.BubbleTea
}

// NewParams represents the parameters for the new method.
type NewParams struct {
	Bubbletea *bubbletea.BubbleTea
}

// New returns a pointer to the `CLI` struct.
func New(p *NewParams) *CLI {
	return &CLI{
		bubbletea: p.Bubbletea,
	}
}

// RunResult is the result of executing the run method.
type RunResult struct {
	// Choice is the option chosen after running the CLI application.
	Choice string
}

// RunParams are the parameters of the run method.
type RunParams struct {
}

// Run runs the CLI application and returns the result.
func (c *CLI) Run(p *RunParams) (*RunResult, error) {
	btRunResult, err := c.bubbletea.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	r, ok := btRunResult.(*bubbletea.ChoicesModelResult)
	if !ok {
		return nil, errors.New("unexpected type")
	}

	return &RunResult{
		Choice: r.Choice,
	}, nil
}
