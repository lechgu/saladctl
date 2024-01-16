package commands

import "github.com/spf13/cobra"

var organizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "Manage organizations",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(organizationCmd)
}
