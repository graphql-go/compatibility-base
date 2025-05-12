package main

import (
	"log"

	"github.com/graphql-go/compatibility-base/bubbletea"
	"github.com/graphql-go/compatibility-base/cmd"
)

func main() {
	handleErr := func(err error) {
		log.Fatal(err)
	}

	params := cmd.NewParams{
		Bubbletea: bubbletea.New(&bubbletea.Params{}),
	}
	cli := cmd.New(&params)

	result, err := cli.Run(&cmd.RunParams{})
	if err != nil {
		handleErr(err)
	}

	log.Println(result)
}
