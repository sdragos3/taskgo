package update

import (
	"fmt"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/models"
	"github.com/spf13/cobra"
)

var id string
var title *string
var description *string
var priority *models.Priority

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update general information for an existing task (title, description, priority)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Update")
	},
}

func init() {
	updateCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the task to update")
	updateCmd.Flags().StringVarP(title, "title", "t", "", "task title")
	updateCmd.Flags().StringVarP(description, "description", "d", "", "task description")
	updateCmd.Flags().VarP(priority, "priority", "p", "task priority")

	updateCmd.MarkFlagRequired("id")
	root.RootCmd.AddCommand(updateCmd)
}
