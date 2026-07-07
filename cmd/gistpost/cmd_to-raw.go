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

// *** Sub-command: to-raw ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The ToRawCommand type defines all the configurable options from cli.
// type ToRawCommand struct {
// }
type ToRawCommand struct {
	gistpost.ToRawCommand
}

var to_rawCommand ToRawCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("to-raw",
		"Get raw content url from GH gist",
		"Usage:\n  gistpost [Options] to-raw gist-url",
		&to_rawCommand)
}

func (x *ToRawCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Get raw content url from GH gist\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
	clis.Setup("gistpost::to-raw", gistpost.Opts.Verbose)
	clis.Verbose(1, "Doing ToRaw, with %+v, %+v", gistpost.Opts, args)
	// fmt.Println()
	r, err := x.Exec(args)
	fmt.Print(r)
	return err
}

// // Exec implements the business logic of command `to-raw`
// func (x *ToRawCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("to-raw::Exec", err)
// 	// or,
// 	// clis.AbortOn("to-raw::Exec", err)
// 	return nil
// }
