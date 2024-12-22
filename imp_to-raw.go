////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

import "strings"

type ToRawCommand struct {
}

// *** Sub-command: to-raw ***
// Exec implements the business logic of command `to-raw`
func (x *ToRawCommand) Exec(args []string) (string, error) {
	// err := ...
	// clis.WarnOn("to-raw::Exec", err)
	// or,
	// clis.AbortOn("to-raw::Exec", err)
	r := ""
	for _, gistURL := range args {
		r += gistToRawURL(gistURL) + "/raw\n"
	}
	return r, nil
}

func gistToRawURL(gistURL string) string {
	if !strings.HasPrefix(gistURL, "https://gist.github.com") {
		return "Invalid Gist URL of " + gistURL
	}

	rawURL := strings.Replace(gistURL, "gist.github.com", "gist.githubusercontent.com", 1)
	// rawURL = strings.Replace(rawURL, "/gist/", "/raw/", 1)
	return rawURL
}
