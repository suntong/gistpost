// gistpost - GH gist post/update tool
//
// Tool to post to GH gist, or update it

package main

////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

import (
//  	"fmt"
//  	"os"

// "github.com/go-easygen/go-flags"
)

// Template for main starts here

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "gistpost"
//          version   = "0.1.0"
//          date = "2024-09-02"

//  	// opts store all the configurable options
//  	opts optsT
//  )
//
//  var gfParser = flags.NewParser(&opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// Function main
//  func main() {
//  	opts.Version = showVersion
//  	opts.Verbflg = func() {
//  		opts.Verbose++
//  	}
//
//  	if _, err := gfParser.Parse(); err != nil {
//  		fmt.Println()
//  		gfParser.WriteHelp(os.Stdout)
//  		os.Exit(1)
//  	}
//  	fmt.Println()
//  	//DoGistpost()
//  }
//
//  //==========================================================================
//  // support functions
//
//  func showVersion() {
//   	fmt.Fprintf(os.Stderr, "gistpost - GH gist post/update tool, version %s\n", version)
//  	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
//   	fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2024, Tong Sun\n\n")
//  	fmt.Fprintf(os.Stderr, "Tool to post to GH gist, or update it\n")
//  	os.Exit(0)
//  }
// Template for main ends here

// DoGistpost implements the business logic of command `gistpost`
//  func DoGistpost() error {
//  	return nil
//  }

// Template for type define starts here

// The optsT type defines all the configurable options from cli.
type optsT struct {
	Token       string `short:"t" long:"token" env:"GISTPOST_TOKEN" description:"The GITHUB_TOKEN*" required:"true"`
	Description string `short:"d" long:"desc" env:"GISTPOST_DESCRIPTION" description:"Gist description"`
	Filename    string `short:"f" long:"fname" env:"GISTPOST_FILENAME" description:"Gist filename" default:"archive.md"`
	Wrap        bool   `short:"w" long:"wrap" env:"GISTPOST_WRAP" description:"Wrap content within a markdown block"`
	Verbflg     func() `short:"v" long:"verbose" description:"Verbose mode (Multiple -v options increase the verbosity)"`
	Verbose     int
	Version     func() `short:"V" long:"version" description:"Show program version and exit"`
}

// Template for type define ends here

// Template for "create" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Tong Sun (c) 2024-2024, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package main

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: create ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The CreateCommand type defines all the configurable options from cli.
//  type CreateCommand struct {
//  	Public	bool	`short:"p" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
//  }

//
//  var createCommand CreateCommand
//
//  ////////////////////////////////////////////////////////////////////////////
//  // Function definitions
//
//  func init() {
//  	gfParser.AddCommand("create",
//  		"Create a new GH gist entry",
//  		"Usage:\n  gistpost [Options] create [-p]",
//  		&createCommand)
//  }
//
//  func (x *CreateCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Create a new GH gist entry\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2024, Tong Sun\n\n")
//   	clis.Setup("gistpost::create", opts.Verbose)
//   	clis.Verbose(1, "Doing Create, with %+v, %+v", opts, args)
//   	// fmt.Println(x.Public)
//  	return x.Exec(args)
//  }
//
// // Exec implements the business logic of command `create`
// func (x *CreateCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("create::Exec", err)
// 	// or,
// 	// clis.AbortOn("create::Exec", err)
// 	return nil
// }
// Template for "create" CLI handling ends here
