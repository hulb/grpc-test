package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "grpc",
	Short: "a grpc test tool",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)
}
