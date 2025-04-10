package delete

import (
	"fmt"
	"github.com/google/uuid"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/persistence/repositories"
	"github.com/spf13/cobra"
)

var id string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		guid, err := uuid.Parse(id)
		if err != nil {
			fmt.Printf("Invalid ID: %s", id)
		}
		err = repositories.DeleteById(guid)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	deleteCmd.Flags().StringVarP(&id, "id", "i", "", "Task ID")

	err := deleteCmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}

	root.RootCmd.AddCommand(deleteCmd)
}
