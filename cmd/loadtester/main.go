package main

import (
	"github.com/ivofulco/goexpert-stress-test/internal/infrastructure"
)

func main() {
	container := infrastructure.NewContainer()

	cli := container.GetCLI()

	cli.Execute()
}
