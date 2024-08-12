// !!! !!!
// WARNING: Code automatically generated. Editing discouraged.
// !!! !!!

package main

import (
	"flag"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "gistpost" // os.Args[0]

// The Options struct defines the structure to hold the commandline values
type Options struct {
	Token       string // The GITHUB_TOKEN*
	Description string // Gist description
	Filename    string // Gist filename
	Public      bool   // Public gist or not
	Help        bool   // show usage help
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line parameters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func initVars() {

	// set default values for command line parameters
	flag.StringVar(&Opts.Token, "t", "",
		"The GITHUB_TOKEN*")
	flag.StringVar(&Opts.Description, "d", "",
		"Gist description")
	flag.StringVar(&Opts.Filename, "f", "archive.md",
		"Gist filename")
	flag.BoolVar(&Opts.Public, "p", false,
		"Public gist or not")
	flag.BoolVar(&Opts.Help, "h", false,
		"show usage help")
}

func initVals() {
	exists := false
	// Now override those default values from environment variables
	if len(Opts.Token) == 0 &&
		len(os.Getenv("GP_GIST_T_TOKEN")) != 0 {
		Opts.Token = os.Getenv("GP_GIST_T_TOKEN")
	}
	if len(Opts.Description) == 0 &&
		len(os.Getenv("GP_GIST_D_DESCRIPTION")) != 0 {
		Opts.Description = os.Getenv("GP_GIST_D_DESCRIPTION")
	}
	if len(Opts.Filename) == 0 &&
		len(os.Getenv("GP_GIST_F_FILENAME")) != 0 {
		Opts.Filename = os.Getenv("GP_GIST_F_FILENAME")
	}
	if _, exists = os.LookupEnv("GP_GIST_P_PUBLIC"); Opts.Public || exists {
		Opts.Public = true
	}
	if _, exists = os.LookupEnv("GP_GIST_H_HELP"); Opts.Help || exists {
		Opts.Help = true
	}

}

const usageSummary = "  -t\tThe GITHUB_TOKEN* (GP_GIST_T_TOKEN)\n  -d\tGist description (GP_GIST_D_DESCRIPTION)\n  -f\tGist filename (GP_GIST_F_FILENAME)\n  -p\tPublic gist or not (GP_GIST_P_PUBLIC)\n  -h\tshow usage help (GP_GIST_H_HELP)\n\nDetails:\n\n"

// Usage function shows help on commandline usage
func Usage(exit int) {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags] \n\nFlags:\n\n",
		progname)
	fmt.Fprintf(os.Stderr, usageSummary)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		``)
	if exit != 0 {
		os.Exit(0)
	}
}
