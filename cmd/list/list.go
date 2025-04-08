package list

import (
	"fmt"
	root "github.com/sdragos3/taskgo/cmd"
	repository "github.com/sdragos3/taskgo/persistence/repositories"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all available commands",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := repository.List()
		if err != nil {
			fmt.Println("Error listing commands:", err)
		}
		for _, task := range tasks {
			fmt.Println(task.ToString())
		}
	},
}

func init() {
	root.RootCmd.AddCommand(listCmd)
}
