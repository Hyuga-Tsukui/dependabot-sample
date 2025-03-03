package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints hello world",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello world")
		return nil
	},
}
