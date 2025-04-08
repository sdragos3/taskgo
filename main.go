package main

import (
	"github.com/sdragos3/taskgo/cmd"
	_ "github.com/sdragos3/taskgo/cmd/create"
	_ "github.com/sdragos3/taskgo/cmd/init"
)

func main() {
	cmd.Execute()
}
