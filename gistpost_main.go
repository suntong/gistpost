package main

//go:generate sh gistpost_cliGen.sh

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type GistFile struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string              `json:"description"`
	Public      bool                `json:"public"`
	Files       map[string]GistFile `json:"files"`
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	// define flags
	initVars()
	// popoulate flag variables from ENV
	initVals()
	// popoulate flag variables from cli
	flag.Parse()
	if Opts.Help {
		Usage(1)
	}

	// == Sanity check on variables from environment
	if Opts.Token == "" {
		Usage(0)
		fmt.Println("\nError: The GP_GIST_T_TOKEN environment variable is required")
		os.Exit(1)
	}
	if Opts.Description == "" {
		t := time.Now()
		Opts.Description = "Archived on " + t.Format(time.DateOnly)
	}
	if Opts.Filename == "" {
		Opts.Filename = "archive.md"
	}

	// == Sanity check on stdin
	// Get file information about stdin
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		os.Exit(1)
	}
	// Check if stdin is from a pipe
	if info.Mode()&os.ModeCharDevice != 0 {
		Usage(0)
		fmt.Println("\nError: This program reads input from pipe.")
		os.Exit(1)
	}

	// Read content from stdin
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin: %v", err)
	}

	// Create the gist
	gist := Gist{
		Description: Opts.Description,
		Public:      Opts.Public,
		Files: map[string]GistFile{
			Opts.Filename: {Content: string(content)},
		},
	}

	// Convert Gist to JSON
	gistJson, err := json.Marshal(gist)
	if err != nil {
		log.Fatalf("Error marshaling Gist JSON: %v", err)
	}

	// Make the POST request to GitHub Gist API
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/gists", bytes.NewBuffer(gistJson))
	if err != nil {
		log.Fatalf("Error creating POST request: %v", err)
	}
	req.Header.Set("Authorization", "token "+Opts.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error posting to GitHub Gist: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Failed to create gist, status: %s", resp.Status)
	}

	// Print the URL of the created gist
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Gist created: ", result["html_url"])

}

//==========================================================================
// support functions

/*

Usage example:

   export GP_GIST_T_TOKEN=github_pat_...

   $ cat gistpost_main.go | gistpost -d 'gistpost main' -f main.go -p
   Gist created:  https://gist.github.com/suntong/923dabcb17c73e7ba435c0067eee3b19

*/
