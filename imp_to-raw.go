////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package gistpost

type ToRawCommand struct {
}

// *** Sub-command: to-raw ***
// Exec implements the business logic of command `to-raw`
func (x *ToRawCommand) Exec(args []string) (string, error) {
	// err := ...
	// clis.WarnOn("to-raw::Exec", err)
	// or,
	// clis.AbortOn("to-raw::Exec", err)
	return "", nil
}
