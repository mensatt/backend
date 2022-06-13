package utils

import "os"

func FileOrFirExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func CreateDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}