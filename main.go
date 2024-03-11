package main

import (
	"log"

	"github.com/mon7792/ccwc/flg"
	"github.com/mon7792/ccwc/fs"
	"github.com/mon7792/ccwc/wc"
)

func main() {

	log.Println("")

	// parse the flags
	fg := flg.New()
	fg.Init()
	fg.Parse()
	fg.Verify()
	fg.Set()
	opsType, dataIn, fsPath := fg.Get()

	if dataIn == flg.StdIn {
		log.Println("Stdin")
	} else {
		// verify if the file exists
		if !fs.PathExist(fsPath) {
			log.Println("error: file does not exist")
			return
		}

		// read the file
		var ch = make(chan []byte)

		go fs.ReadFs(fsPath, ch)

		// perform the operation
		switch opsType {
		case flg.CharCount:
			res := wc.CharacterCount(ch)
			log.Println("Character count: ", res)
		case flg.LineCount:
			res := wc.LineCount(ch)
			log.Println("Line count: ", res)
		case flg.WordCount:
			res := wc.WordCount(ch)
			log.Println("Word count: ", res)
		default:
			log.Println("error: invalid operation")

		}
	}

}
