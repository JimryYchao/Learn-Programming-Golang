package gostd

import (
	"fmt"
	"testing"
)

func logCase(_case string) {
	logfln("case : %s", _case)
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
	fmt.Printf("LOG ERROR: \n%s", err)
}

func log(s any) {
	fmt.Println(s)
}
func logfln(format string, args ...any) {
	fmt.Printf(format+"\n", args...)

}
