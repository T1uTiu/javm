package main

import (
	"fmt"
	"os"

	"github.com/t1utiu/javm/cmd"
)

func main() {
	cmd.Init()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
