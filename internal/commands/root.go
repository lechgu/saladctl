package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "saladctl",
	Short: "CLI to interact with the Salad Portal",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
