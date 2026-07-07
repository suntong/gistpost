// gistpost - GH gist post/update tool
//
// Tool to post to GH gist, or update it

package gistpost

////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

import (
//  	"fmt"
//  	"os"

// "github.com/go-easygen/go-flags"
)

// Template for main starts here

//  // for `go generate -x`
//  //go:generate sh gistpost_cliGen.sh

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "gistpost"
//          version   = "0.1.0"
//          date = "2026-07-07"

//  	// Opts store all the configurable options
//  	Opts OptsT
//  )
//
//  var gfParser = flags.NewParser(&Opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

//==========================================================================
// Function main
//  func main() {
//  	Opts.Version = showVersion
//  	Opts.Verbflg = func() {
//  		Opts.Verbose++
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
//   	fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
//  	fmt.Fprintf(os.Stderr, "Tool to post to GH gist, or update it\n")
//  	os.Exit(0)
//  }
// Template for main ends here

// DoGistpost implements the business logic of command `gistpost`
//  func DoGistpost() error {
//  	return nil
//  }

// Template for type define starts here

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
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
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package gistpost

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
//  	Public	bool	`short:"p" long:"pub" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
//  }

//
//  var createCommand CreateCommand
//
//  ////////////////////////////////////////////////////////////////////////////
//  // Function definitions
//
//  func init() {
//  	gfParser.AddCommand("create",
//  		"Create a new GH gist entry (file)",
//  		"Usage:\n  gistpost [Options] create [-p]",
//  		&createCommand)
//  }
//
//  func (x *CreateCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Create a new GH gist entry (file)\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
//   	clis.Setup("gistpost::create", Opts.Verbose)
//   	clis.Verbose(1, "Doing Create, with %+v, %+v", Opts, args)
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

// Template for "update" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package gistpost

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: update ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The UpdateCommand type defines all the configurable options from cli.
//  type UpdateCommand struct {
//  	GistID	string	`short:"g" long:"id" env:"GISTPOST_GISTID" description:"Existing GH gist id*" required:"true"`
//  }

//
//  var updateCommand UpdateCommand
//
//  ////////////////////////////////////////////////////////////////////////////
//  // Function definitions
//
//  func init() {
//  	gfParser.AddCommand("update",
//  		"Update an existing GH gist entry (file)",
//  		"Usage:\n  gistpost [Options] update --id",
//  		&updateCommand)
//  }
//
//  func (x *UpdateCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Update an existing GH gist entry (file)\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
//   	clis.Setup("gistpost::update", Opts.Verbose)
//   	clis.Verbose(1, "Doing Update, with %+v, %+v", Opts, args)
//   	// fmt.Println(x.GistID)
//  	return x.Exec(args)
//  }
//
// // Exec implements the business logic of command `update`
// func (x *UpdateCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("update::Exec", err)
// 	// or,
// 	// clis.AbortOn("update::Exec", err)
// 	return nil
// }
// Template for "update" CLI handling ends here

// Template for "folder" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package gistpost

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: folder ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The FolderCommand type defines all the configurable options from cli.
//  type FolderCommand struct {
//  	Dir	string	`short:"D" long:"dir" description:"Directory to upload as gist*" required:"true"`
//  	Public	bool	`short:"p" long:"pub" env:"GISTPOST_PUBLIC" description:"Public gist or not"`
//  	Extra	bool	`short:"e" long:"extra" env:"GISTPOST_EXTRA" description:"Extra files will be added to gist for better name/doc."`
//  }

//
//  var folderCommand FolderCommand
//
//  ////////////////////////////////////////////////////////////////////////////
//  // Function definitions
//
//  func init() {
//  	gfParser.AddCommand("folder",
//  		"Upload the whole folder as GH gist",
//  		"Usage:\n  gistpost [Options] folder [-p] [-e]",
//  		&folderCommand)
//  }
//
//  func (x *FolderCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Upload the whole folder as GH gist\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
//   	clis.Setup("gistpost::folder", Opts.Verbose)
//   	clis.Verbose(1, "Doing Folder, with %+v, %+v", Opts, args)
//   	// fmt.Println(x.Dir, x.Public, x.Extra)
//  	return x.Exec(args)
//  }
//
// // Exec implements the business logic of command `folder`
// func (x *FolderCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("folder::Exec", err)
// 	// or,
// 	// clis.AbortOn("folder::Exec", err)
// 	return nil
// }
// Template for "folder" CLI handling ends here

// Template for "to-raw" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package gistpost

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: to-raw ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The ToRawCommand type defines all the configurable options from cli.
//  type ToRawCommand struct {
//  }

//
//  var to_rawCommand ToRawCommand
//
//  ////////////////////////////////////////////////////////////////////////////
//  // Function definitions
//
//  func init() {
//  	gfParser.AddCommand("to-raw",
//  		"Get raw content url from GH gist",
//  		"Usage:\n  gistpost [Options] to-raw gist-url",
//  		&to_rawCommand)
//  }
//
//  func (x *ToRawCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Get raw content url from GH gist\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
//   	clis.Setup("gistpost::to-raw", Opts.Verbose)
//   	clis.Verbose(1, "Doing ToRaw, with %+v, %+v", Opts, args)
//   	// fmt.Println()
//  	return x.Exec(args)
//  }
//
// // Exec implements the business logic of command `to-raw`
// func (x *ToRawCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("to-raw::Exec", err)
// 	// or,
// 	// clis.AbortOn("to-raw::Exec", err)
// 	return nil
// }
// Template for "to-raw" CLI handling ends here

// Template for "append" CLI handling starts here
////////////////////////////////////////////////////////////////////////////
// Program: gistpost
// Purpose: GH gist post/update tool
// Authors: Chenhua Sun (c) 2024-2026, All rights reserved
////////////////////////////////////////////////////////////////////////////

//  package gistpost

//  import (
//  	"fmt"
//  	"os"
//
//  	"github.com/go-easygen/go-flags/clis"
//  )

// *** Sub-command: append ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The AppendCommand type defines all the configurable options from cli.
//  type AppendCommand struct {
//  	GistID	string	`short:"g" long:"id" env:"GISTPOST_GISTID" description:"Existing GH gist id*" required:"true"`
//  	Dir	string	`short:"D" long:"dir" description:"Directory whose files to append*" required:"true"`
//  	Extra	bool	`short:"e" long:"extra" env:"GISTPOST_EXTRA" description:"Extra files will be added to gist for better name/doc."`
//  }

//
//  var appendCommand AppendCommand
//
//  ////////////////////////////////////////////////////////////////////////////
//  // Function definitions
//
//  func init() {
//  	gfParser.AddCommand("append",
//  		"Append all given folder files to an existing GH gist entry",
//  		"Usage:\n  gistpost [Options] append -g <gid> -D <dir> [-e]",
//  		&appendCommand)
//  }
//
//  func (x *AppendCommand) Execute(args []string) error {
//   	fmt.Fprintf(os.Stderr, "Append all given folder files to an existing GH gist entry\n")
//   	// fmt.Fprintf(os.Stderr, "Copyright (C) 2024-2026, Chenhua Sun\n\n")
//   	clis.Setup("gistpost::append", Opts.Verbose)
//   	clis.Verbose(1, "Doing Append, with %+v, %+v", Opts, args)
//   	// fmt.Println(x.GistID, x.Dir, x.Extra)
//  	return x.Exec(args)
//  }
//
// // Exec implements the business logic of command `append`
// func (x *AppendCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("append::Exec", err)
// 	// or,
// 	// clis.AbortOn("append::Exec", err)
// 	return nil
// }
// Template for "append" CLI handling ends here
