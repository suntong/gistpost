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

// *** Sub-command: folder ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The FolderCommand type defines all the configurable options from cli.
type FolderCommand struct {
	Dir    string `short:"D" long:"dir" description:"Directory to upload as gist*" required:"true"`
	Public bool   `short:"p" long:"pub" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
	Extra  bool   `short:"e" long:"extra" env:"GISTPOST_EXTRA" description:"Extra files will be added to gist for better name/doc."`
}

var folderCommand FolderCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("folder",
		"Upload the whole folder as GH gist",
		"Usage:\n  gistpost [Options] folder [-p] [-e]",
		&folderCommand)
}

func (x *FolderCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Upload the whole folder as GH gist\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2024, Tong Sun\n\n")
	clis.Setup("gistpost::folder", opts.Verbose)
	clis.Verbose(1, "Doing Folder, with %+v, %+v", opts, args)
	// fmt.Println(x.Dir, x.Public, x.Extra)
	return x.Exec(args)
}

// // Exec implements the business logic of command `folder`
// func (x *FolderCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("folder::Exec", err)
// 	// or,
// 	// clis.AbortOn("folder::Exec", err)
// 	return nil
// }
