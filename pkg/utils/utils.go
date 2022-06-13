package utils

import "os"

func FileOrFirExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
