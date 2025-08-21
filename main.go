package main

import (
	"os"
	"sliplane-cli/cmd"
	"sliplane-cli/internal/api"

	"github.com/joho/godotenv"
)

func main() {

	// Initialize the API client with environment variables

	godotenv.Load()
	api.Init(os.Getenv("SLIPLANE_API_KEY"), os.Getenv("SLIPLANE_ORG_ID"))
	cmd.Execute()
}
