// gistpost - GH gist post/update tool

// Tool to post to GH gist, or update it

package main

////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh gistpost_cliGen.sh
//go:generate emd gen -in README.beg.e.md -in README.e.md -in README.end.e.md -out README.md

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags"
	"github.com/suntong/gistpost"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "gistpost"
	version  = "2.1.0"
	date     = "2024-12-15"

	// gistpost.Opts store all the configurable options
	//gistpost.Opts gistpost.OptsT
)

var gfParser = flags.NewParser(&gistpost.Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// ==========================================================================
// Function main
func main() {
	gistpost.Opts.Version = showVersion
	gistpost.Opts.Verbflg = func() {
		gistpost.Opts.Verbose++
	}

	if _, err := gfParser.Parse(); err != nil {
		fmt.Println()
		gfParser.WriteHelp(os.Stdout)
		//fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	//DoGistpost()
}

//==========================================================================
// support functions

func showVersion() {
	fmt.Fprintf(os.Stderr, "gistpost - GH gist post/update tool, version %s\n", version)
	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
	fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2024, Tong Sun\n\n")
	fmt.Fprintf(os.Stderr, "Tool to post to GH gist, or update it\n")
	os.Exit(0)
}
