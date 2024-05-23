package gostd

import (
	"fmt"
	"strings"
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
	fmt.Printf("LOG ERROR: \n%s", err)
}

func log(s any) {
	fmt.Println(s)
}
func logfln(format string, args ...any) {
	fmt.Printf(format+"\n", args...)

}

func checkVerbs(v any, verbs ...string) {
	if v == nil {
		return
	}
	var rt []byte
	if _, ok := v.(string); ok {
		rt = []byte(fmt.Sprintf("(T:%T, V:%s)", v, fmt.Sprintf("%q", v)))
	} else {
		rt = []byte(fmt.Sprintf("(T:%T, V:%v)", v, v))
	}

	if len(verbs) == 0 {
		return
	}
	fmts, okCh := make([][]byte, len(verbs)), make(chan bool)

	for i, verb := range verbs {
		i, verb := i, verb
		go func() {
			fmts[i] = []byte(fmt.Sprintf("%-10s", verb) + fmt.Sprintf("%#q", fmt.Sprintf(verb, v)))
			okCh <- true
		}()
	}

	var c = 0
	for {
		<-okCh
		c++
		if c >= len(verbs) {
			break
		}
	}
	var sb strings.Builder
	sb.Write(rt)
	for _, format := range fmts {
		sb.WriteString("\n    " + string(format))
	}
	logfln("%s\n", sb.String())
}

func TestFormat(t *testing.T) {
	beforeTest(t)
	// "a", []byte("a"), [1]byte{'a'}, &[1]byte{'a'} 在字符串格式化上等效
	checkVerbs("abcxyz", "%s", "%q", "%x", "%X", "% x", "% X", "%#x", "%#X", "%# x", "% #X")

	checkVerbs("`az我☺⌘\xff\U0010ffff\a\b\f\n\r\t\v\\\""+`\n`,
		"%x", "%X", "% x", "% X", "%#x", "%#X", "%# x", "% #X", "%s", "%q")
}
