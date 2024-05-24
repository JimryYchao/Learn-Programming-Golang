package gostd

import (
	"fmt"
	"testing"
)

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
	logf("LOG ERROR: \n%s", err)
}
func logCase(_case string) {
	logf("case : %s", _case)
}
func logf(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}
