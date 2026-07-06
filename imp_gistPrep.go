////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-easygen/go-flags/clis"
)

// *** Shared gist-prep helpers ***

// gistOpFromFiles builds the API op from a file map.
//   method=="POST"  -> wrapped as gistCreate (carries Public flag).
//   method=="PATCH" -> wrapped as gistT       (GitHub ignores visibility on PATCH).
func gistOpFromFiles(files map[string]gistFile, method, url string, public bool) gistOp {
	gist := gistT{
		Description: Opts.Description,
		Files:       files,
	}

	var payload []byte
	var err error
	if method == "POST" {
		payload, err = json.Marshal(gistCreate{gist, public})
	} else {
		payload, err = json.Marshal(gist)
	}
	if err != nil {
		log.Fatalf("Error marshaling Gist JSON: %v", err)
	}

	return gistOp{method, url, payload}
}

// walkDir enumerates files in a (flat) directory into a gistFile map.
// Reused by folder (POST) and append (PATCH). Nested subfolders are
// rejected, matching the original imp_folder.go behavior.
func walkDir(dirPath string, extra bool) (map[string]gistFile, error) {
	files := make(map[string]gistFile)

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
				return errors.New("No folders allowed within gist")
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error collecting files in given directory: %w", err)
	}

	if extra {
		// add dirname to the top of the file list, and zzComments.md to the bottom
		// with empty content
		abs, _ := filepath.Abs(dirPath)
		files["00_"+filepath.Base(abs)+".md"] = gistFile{"."}
		files["zzComments.md"] = gistFile{"."}
	}

	return files, nil
}
