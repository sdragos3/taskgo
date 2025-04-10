package complete

import (
	"fmt"
	"github.com/google/uuid"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/persistence/repositories"
	"github.com/spf13/cobra"
)

var id string

var completionCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		guid, err := uuid.Parse(id)
		if err != nil {
			fmt.Println("Invalid ID")
		}
		task, err := repositories.GetById(guid)
		if err != nil {
			fmt.Println(err)
		}

		task.Completed = true
		err = repositories.Insert(task)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	completionCmd.Flags().StringVarP(&id, "id", "i", "", "Task ID to complete")

	err := completionCmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}

	root.RootCmd.AddCommand(completionCmd)
}
