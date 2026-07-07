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

// *** Sub-command: append ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The AppendCommand type defines all the configurable options from cli.
/*
type AppendCommand struct {
	GistID string `short:"g" long:"id" env:"GISTPOST_GISTID" description:"Existing GH gist id*" required:"true"`
	Dir    string `short:"D" long:"dir" description:"Directory whose files to append*" required:"true"`
	Extra  bool   `short:"e" long:"extra" env:"GISTPOST_EXTRA" description:"Extra files will be added to gist for better name/doc."`
}
*/

type AppendCommand struct {
	gistpost.AppendCommand
}

var appendCommand AppendCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("append",
		"Append all given folder files to an existing GH gist entry",
		`Usage:\n  gistpost [Options] append -g <gid> -D <dir> [-e]
`,
		&appendCommand)
}

func (x *AppendCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Append all given folder files to an existing GH gist entry\n")
	clis.Setup("gistpost::append", gistpost.Opts.Verbose)
	clis.Verbose(1, "Doing Append, with %+v, %+v", gistpost.Opts, args)
	r, err := x.Exec(args)
	fmt.Print(x.Extract(r))
	return err
}

// // Exec implements the business logic of command `append`
// func (x *AppendCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("append::Exec", err)
// 	// or,
// 	// clis.AbortOn("append::Exec", err)
// 	return nil
// }
