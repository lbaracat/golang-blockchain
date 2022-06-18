package main

import (
	"os"

	"github.com/lbaracat/golang-blockchain/internal/cli"
)

func main() {
	defer os.Exit(0)

	cli := cli.CommandLine{}
	cli.Run()
}
