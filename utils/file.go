package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

type CopyConvertFn func(in []byte) (out []byte, err error)

func Tree(dir string) (o []byte, err error) {
	cmd := exec.Command("tree", ".", "-a", "-p", "-s", "-D", "-L", "2")
	cmd.Dir = dir
	o, err = cmd.CombinedOutput()
	return
}

func Md5sum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5h := md5.New()
	if _, err := io.Copy(md5h, f); err != nil {
		return "", err
	}

	md5b := md5h.Sum([]byte(""))
	return hex.EncodeToString(md5b[:]), nil
}

func IsSameFileContent(f1, f2 string) (bool, error) {
	md51, err := Md5sum(f1)
	if err != nil {
		return false, err
	}

	md52, err := Md5sum(f2)
	if err != nil {
		return false, err
	}

	return md51 == md52, nil
}

func IsFileExist(fpath string) bool {
	if _, err := os.Stat(fpath); os.IsNotExist(err) {
		return false
	}
	return true
}

func Chown(f, usr string) (err error) {
	u, err := user.Lookup(usr)
	if err != nil {
		return
	}

	uid, e1 := strconv.Atoi(u.Uid)
	gid, e2 := strconv.Atoi(u.Gid)
	if e1 != nil || e2 != nil {
		return errors.New(fmt.Sprintf("Convert uid/gid to int error! uid: %v, gid: %v", u.Uid, u.Gid))
	}

	return os.Chown(f, uid, gid)
}

func CreateParentDir(f string) error {
	return os.MkdirAll(path.Dir(f), 0755)
}

func SafeCreate(f string) (file *os.File, err error) {
	err = CreateParentDir(f)
	if err != nil {
		return
	}

	if _, err = os.Stat(f); err == nil {
		os.Rename(f, f+".old")
	}

	file, err = os.Create(f)
	return
}

func Copy(src, dest string) error {
	fr, _ := os.Open(src)
	defer fr.Close()
	fw, _ := SafeCreate(dest)
	defer fw.Close()

	_, err := ioutil.ReadAll(io.TeeReader(fr, fw))
	return err
}

func MoreCopy(src, dest string, convertFn CopyConvertFn) error {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	converted, err := convertFn(content)
	if err != nil {
		return err
	}

	destHandle, err := SafeCreate(dest)
	if err != nil {
		return err
	}
	defer destHandle.Close()

	_, err = destHandle.Write(converted)
	return err
}

// recursiveFindFiles find files with pattern in the root with depth.
func RecursiveFindFiles(root string, pattern string) ([]string, error) {
	files := make([]string, 0)
	findfile := func(path string, f os.FileInfo, err error) (inner error) {
		if err != nil {
			return
		}
		if f.IsDir() {
			return
		} else if match, innerr := filepath.Match(pattern, f.Name()); innerr == nil && match {
			files = append(files, path)
		}
		return
	}
	err := filepath.Walk(root, findfile)
	if len(files) == 0 {
		return files, err
	} else {
		return files, err
	}
}

// return retry times and error
func DownloadFromUrl(url, localFile string, overWrite bool) (int, error) {
	if !overWrite && IsFileExist(localFile) {
		return 0, errors.New("local file " + localFile + " exist and overWrite set to false")
	}

	return Retry(func() error {
		output, err := os.Create(localFile)
		if err != nil {
			return errors.New("create localfile failed: " + err.Error())
		}
		defer output.Close()

		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("error while downloading %v: %v", url, err.Error())
		}
		defer resp.Body.Close()

		_, err = io.Copy(output, resp.Body)
		if err != nil {
			return fmt.Errorf("error while downloading %v: %v", url, err.Error())
		}
		return nil
	}, 3, time.Duration(500)*time.Millisecond)

}

func DownloadViaWget(url, localFile string, overWrite bool) error {
	if !overWrite && IsFileExist(localFile) {
		return errors.New("local file " + localFile + " exist and overWrite set to false")
	}

	cmd := exec.Command("wget", "-c", "--waitretry=64", "-O", localFile, url)
	o, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error while downloading %v: %v,%v", url, err.Error(), o)
	}

	return nil
}
