package create

import (
	"fmt"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/models"
	"github.com/spf13/cobra"
)

var title string
var description string
var priority models.Priority = models.None

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new task",
	Long:  `Create new task`,
	Run: func(cmd *cobra.Command, args []string) {
		task := models.TaskCreate(title, description, priority)
		fmt.Printf("%s", task)
	},
}

func init() {
	createCmd.Flags().StringVarP(&title, "title", "t", "", "task title")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "task description")
	createCmd.Flags().VarP(&priority, "priority", "p", "task priority")

	err := createCmd.MarkFlagRequired("title")
	if err != nil {
		panic(err)
	}
	root.RootCmd.AddCommand(createCmd)
}
