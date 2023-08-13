package store

import (
	"os"
	"path/filepath"
)

func Scrape(root *string) ([]string, []string) {
	filenames := []string{}
	filedirs := []string{}
	filepath.Walk(*root, func(path string, info os.FileInfo, err error) error {
		extension := filepath.Ext(info.Name())
		if extension != "" {
			name := info.Name()
			// name, _ := strings.TrimSuffix(info.Name(), extension), info.Size()
			filenames = append(filenames, name)
			filedirs = append(filedirs, path)
		}
		return nil
	})
	return filenames, filedirs
}
