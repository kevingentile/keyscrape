package main

import (
	"fmt"
	"os"

	"github.com/kevingentile/keyscrape/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
