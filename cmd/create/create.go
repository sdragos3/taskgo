package create

import (
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/models"
	repository "github.com/sdragos3/taskgo/persistence/repositories"
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
		err := repository.Create(&task)
		if err != nil {
			panic(err)
		}
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
