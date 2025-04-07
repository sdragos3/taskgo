package create

import (
	"fmt"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/cmd/types"
	"github.com/spf13/cobra"
)

var title string
var description string
var priority types.Priority = types.None

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new task",
	Long:  `Create new task`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating task with title '%s' and description '%s' and priority '%s'\n", title, description, priority)
	},
}

func init() {
	createCmd.Flags().StringVarP(&title, "title", "t", "", "task title")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "task description")
	createCmd.Flags().VarP(&priority, "priority", "p", "task priority")

	createCmd.MarkFlagRequired("title")
	root.RootCmd.AddCommand(createCmd)
}
