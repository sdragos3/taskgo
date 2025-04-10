package main

import (
	"github.com/sdragos3/taskgo/cmd"
	_ "github.com/sdragos3/taskgo/cmd/create"
	_ "github.com/sdragos3/taskgo/cmd/delete"
	_ "github.com/sdragos3/taskgo/cmd/init"
	_ "github.com/sdragos3/taskgo/cmd/list"
	_ "github.com/sdragos3/taskgo/cmd/update"
)

func main() {
	cmd.Execute()
}
