package helper_test

import (
	. "gostd/reflect/helper"
	"testing"
)

func Test(t *testing.T) {
	log(MethodOf(TypeFor[ArrayType]()).NumMethod())
	iterMethods(MethodOf(TypeFor[ArrayType]()))
}

func testMethodSet(set *MethodSet) {
	iterMethods(set)
}

func iterMethods(set *MethodSet) {
	for _, m := range set.Methods() {
		logf("name:%s, type:%s", m.Name, m.Type)
	}
}
