package cmd

import (
	"fmt"
	"sliplane-cli/internal/api"

	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var credentialsCmd = &cobra.Command{
	Use:   "credentials",
	Short: "Interact with credentials in your Sliplane account",
}

var listCredentialsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all credentials",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		sliplaneApiClient := api.GetClient()

		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading credentials..."
		s.Start()
		credentialsList, err := sliplaneApiClient.ListCredentials()
		s.Stop()
		if err != nil {
			fmt.Printf("Error listing credentials: %v\n", err)
			return
		}
		if len(credentialsList) == 0 {
			fmt.Printf("No Credentials found.")
		} else {
			for _, credential := range credentialsList {
				fmt.Printf("- %s (%s)\n", credential.Name, credential.ID)
			}
		}
	},
}

func init() {
	credentialsCmd.AddCommand(listCredentialsCmd)
	rootCmd.AddCommand(credentialsCmd)
}
