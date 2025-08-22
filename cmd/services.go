package cmd

import (
	"fmt"
	"sliplane-cli/internal/api"

	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Interact with service in your Sliplane account",
}

var listServicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all services",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sliplaneApiClient := api.GetClient()
		// Spinner for fetching projects
		projectSpinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		projectSpinner.Suffix = " Loading projects..."
		projectSpinner.Start()
		projects, err := sliplaneApiClient.ListProjects()
		projectSpinner.Stop()
		if err != nil {
			fmt.Printf("Error fetching projects: %v\n", err)
			return
		}
		if len(projects) == 0 {
			fmt.Println("No projects found.")
			return
		}
		projectLabels := make([]string, len(projects))
		for i, project := range projects {
			projectLabels[i] = fmt.Sprintf("%s (%s)", project.Name, project.ID)
		}
		prompt := promptui.Select{
			Label: "Select project",
			Items: projectLabels,
		}
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}
		projectID := projects[idx].ID
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading services..."
		s.Start()
		servicesList, err := sliplaneApiClient.ListServices(projectID)
		s.Stop()
		if err != nil {
			fmt.Printf("Error listing services: %v\n", err)
			return
		}
		if len(servicesList) == 0 {
			fmt.Printf("No Services found.")
		} else {
			for _, service := range servicesList {
				fmt.Printf("- %s (%s)\n", service.Name, service.ID)
			}
		}
	},
}

func init() {
	serviceCmd.AddCommand(listServicesCmd)
	rootCmd.AddCommand(serviceCmd)
}
