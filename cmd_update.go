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

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The UpdateCommand type defines all the configurable options from cli.
type UpdateCommand struct {
	GistID string `short:"g" long:"id" env:"GISTPOST_GISTID" description:"Existing GH gist id*" required:"true"`
}

var updateCommand UpdateCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("update",
		"Update an existing GH gist entry",
		"Usage:\n  gistpost [Options] update --id",
		&updateCommand)
}

func (x *UpdateCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Update an existing GH gist entry\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2024, Tong Sun\n\n")
	clis.Setup("gistpost::update", opts.Verbose)
	clis.Verbose(1, "Doing Update, with %+v, %+v", opts, args)
	// fmt.Println(x.GistID)
	return x.Exec(args)
}

// // Exec implements the business logic of command `update`
// func (x *UpdateCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("update::Exec", err)
// 	// or,
// 	// clis.AbortOn("update::Exec", err)
// 	return nil
// }
