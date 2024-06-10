package gostd

import (
	"fmt"
	"testing"
)

func _logCase(_case string) {
	_logfln("case : %s", _case)
}

var fileNames = []string{
	"e.txt",
	"compress.pdf",
	"Isaac.Newton-Opticks.txt",
}

var EnterTest = ">>> Enter %s :\n"
var EndTest = ">>> End   %s\n"

func beforeTest[TBF testing.TB](t TBF) {
	if !testing.Verbose() {
		return
	}
	fmt.Printf(EnterTest, t.Name())
	t.Cleanup(func() {
		fmt.Printf(EndTest, t.Name())
	})
}
func checkErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("LOG ERROR: %s\n", err)
}

func _log(a ...any) {
	s := fmt.Sprintln(a...)
	if s[len(s)-1] != '\n' {
		s += "\n"
	}
	fmt.Print(s)
}
func _logfln(format string, args ...any) {
	s := fmt.Sprintf(format, args...)
	if s[len(s)-1] != '\n' {
		s += "\n"
	}
	fmt.Print(s)
}
