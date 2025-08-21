package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "v1.0.0",
	Use:     "sliplane",
	Short:   "Deploy to your sliplane account",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Sliplane '%s'\n", err)
		os.Exit(1)
	}
}
