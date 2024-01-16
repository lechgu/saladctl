package commands

import "github.com/spf13/cobra"

var containerGroupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get container group",
	RunE:  getContainerGroup,
}

func getContainerGroup(cmd *cobra.Command, args []string) error {
	return nil
}
