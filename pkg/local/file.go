package local

import (
	"io/ioutil"
	"os"
	"io"
	"strings"
)

func SPutContent(data, path string) error {
	fo, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fo.Close()
	_, err = io.Copy(fo, strings.NewReader(data))
	if err != nil {
		return err
	}
	return nil
}

func SGetContent(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		//RecoverableErrorf("reading %s: %v", path, err)
		return ""
	}
	return string(b)
}

func BGetContent(path string) []byte {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		//RecoverableErrorf("reading %s: %v", path, err)
		return []byte("")
	}
	return b
}
