////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags/clis"
	"github.com/suntong/gistpost"
)

// *** Sub-command: update ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The UpdateCommand type defines all the configurable options from cli.
/*
type UpdateCommand struct {
	GistID string `short:"g" long:"id" env:"GISTPOST_GISTID" description:"Existing GH gist id*" required:"true"`
}
*/

type UpdateCommand struct {
	gistpost.UpdateCommand
}

var updateCommand UpdateCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("update",
		"Update an existing GH gist entry (file)",
		"Usage:\n  gistpost [Options] update --id",
		&updateCommand)
}

func (x *UpdateCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Update an existing GH gist entry (file)\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
	clis.Setup("gistpost::update", gistpost.Opts.Verbose)
	clis.Verbose(1, "Doing Update, with %+v, %+v", gistpost.Opts, args)
	// fmt.Println(x.GistID)

	gistpost.From = os.Stdin
	r, err := x.Exec(args)
	fmt.Print(x.Extract(r))
	return err
}

// // Exec implements the business logic of command `update`
// func (x *UpdateCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("update::Exec", err)
// 	// or,
// 	// clis.AbortOn("update::Exec", err)
// 	return nil
// }
