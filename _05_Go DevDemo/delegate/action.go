package delegate

type ActionFunc *func(...any)

type (
	ActionFunc_0             func()
	ActionFunc_1[T any]      func(t T)
	ActionFunc_2[T1, T2 any] func(t1 T1, t2 T2)
)

type BaseInvoke interface {
	Invoke(...any)
}

var Action AcBase

type AcBase struct {
	BaseInvoke
}
type Ac_0 struct {
	AcBase
	ac ActionFunc
}

type Ac_1[T any] struct {
	AcBase
	acs   *[]*ActionFunc_1[T]
	index int
}

func (ac AcBase) Init() {

}
func (ac *Ac_1[T]) Invoke(t T) {
	for _, f := range *ac.acs {
		(*f)(t)
	}
}

func (ac *Ac_1[T]) Init(fun *ActionFunc_1[T]) {
	if fun == nil {
		return
	}
	s := make([]*ActionFunc_1[T], 1)
	ac.acs = &s
	(*ac.acs)[0] = fun
	ac.index = 1
}
func (ac *Ac_1[T]) Add(fun *ActionFunc_1[T], funs ...*ActionFunc_1[T]) *Ac_1[T] {
	*ac.acs = append(*ac.acs, fun)
	if len(funs) > 0 {
		*ac.acs = append(*ac.acs, funs...)
	}
	ac.index += 1 + len(funs)
	return ac
}

func (ac *Ac_1[T]) Remove(fun *ActionFunc_1[T]) {
	if fun == nil {
		return
	}
	for i, f := range *ac.acs {
		if f == fun {
			copy((*ac.acs)[i:], (*ac.acs)[i+1:ac.index])
			ac.index--
			*ac.acs = (*ac.acs)[0:ac.index]
			return
		}
	}
}

type GAction[T any] struct {
}

func (a AcBase) Invoke(args ...any) {
}

func FFFF[T any](ac ActionFunc_1[T], arg T) {
	// ac(&arg)
}

func main() {

	var ac Ac_1[int]
	var f1, f2, f3 ActionFunc_1[int] = func(int) { println(1) }, func(int) { println(2) }, func(int) { println(3) }
	ac.Init(&f1)
	ac.Add(&f2, &f3)
	ac.Invoke(10)

	ac.Remove(&f2)
	ac.Invoke(1)
}
