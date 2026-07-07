////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: folder ***
// Exec implements the business logic of command `folder`
func (x *FolderCommand) Exec(args []string) (GistRet, error) {
	files, err := walkDir(x.Dir, x.Extra)
	if err != nil {
		return nil, err
	}
	gop := gistOpFromFiles(
		files,
		"POST", "https://api.github.com/gists", x.Public)
	result := gistAction(gop)
	clis.Verbose(3, "Got %+v", result)
	return result, nil
}

func (x *FolderCommand) Extract(result GistRet) string {
	//if err != nil { return err }
	url_http := result["git_push_url"]
	url_git := strings.Replace(url_http.(string),
		"https://gist.github.com/", "git@gist.github.com:", 1)
	abs, _ := filepath.Abs(x.Dir)
	return fmt.Sprintf("Gist git url: %s\n cd ../\n mv -v %s{,.org}\n git clone %s %[2]s\n cd %[2]s\n",
		url_http, filepath.Base(abs), url_git)
}
