////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import (
	"encoding/json"
	"fmt"
	"log"
)

// *** Sub-command: update ***
// Exec implements the business logic of command `update`
func (x *UpdateCommand) Exec(args []string) (GistRet, error) {
	// err := ...
	// clis.WarnOn("update::Exec", err)
	// or,
	// clis.AbortOn("update::Exec", err)
	gop := x.gistPrep(readStdin())
	result := gistAction(gop)
	return result, nil
}

func (x *UpdateCommand) Extract(result GistRet) string {
	return fmt.Sprintf("Gist updated: %v\n", result["html_url"])
}

func (x *UpdateCommand) gistPrep(content []byte) gistOp {
	gist := gistT{
		Description: Opts.Description,
		Files: map[string]gistFile{
			Opts.Filename: {Content: string(content)},
		},
	}

	// Convert gist to JSON
	gistJson, err := json.Marshal(gist)
	if err != nil {
		log.Fatalf("Error marshaling Gist JSON: %v", err)
	}

	return gistOp{"PATCH", "https://api.github.com/gists/" + x.GistID, gistJson}
}

/*

Usage example:

   export GISTPOST_TOKEN=github_pat_...

   $ echo abc | gistpost create -w -d 'gistpost test create' -p
   Gist url:  https://gist.github.com/suntong/25c53ccc65330a8c20f8ba79d4a9eed4

   $ echo abcd | gistpost update -w -d 'gistpost test update' -g 25c53ccc65330a8c20f8ba79d4a9eed4
   Gist url:  https://gist.github.com/suntong/25c53ccc65330a8c20f8ba79d4a9eed4

*/
