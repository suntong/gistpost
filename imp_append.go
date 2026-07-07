////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import (
	"fmt"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: append ***
// Exec implements the business logic of command `append`.
//
// `append` walks the given directory and issues a PATCH against the
// existing gist. GitHub's Gist API treats each file name in the PATCH
// `files` map as:
//   - a NEW file name -> the file is ADDED to the gist (appended), and
//   - an EXISTING file name -> the file content is REPLACED.
// So as long as the appended folder's file names don't collide with the
// existing gist's file names, every file is appended. This is the same
// collision caveat that the `folder` command has internally.
func (x *AppendCommand) Exec(args []string) (GistRet, error) {
	files, err := walkDir(x.Dir, x.Extra)
	if err != nil {
		return nil, err
	}
	gop := gistOpFromFiles(
		files,
		"PATCH", "https://api.github.com/gists/"+x.GistID, false)
	result := gistAction(gop)
	clis.Verbose(3, "Got %+v", result)
	return result, nil
}

func (x *AppendCommand) Extract(result GistRet) string {
	return fmt.Sprintf("Gist appended: %v\n", result["html_url"])
}

/*

Usage example:

   export GISTPOST_TOKEN=github_pat_...

   $ gistpost folder -D a/ -d 'demo' -p
   Gist git url: https://gist.github.com/suntong/<GID>
   ...

   $ gistpost append -g <GID> -D b/
   Gist appended: https://gist.github.com/suntong/<GID>

   The gist URL stays the same; the files from b/ have been added to it.

*/
