package cmd

import (
	"github.com/spf13/cobra"
	"wxjkzcy/consts"
	"wxjkzcy/runner"
)

// cakeCmd represents the unicake command
var cakeCmd = &cobra.Command{
	Use:   "mint",
	Short: "mint",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		runner.NewEthRunner().Start(consts.ChainTypeBsc)
	},
}

func init() {
	rootCmd.AddCommand(cakeCmd)
}
