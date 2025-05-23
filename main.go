package main

import (
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
	header := cfg.GraphqlJSImplementation.Repo.String(implementation.RefImplementationPrefix)
	headerWidth := uint(15)

	params := cmd.NewParams{
		Bubbletea: bubbletea.New(&bubbletea.Params{
			Models: bubbletea.Models{
				bubbletea.NewChoicesModel(&bubbletea.ChoicesModelParams{
					Order:   1,
					Choices: cfg.AvailableImplementations,
					UI: bubbletea.ChoicesModelUIParams{
						Header: header,
					},
				}),
				bubbletea.NewTableModel(&bubbletea.TableModelParams{
					Order: 2,
					Headers: []bubbletea.TableHeader{
						{Title: "Metric", Width: 35},
						{Title: "Spec: https://github.com/graphql/graphql-js", Width: headerWidth},
						{Title: "Impl: https://github.com/graphql-go/graphql", Width: headerWidth},
						{Title: "Diff Ratio", Width: headerWidth},
						{Title: "Max Diff", Width: headerWidth},
						{Title: "Result", Width: headerWidth},
					},
					Rows: [][]string{
						[]string{"GitHub:", "", "", "", "", ""},
						[]string{"License", "MIT", "MIT", "0%", "0%", "✅"},
						[]string{"Number Of Stars", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Issues Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Issues Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Pull Requests Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Pull Requests Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Forks", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Last Commit Date", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Contributors", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"GraphQL Compatibility Keywords:", "", "", "", "", ""},
						[]string{"Number Of Comments Open", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"Number Of Comments Closed", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
						[]string{"GraphQL:", "", "", "", "", ""},
						[]string{"Specification Version", "Loading...", "Loading...", "Loading...", "Loading...", "Loading..."},
					},
				}),
			},
			BaseStyle: bubbletea.NewBaseStyle(),
		}),
	}
	cli := cmd.New(&params)

	result, err := cli.Run(&cmd.RunParams{})
	if err != nil {
		handleErr(err)
	}

	log.Printf("ChoicesModelResult: %+v", result.ChoicesModelResult)
	log.Printf("TableModelResult: %+v", result.TableModelResult)
}
