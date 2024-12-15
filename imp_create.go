////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-easygen/go-flags/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type gistFile struct {
	Content string `json:"content"`
}

type gistT struct {
	Description string              `json:"description"`
	Files       map[string]gistFile `json:"files"`
}

type gistCreate struct {
	gistT
	Public bool `json:"public"`
}

type gistOp struct {
	method, url string
	gistJson    []byte
}

type GistRet map[string]any

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

////////////////////////////////////////////////////////////////////////////
// Function definitions

func OptsCheck() error {
	// == Sanity check on variables from environment
	if Opts.Token == "" {
		return fmt.Errorf("\nError: GISTPOST_TOKEN environment variable required")
	}
	if Opts.Description == "" {
		t := time.Now()
		Opts.Description = "Archived on " + t.Format(time.DateOnly)
	}
	if Opts.Filename == "" {
		Opts.Filename = "archive.md"
	}
	// Not checking .From. Make sure assigned!!
	return nil
}

// // Exec implements the business logic of command `create`
func (x *CreateCommand) Exec(args []string) (GistRet, error) {
	// err := ...
	// clis.WarnOn("create::Exec", err)
	// or,
	// clis.AbortOn("create::Exec", err)

	gop := x.gistPrep(readStdin())
	result := gistAction(gop)
	return result, nil
}

func (x *CreateCommand) Extract(result GistRet) string {
	return fmt.Sprintf("Gist created: %v\n", result["html_url"])
}

func readStdin() []byte {
	// Read content from stdin
	content, err := io.ReadAll(From)
	if err != nil {
		log.Fatalf("Error reading stdin: %v", err)
	}

	if Opts.Wrap {
		content = []byte("```\n" + string(content) + "\n```\n")
	}
	return content
}

func (x *CreateCommand) gistPrep(content []byte) gistOp {
	gist := gistT{
		Description: Opts.Description,
		Files: map[string]gistFile{
			Opts.Filename: {Content: string(content)},
		},
	}

	// Create-gist
	gc := gistCreate{gist, x.Public}

	// Convert gist to JSON
	gistJson, err := json.Marshal(gc)
	if err != nil {
		log.Fatalf("Error marshaling Gist JSON: %v", err)
	}

	return gistOp{"POST", "https://api.github.com/gists", gistJson}
}

func gistAction(gop gistOp) GistRet {
	clis.Verbose(3, "%s Requesting to GitHub %s with %+v",
		gop.method, gop.url, string(gop.gistJson))

	// Make the request to GitHub Gist API
	client := &http.Client{}
	req, err := http.NewRequest(gop.method, gop.url, bytes.NewBuffer(gop.gistJson))
	if err != nil {
		log.Fatalf("Error creating %s request: %v", gop.method, err)
	}
	req.Header.Set("Authorization", "token "+Opts.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error posting to GitHub Gist: %v", err)
	}
	defer resp.Body.Close()

	if !(resp.StatusCode == http.StatusCreated ||
		resp.StatusCode == http.StatusOK) {
		// Read the response body
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Fatalf("Failed to %s to gist, status: %s\n  %s",
			gop.method, resp.Status, bodyBytes)
	}

	// Get the URL of the gist
	var result GistRet
	json.NewDecoder(resp.Body).Decode(&result)

	return result
}

//==========================================================================
// support functions

/*

Usage example:

   export GISTPOST_TOKEN=github_pat_...

   $ cat imp_create.go | gistpost create -d 'gistpost main' -f main.go -p
   Gist created:  https://gist.github.com/suntong/d3203029f4dd088e3b5ed2b0aaee652c

*/
