package cmd

import (
	"github.com/spf13/cobra"
)

var version = "latest"
var commit = "head"
var buildTime = "now"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return version information of taskgo",
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
