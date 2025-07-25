package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/compatibility-base/bubbletea"
	"github.com/graphql-go/compatibility-base/cmd"
	"github.com/graphql-go/compatibility-base/config"
	"github.com/graphql-go/compatibility-base/implementation"
)

func main() {
	handleErr := func(err error) {
		log.Fatal(err)
	}

	cfg := config.New()
	defaultSpecTableHeader := fmt.Sprintf("Ref: %s", cfg.GraphqlJSImplementation.Repo.URL)
	defaultImplTableHeader := "Impl: https://github.com/graphql-go/graphql"
	choicesModelUIHeader := cfg.GraphqlJSImplementation.Repo.String(implementation.RefImplementationPrefix)

	cmdParams := cmd.NewParams{
		Bubbletea: bubbletea.New(&bubbletea.Params{
			Models: bubbletea.Models{
				bubbletea.NewChoicesModel(&bubbletea.ChoicesModelParams{
					Order:   1,
					Choices: cfg.AvailableImplementations,
					UI: bubbletea.ChoicesModelUIParams{
						Header: choicesModelUIHeader,
					},
				}),
				newTableModel(defaultSpecTableHeader, defaultImplTableHeader),
			},
			BaseStyle: bubbletea.NewBaseStyle(),
		}),
	}
	cli := cmd.New(&cmdParams)

	resultCallback := func(result *bubbletea.BubbleTeaResult) error {
		implementationHeader := result.ChoicesModelResult.Choice
		tableModel := newTableModel(defaultSpecTableHeader, implementationHeader)

		if err := cli.UpdateModel(tableModel); err != nil {
			log.Printf("failed to update table model: %v", err)
			return err
		}

		return nil
	}

	runParams := &cmd.RunParams{
		ResultCallback: resultCallback,
	}

	if _, err := cli.Run(runParams); err != nil {
		handleErr(err)
	}
}

// `newTableModel` creates and returns a pointer to `bubbletea.TableModel`.
func newTableModel(specificationHeader string, implementationHeader string) *bubbletea.TableModel {
	headerWidth := uint(16)

	return bubbletea.NewTableModel(&bubbletea.TableModelParams{
		Order: 2,
		Headers: []bubbletea.TableHeader{
			{Title: "Metric", Width: 35},
			{Title: specificationHeader, Width: headerWidth},
			{Title: implementationHeader, Width: headerWidth},
			{Title: "Diff Ratio", Width: headerWidth},
			{Title: "Max Diff", Width: headerWidth},
			{Title: "Result", Width: headerWidth},
		},
		Rows: [][]string{
			{"GitHub:", "", "", "", "", ""},
			{"License", "MIT", "MIT", "0%", "0%", "✅"},
			{"Number Of Stars", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Issues Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Issues Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Pull Requests Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Pull Requests Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Forks", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Last Commit Date", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Contributors", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"GraphQL Compatibility Keywords:", "", "", "", "", ""},
			{"Number Of Comments Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"Number Of Comments Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
			{"GraphQL:", "", "", "", "", ""},
			{"Specification Version", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
		},
	})
}
