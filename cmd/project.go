package cmd

import (
	"fmt"
	"sliplane-cli/internal/api"
	"strings"

	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Interact with projects in your Sliplane account",
}

var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		fmt.Print("Project Name: ")
		fmt.Scanln(&name)
		if name == "" {
			fmt.Println("Project name cannot be empty.")
			return
		}
		apiClient := api.GetClient()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Creating project..."
		s.Start()
		project, err := apiClient.CreateProject(name)
		s.Stop()
		if err != nil {
			fmt.Printf("Error creating project: %v\n", err)
			return
		}
		fmt.Printf("Project created: %s (%s)\n", project.Name, project.ID)
	},
}

var listProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := api.GetClient()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading projects..."
		s.Start()
		projectsList, err := apiClient.ListProjects()
		s.Stop()
		if err != nil {
			fmt.Printf("Error listing projects: %v\n", err)
			return
		}
		if len(projectsList) == 0 {
			fmt.Printf("No Projects found.")
		} else {
			for _, project := range projectsList {
				fmt.Printf("- %s (%s)\n", project.Name, project.ID)
			}
		}

	},
}

var deleteProjectCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project",
	Run: func(cmd *cobra.Command, args []string) {
		sliplaneApiClient := api.GetClient()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading projects..."
		s.Start()
		projectsList, err := sliplaneApiClient.ListProjects()
		s.Stop()
		if err != nil {
			fmt.Printf("Error loading projects: %v\n", err)
			return
		}
		if len(projectsList) == 0 {
			fmt.Println("No projects found to delete.")
			return
		}
		projectLabels := make([]string, len(projectsList))
		for i, project := range projectsList {
			projectLabels[i] = fmt.Sprintf("%s (%s)", project.Name, project.ID)
		}
		prompt := promptui.Select{
			Label: "Select project to delete",
			Items: projectLabels,
		}
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}
		projectID := projectsList[idx].ID
		confirmPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("Are you sure you want to delete project %s? Type 'yes' to confirm", projectLabels[idx]),
			IsConfirm: true,
		}
		confirm, err := confirmPrompt.Run()
		if err != nil || (strings.ToLower(confirm) != "yes" && strings.ToLower(confirm) != "y") {
			fmt.Println("Aborted.")
			return
		}
		s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Deleting project..."
		s.Start()
		err = sliplaneApiClient.DeleteProject(projectID)
		s.Stop()
		if err != nil {
			fmt.Printf("Error deleting project: %v\n", err)
			return
		}
		fmt.Println("Project deleted successfully.")
	},
}

func init() {
	projectCmd.AddCommand(deleteProjectCmd)
	projectCmd.AddCommand(listProjectsCmd)
	projectCmd.AddCommand(createProjectCmd)
	rootCmd.AddCommand(projectCmd)
}
