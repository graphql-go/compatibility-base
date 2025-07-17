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
			},
			BaseStyle: bubbletea.NewBaseStyle(),
		}),
	}
	cli := cmd.New(&cmdParams)

	runParams := &cmd.RunParams{
		ResultCallback: nil,
	}

	if _, err := cli.Run(runParams); err != nil {
		handleErr(err)
	}
}

