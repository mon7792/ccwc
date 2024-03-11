package flg

import (
	"flag"
	"log"
	"os"
)

// OpsType is the type of operation to be performed
type OpsType int

const (
	CharCount OpsType = iota
	LineCount
	WordCount
	MultiCharCount
	Def
)

// DataInput is the type of input data
type DataInput int

const (
	FsIn DataInput = iota
	StdIn
)

// defWCUsage is the default usage for the wc command
const defWCUsage = `
USAGE ccwc

NAME
     ccwc – word, line, character, and byte count

SYNOPSIS
     ccwc [-clmw] [file ...]

DESCRIPTION
     The wc utility displays the number of lines, words, and bytes contained in each input file, or standard input (if no file is specified) to
     the standard output.  A line is defined as a string of characters delimited by a ⟨newline⟩ character.  Characters beyond the final
     ⟨newline⟩ character will not be included in the line count.

     A word is defined as a string of characters delimited by white space characters.  White space characters are the set of characters for
     which the iswspace(3) function returns true.  If more than one input file is specified, a line of cumulative counts for all the files is
     displayed on a separate line after the output for the last file.

     The following options are available:
	 -c      The number of bytes in each input file is written to the standard output.  This will cancel out any prior usage of the -m option.

     -l      The number of lines in each input file is written to the standard output.

     -m      The number of characters in each input file is written to the standard output.  If the current locale does not support multibyte
             characters, this is equivalent to the -c option.  This will cancel out any prior usage of the -c option.

     -w      The number of words in each input file is written to the standard output.

EXAMPLES
	 1. ccwc -c file
	 will output the number of bytes in the file.
	 2. ccwc -l file
	 will output the number of lines in the file.
	 3. ccwc -w file
	 will output the number of words in the file.
	 4. ccwc -m file
	 will output the number of multi characters in the file.
`

// Ops is the struct for the operations
type Ops struct {
	flgType   OpsType
	dataInput DataInput
	fsPath    string
}

// New returns a new Operations struct
func New() *Ops {
	return &Ops{}
}

// Init initializes the flags
func (f *Ops) Init() {
	flag.Bool("c", false, "count the number of bytes")
	flag.Bool("l", false, "count the number of lines")
	flag.Bool("w", false, "count the number of words")
	flag.Bool("m", false, "count the number of bytes for multibyte characters")
}

// Parse parses the flags
func (f *Ops) Parse() {
	flag.Usage = func() {
		log.Print(defWCUsage)
	}
	flag.Parse()
}

// Verify verifies the flags
func (f *Ops) Verify() {
	// check if the flags are proper
	// at max only one flag can be set
	if flag.NFlag() > 1 {
		log.Println("error: More than one flag is used")
		flag.Usage()
		os.Exit(1)
	}

	// check if the file args are proper
	if len(flag.Args()) > 1 {
		log.Println("error: No file is specified")
		flag.Usage()
		os.Exit(1)
	}
}

// Set sets the flags
func (f *Ops) Set() {
	flag.Visit(func(fl *flag.Flag) {
		switch fl.Name {
		case "c":
			f.flgType = CharCount
		case "l":
			f.flgType = LineCount
		case "w":
			f.flgType = WordCount
		case "m":
			f.flgType = MultiCharCount
		default:
			f.flgType = Def
		}
	})

	if len(flag.Args()) == 1 {
		f.dataInput = FsIn
		f.fsPath = flag.Args()[0]
	} else {
		f.dataInput = StdIn
	}
}

// Get returns the ops type and the data input
func (f *Ops) Get() (OpsType, DataInput, string) {
	return f.flgType, f.dataInput, f.fsPath
}
