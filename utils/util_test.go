package utils

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestFQDN(t *testing.T) {
	libfqdn := FQDN()

	cmd := exec.Command("/bin/hostname", "-f")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fqdn := out.String()
	fqdn = fqdn[:len(fqdn)-1]

	if libfqdn != fqdn {
		t.Fail()
	}
}

func TestInStringSlick(t *testing.T) {
	s := []string{"b", "d", "f"}

	if !InStringSlice(s, "b") {
		t.Fatal("b")
	}
	if !InStringSlice(s, "d") {
		t.Fatal("d")
	}
	if !InStringSlice(s, "f") {
		t.Fatal("f")
	}
	if InStringSlice(s, "a") {
		t.Fatal("a")
	}
	if InStringSlice(s, "c") {
		t.Fatal("c")
	}
	if InStringSlice(s, "e") {
		t.Fatal("e")
	}
	if InStringSlice(s, "g") {
		t.Fatal("g")
	}
}
