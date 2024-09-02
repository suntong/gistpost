////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: update ***
// Exec implements the business logic of command `update`
func (x *UpdateCommand) Exec(args []string) error {
	// err := ...
	// clis.WarnOn("update::Exec", err)
	// or,
	// clis.AbortOn("update::Exec", err)
	return nil
}
