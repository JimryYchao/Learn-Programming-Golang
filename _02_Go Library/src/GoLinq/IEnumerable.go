package GoLinq

type T interface {
	any
}

type TSource interface {
	any
}

type TResult interface {
	any
}

type T2 interface {
	any
}

type T1 interface {
	any
}

type Enumerable struct {
	IEnumerable
}

type IEnumerable interface {
	GetEnumerator() IEnumerator
}

type IEnumerator interface {
	Current() any
	MoveNext() bool
	Dispose()
}
