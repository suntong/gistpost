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

// *** Sub-command: create ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The CreateCommand type defines all the configurable options from cli.
type CreateCommand struct {
	Public bool `short:"p" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
}

var createCommand CreateCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("create",
		"Create a new GH gist entry (file)",
		"Usage:\n  gistpost [Options] create [-p]",
		&createCommand)
}

func (x *CreateCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Create a new GH gist entry (file)\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2024, Tong Sun\n\n")
	clis.Setup("gistpost::create", opts.Verbose)
	clis.Verbose(1, "Doing Create, with %+v, %+v", opts, args)
	// fmt.Println(x.Public)
	return x.Exec(args)
}

// // Exec implements the business logic of command `create`
// func (x *CreateCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("create::Exec", err)
// 	// or,
// 	// clis.AbortOn("create::Exec", err)
// 	return nil
// }
