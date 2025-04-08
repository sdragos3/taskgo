package clean

import (
	root "github.com/sdragos3/taskgo/cmd"
	database "github.com/sdragos3/taskgo/persistence"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean taskgo database",
	Long:  `Clean taskgo database`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Delete()
		if err != nil {
			panic(err)
		}
		err = database.Initialize()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(cleanCmd)
}
