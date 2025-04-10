package main

import (
	"github.com/sdragos3/taskgo/cmd"
	_ "github.com/sdragos3/taskgo/cmd/create"
	_ "github.com/sdragos3/taskgo/cmd/init"
	_ "github.com/sdragos3/taskgo/cmd/list"
)

func main() {
	cmd.Execute()
}
