////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: folder ***
// Exec implements the business logic of command `folder`
func (x *FolderCommand) Exec(args []string) error {
	// err := ...
	// clis.WarnOn("folder::Exec", err)
	// or,
	// clis.AbortOn("folder::Exec", err)
	gop := x.gistPrep()
	result := gistAction(gop)
	clis.Verbose(3, "Got %+v", result)
	url_http := result["git_push_url"]
	url_git := strings.Replace(url_http.(string),
		"https://gist.github.com/", "git@gist.github.com:", 1)
	abs, _ := filepath.Abs(x.Dir)
	fmt.Printf("Gist git url: %s\n cd ../\n mv -v %s{,.org}\n git clone %s %[2]s\n cd %[2]s\n",
		url_http, filepath.Base(abs), url_git)
	return nil
}

func (x *FolderCommand) gistPrep() gistOp {
	// Directory to walk through
	dirPath := x.Dir
	extraFiles := x.Extra

	files := make(map[string]gistFile)

	// Walk through the directory and collect file names & contents
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fc := clis.ReadInput(path)
			files[path] = gistFile{string(fc)}
		} else {
			// fmt.Println("Skipping folder", path)
			if path != "." {
				clis.AbortOn("folder::Exec", errors.New("No folders allowed within gist"))
			}
		}
		return nil
	})
	if err != nil {
		clis.AbortOn("folder::Exec", errors.New("Error collecting files in given directory"))
	}

	if extraFiles {
		// add dirname to the top of the file list, and zzComments.md to the bottom
		// with empty content
		abs, _ := filepath.Abs(dirPath)
		files["00_"+filepath.Base(abs)+".md"] = gistFile{"."}
		files["zzComments.md"] = gistFile{"."}
	}

	gist := gistT{
		Description: opts.Description,
		Files:       files,
	}
	gc := gistCreate{gist, x.Public}
	// Convert gist to JSON
	gistJson, err := json.Marshal(gc)
	if err != nil {
		log.Fatalf("Error marshaling Gist JSON: %v", err)
	}

	return gistOp{"POST", "https://api.github.com/gists", gistJson}
}
