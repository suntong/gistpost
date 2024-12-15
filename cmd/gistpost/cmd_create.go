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
	"github.com/suntong/gistpost"
)

// *** Sub-command: create ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The CreateCommand type defines all the configurable options from cli.
/*
type CreateCommand struct {
	Public bool `short:"p" long:"pub" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
}
*/

type CreateCommand struct {
	gistpost.CreateCommand
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
	clis.Setup("gistpost::create", gistpost.Opts.Verbose)
	clis.Verbose(1, "Doing Create, with %+v, %+v", gistpost.Opts, args)
	// fmt.Println(x.Public)

	// == Sanity check on stdin
	// Get file information about stdin
	info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("Error checking stdin:", err)
		os.Exit(1)
	}
	// Check if stdin is from a pipe
	if info.Mode()&os.ModeCharDevice != 0 {
		gfParser.WriteHelp(os.Stdout)
		fmt.Println("\nError: This program reads input from pipe.")
		os.Exit(1)
	}
	if err = gistpost.OptsCheck(); err != nil {
		gfParser.WriteHelp(os.Stdout)
		fmt.Println(err)
		os.Exit(1)
	}

	gistpost.From = os.Stdin
	r, err := x.Exec(args)
	fmt.Print(x.Extract(r))
	return err
}

// // Exec implements the business logic of command `create`
// func (x *CreateCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("create::Exec", err)
// 	// or,
// 	// clis.AbortOn("create::Exec", err)
// 	return nil
// }
