package utils

import (
	"path/filepath"
	"os"
	"fmt"
)

func GetBinDirectory() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Errorf("Can't get current directory: %s \n", err))
	}

	return path
}
