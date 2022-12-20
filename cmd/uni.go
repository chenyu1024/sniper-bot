package cmd

import (
	"github.com/spf13/cobra"
	"wxjkzcy/consts"
	"wxjkzcy/runner"
)

// cakeCmd represents the unicake command
var uniCmd = &cobra.Command{
	Use:   "uni",
	Short: "sniper on uniswap v2",
	Run: func(cmd *cobra.Command, args []string) {
		runner.NewEthRunner().SniperUniCake(consts.ChainTypeEth)
	},
}

func init() {
	rootCmd.AddCommand(uniCmd)
}
