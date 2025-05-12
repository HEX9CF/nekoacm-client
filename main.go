package main

import (
	"nekoacm-client/cmd/bootstrap"
	"nekoacm-client/cmd/cli"
)

func main() {
	err := bootstrap.Init()
	if err != nil {
		panic(err)
	}
	cli.Execute()
}
