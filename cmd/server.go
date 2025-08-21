package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sliplane-cli/internal/api"
	"strings"

	"time"

	"github.com/briandowns/spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var createServerCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new server",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		locationOptions := []struct {
			Code, Label string
		}{
			{"fsn", "Falkenstein, GER (fsn)"},
			{"nbg", "Nuremberg, GER (nbg)"},
			{"ash", "Ashburn, VA, US (ash)"},
			{"hil", "Hillsboro, OR, US (hil)"},
			{"hel", "Helsinki, FIN (hel)"},
			{"sin", "Singapore, SIN (sin)"},
		}
		locationLabels := make([]string, len(locationOptions))
		for i, l := range locationOptions {
			locationLabels[i] = l.Label
		}
		prompt := promptui.Select{
			Label: "Select Server Location",
			Items: locationLabels,
		}
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}
		location := locationOptions[idx].Code

		serverTypeOptions := []struct {
			Code, Label string
		}{
			{"base", "Base"},
			{"medium", "Medium"},
			{"large", "Large"},
			{"x-large", "X-Large"},
			{"xx-large", "XX-Large"},
			{"dedicated-base", "Dedicated Base"},
			{"dedicated-medium", "Dedicated Medium"},
			{"dedicated-large", "Dedicated Large"},
			{"dedicated-x-large", "Dedicated X-Large"},
			{"dedicated-xx-large", "Dedicated XX-Large"},
			{"dedicated-xxx-large", "Dedicated XXX-Large"},
		}
		serverTypeLabels := make([]string, len(serverTypeOptions))
		for i, t := range serverTypeOptions {
			serverTypeLabels[i] = t.Label
		}
		typePrompt := promptui.Select{
			Label: "Select Server Type",
			Items: serverTypeLabels,
		}
		typeIdx, _, err := typePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}
		serverType := serverTypeOptions[typeIdx].Code

		fmt.Print("Server Name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		sliplaneApiClient := api.GetClient()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Creating server..."
		s.Start()
		server, err := sliplaneApiClient.CreateServer(name, serverType, location)
		s.Stop()
		if err != nil {
			fmt.Printf("Error creating server: %v\n", err)
			return
		}
		fmt.Printf("Server created: %s (%s)\n", server.Name, server.ID)
	},
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Interact with servers in your Sliplane account",
}

var listServersCmd = &cobra.Command{
	Use:   "list",
	Short: "List all servers",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		sliplaneApiClient := api.GetClient()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading servers..."
		s.Start()
		serversList, err := sliplaneApiClient.ListServers()
		s.Stop()
		if err != nil {
			fmt.Printf("Error listing servers: %v\n", err)
			return
		}
		if len(serversList) == 0 {
			fmt.Printf("No Servers found.")
		} else {
			for _, server := range serversList {
				fmt.Printf("- %s (%s)\n", server.Name, server.ID)
			}
		}

	},
}

var deleteServerCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a server",
	Run: func(cmd *cobra.Command, args []string) {
		sliplaneApiClient := api.GetClient()
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Loading servers..."
		s.Start()
		serversList, err := sliplaneApiClient.ListServers()
		s.Stop()
		if err != nil {
			fmt.Printf("Error loading servers: %v\n", err)
			return
		}
		if len(serversList) == 0 {
			fmt.Println("No servers found to delete.")
			return
		}
		serverLabels := make([]string, len(serversList))
		for i, server := range serversList {
			serverLabels[i] = fmt.Sprintf("%s (%s)", server.Name, server.ID)
		}
		prompt := promptui.Select{
			Label: "Select server to delete",
			Items: serverLabels,
		}
		idx, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}
		serverID := serversList[idx].ID
		confirmPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("Are you sure you want to delete server %s? Type 'yes' to confirm", serverLabels[idx]),
			IsConfirm: true,
		}
		confirm, err := confirmPrompt.Run()
		if err != nil || (strings.ToLower(confirm) != "yes" && strings.ToLower(confirm) != "y") {
			fmt.Println("Aborted.")
			return
		}
		s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Deleting server..."
		s.Start()
		err = sliplaneApiClient.DeleteServer(serverID)
		s.Stop()
		if err != nil {
			fmt.Printf("Error deleting server: %v\n", err)
			return
		}
		fmt.Println("Server deleted successfully.")
	},
}

func init() {
	serverCmd.AddCommand(listServersCmd)
	serverCmd.AddCommand(createServerCmd)
	serverCmd.AddCommand(deleteServerCmd)
	rootCmd.AddCommand(serverCmd)
}
