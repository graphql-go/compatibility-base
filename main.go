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

	log.Println(result)
}
