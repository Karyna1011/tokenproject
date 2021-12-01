package main

import (
	"os"

  "gitlab.com/tokend/subgroup/tokenproject/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
