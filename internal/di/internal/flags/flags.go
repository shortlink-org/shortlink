package flags

import "github.com/spf13/cobra"

func New() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	return rootCmd, nil
}
