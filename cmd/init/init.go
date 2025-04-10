package init

import (
	"fmt"
	root "github.com/sdragos3/taskgo/cmd"
	"github.com/sdragos3/taskgo/persistence/database"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize taskgo",
	Long:  `Initializes taskgo database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		err := database.Init()
		if err != nil {
			fmt.Println("Error initializing database:", err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(initCmd)
}
