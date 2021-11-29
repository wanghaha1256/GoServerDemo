// This package contains universal function
// such as operate file, etc.
package pkg

import (
	"io/ioutil"
	"log"
	"strings"
)

// GetFileListWithSuffix return files' name list with specific suffix
func GetFileListWithSuffix(root string, suffix string) []string {
	res := []string{}
	files, err := ioutil.ReadDir(root)

	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			res = append(res, f.Name())
		}
	}
	return res
}

func ReadFile(filePath string) []byte {

	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalln(err)
	}
	return fileContent
}
