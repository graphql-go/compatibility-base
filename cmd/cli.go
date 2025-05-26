package cmd

import (
	"fmt"

	"github.com/graphql-go/compatibility-base/bubbletea"
)

// CLI represents the command line interface component.
type CLI struct {
	// bubbletea is the component that wraps the bubbletea library.
	bubbletea *bubbletea.BubbleTea
}

// NewParams represents the parameters for the new method.
type NewParams struct {
	// Bubbletea is the component parameter that wraps the bubbletea library.
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
	// `ChoicesModelResult` is the result of the `ChoicesModel`.
	ChoicesModelResult *bubbletea.ChoicesModelResult

	// `TableModelResult` is the result of the `TableModel`.
	TableModelResult *bubbletea.TableModelResult
}

// RunParams are the parameters of the run method.
type RunParams struct {
	// `ResultCallback` is the results callback Run parameter.
	ResultCallback func(result *bubbletea.BubbleTeaResult) error
}

// Run runs the CLI application and returns the result.
func (c *CLI) Run(p *RunParams) (*RunResult, error) {
	runParams := bubbletea.RunParams{
		ResultCallback: p.ResultCallback,
	}

	btRunResult, err := c.bubbletea.Run(runParams)
	if err != nil {
		return nil, fmt.Errorf("failed to run: %w", err)
	}

	r := &RunResult{
		ChoicesModelResult: btRunResult.ChoicesModelResult,
		TableModelResult:   btRunResult.TableModelResult,
	}

	return r, nil
}

// `UpdateModel` calls the `BubbleTea` component `UpdateModel` method.
func (c *CLI) UpdateModel(model bubbletea.Model) error {
	return c.bubbletea.UpdateModel(model)
}
