package init

import (
	root "github.com/sdragos3/taskgo/cmd"
	database "github.com/sdragos3/taskgo/persistence"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize taskgo",
	Long:  `Initializes taskgo database`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Initialize()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(createCmd)
}
