////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"log"
)

// *** Sub-command: update ***
// Exec implements the business logic of command `update`
func (x *UpdateCommand) Exec(args []string) error {
	// err := ...
	// clis.WarnOn("update::Exec", err)
	// or,
	// clis.AbortOn("update::Exec", err)
	gop := x.gistPrep(readStdin())
	return gistAction(gop)
}

func (x *UpdateCommand) gistPrep(content []byte) gistOp {
	gist := gistT{
		Description: opts.Description,
		Files: map[string]gistFile{
			opts.Filename: {Content: string(content)},
		},
	}

	// Convert gist to JSON
	gistJson, err := json.Marshal(gist)
	if err != nil {
		log.Fatalf("Error marshaling Gist JSON: %v", err)
	}

	return gistOp{"PATCH", "https://api.github.com/gists/" + x.GistID, gistJson}
}
