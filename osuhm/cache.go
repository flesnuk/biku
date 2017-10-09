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

	os.Remove(path)
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	// Change the file into a hidden file
	err = syscall.SetFileAttributes(nameptr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		os.Remove(path)
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
