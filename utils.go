package main

import (
	"os"
)

func exists(path string) bool {
	var _, err = os.Stat(path)
	return err == nil || os.IsExist(err)
}
