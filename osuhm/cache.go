package osuhm

import (
	"encoding/gob"
	"os"
	"syscall"
)

// Encode via Gob to file
func save(path string, object interface{}) error {
	nameptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		file.Close()
		return err
	}

	// Change the file into a hidden file
	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		os.Remove(path) // XXX do we want to remove it? check for error
		file.Close()    // XXX check error
		return err
	}

	encoder := gob.NewEncoder(file)
	encoder.Encode(object)
	return nil

}

// Decode Gob file
func load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
