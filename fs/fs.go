package fs

import (
	"errors"
	"io"
	"log"
	"os"
)

// PathExist checks if the file exists
func PathExist(fsPath string) bool {
	_, err := os.Stat(fsPath)
	return !errors.Is(err, os.ErrNotExist)
}

// ReadFs reads the file and sends the data to the channel
func ReadFs(fsPath string, outCh chan<- []byte) (err error) {
	// open the file
	var fs *os.File
	fs, err = os.Open(fsPath)
	if err != nil {
		log.Println("error: unable to open the file ", err)
		return err
	}
	defer func() {
		err = fs.Close()
	}()

	// read the file
	var bufSize = 4
	for {
		var buf = make([]byte, bufSize)
		_, err = fs.Read(buf)
		if err == io.EOF {
			close(outCh)
			break
		}
		if err != nil {
			log.Println("error: reading file : ", err)
			return err
		}

		// perform the operation here.
		outCh <- buf
	}
	// read the file
	return nil

}
