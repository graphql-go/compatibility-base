package cmd

import (
	"fmt"

	"graphql-go/compatibility-base/bubbletea"
)

// CLI represents the command line interface component.
type CLI struct {
}

// RunResult is the result of executing the run method.
type RunResult struct {
	Choice string
}

// RunParams are the parameters of the run method.
type RunParams struct {
	Choices []string
	Header  string
}

// Run runs the CLI application and returns the result.
func (c *CLI) Run(p *RunParams) (*RunResult, error) {
	bt := bubbletea.New(&bubbletea.Params{
		Choices: p.Choices,
		UI: bubbletea.UIParams{
			Header: p.Header,
		},
	})

	btRunResult, err := bt.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	return &RunResult{
		Choice: btRunResult.Choice,
	}, nil
}
