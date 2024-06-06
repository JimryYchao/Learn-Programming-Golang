package GoLinq

func (source Enumerable) Aggregate(fun Func[TSource, TSource, TSource]) {
	if source.IEnumerable == nil {
		panic("source is null")
	}

	if fun == nil {
		panic("fun is null")
	}
	var e = source.IEnumerable.GetEnumerator()
	defer e.Dispose()

	if !e.MoveNext() {
		panic("no elements")
	}

}
