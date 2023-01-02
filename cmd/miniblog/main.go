package main

import (
	"os"

	"github.com/cjhw/miniblog/internal/miniblog"
)

func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
