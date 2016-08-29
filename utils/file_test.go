package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testFileName    = "utils.test"
	testFileContent = "utils.test"
	testFileMD5     = "1b4b330c4cb02b20bd6bab694920406b"
)

func TestFileMd5sum(t *testing.T) {
	if err := ioutil.WriteFile(testFileName, []byte(testFileContent), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFileName)

	md5, err := Md5sum(testFileName)
	if err != nil {
		t.Fatal(err)
	}

	if md5 != testFileMD5 {
		t.Fail()
	}
}

func TestTree(t *testing.T) {
	o, err := Tree(".")
	fmt.Printf("%s %v", o, err)
}
