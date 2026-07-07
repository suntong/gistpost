////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import (
	"fmt"
)

// *** Sub-command: update ***
// Exec implements the business logic of command `update`
func (x *UpdateCommand) Exec(args []string) (GistRet, error) {
	files := map[string]gistFile{
		Opts.Filename: {Content: string(readStdin())},
	}
	gop := gistOpFromFiles(
		files, "PATCH", "https://api.github.com/gists/"+x.GistID, false)
	result := gistAction(gop)
	return result, nil
}

func (x *UpdateCommand) Extract(result GistRet) string {
	return fmt.Sprintf("Gist updated: %v\n", result["html_url"])
}

/*

Usage example:

   export GISTPOST_TOKEN=github_pat_...

   $ echo abc | gistpost create -w -d 'gistpost test create' -p
   Gist url:  https://gist.github.com/suntong/25c53ccc65330a8c20f8ba79d4a9eed4

   $ echo abcd | gistpost update -w -d 'gistpost test update' -g 25c53ccc65330a8c20f8ba79d4a9eed4
   Gist url:  https://gist.github.com/suntong/25c53ccc65330a8c20f8ba79d4a9eed4

*/
