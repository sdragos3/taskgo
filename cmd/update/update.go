package update

import (
	"fmt"
	"github.com/google/uuid"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/models"
	"github.com/sdragos3/taskgo/persistence/repositories"
	"github.com/spf13/cobra"
)

var id string
var title string
var description string
var priority models.Priority

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update general information for an existing task (title, description, priority)",
	Run: func(cmd *cobra.Command, args []string) {
		guid, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Invalid ID")
		}
		task, err := repositories.GetById(guid)
		if err != nil {
			fmt.Println(err)
		}

		if cmd.Flags().Changed("title") {
			task.Title = title
		}
		if cmd.Flags().Changed("description") {
			task.Description = description
		}
		if cmd.Flags().Changed("priority") {
			task.Priority = priority
		}
		err = repositories.Insert(task)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the task to update")
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "task title")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "task description")
	updateCmd.Flags().VarP(&priority, "priority", "p", "task priority")

	err := updateCmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}
	root.RootCmd.AddCommand(updateCmd)
}
