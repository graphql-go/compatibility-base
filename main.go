package main

import (
	"log"

	"github.com/graphql-go/compatibility-base/cmd"
)

func main() {
	params := cmd.NewParams{}
	cli := cmd.New(&params)

	log.Println(cli)
}
