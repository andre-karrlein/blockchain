package main

import (
	"os"

	"github.com/andre-karrlein/blockchain/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}

	cli.Run()

}
