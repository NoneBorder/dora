package dora

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestTextWriter(t *testing.T) {
	Info().Msg("log before SetToTextWriter")
	Logger = TextWriter(nil)
	Info().Msg("log after SetToTextWriter")
}

func removeTestFile(fileWildcard string, t *testing.T) {
	files, err := filepath.Glob(fileWildcard)
	if err != nil {
		t.Fatalf("find file failed for file wildcard " + fileWildcard)
	}

	for _, f := range files {
		if err := os.Remove(f); err != nil {
			t.Error(fmt.Sprintf("remote file failed, file=%s error=%s", f, err.Error()))
		}
	}
}

func TestFileWriter(t *testing.T) {
	randNo := strconv.Itoa(rand.Intn(1000))
	logFileName := "test." + randNo + ".log"
	fileWildcard := "test." + randNo + "*.log"
	c := `{"filename":"` + logFileName + `","maxLines":3,"maxsize":1024,"daily":true,"maxDays":15,"rotate":true,"perm":"0666","rotateperm":"0666"}`

	removeTestFile(fileWildcard, t)

	Logger = NewLogWithWriter("file", c, "debug")
	for i := 0; i < 5; i++ {
		Info().Int("index", i).Msg("log for file write test")
	}

	files, err := filepath.Glob(fileWildcard)
	if err != nil {
		t.Fatalf("find file failed for file wildcard " + fileWildcard)
	}

	if len(files) != 2 {
		t.Fatalf("should have to log files")
	}

	removeTestFile(fileWildcard, t)
}
