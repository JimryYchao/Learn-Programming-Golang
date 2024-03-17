package GoLinq

type Func[TResult, T1, T2 any] func(t1 T1, t2 T2) TResult

type Func1 func(t T) TResult

type Func2 func(t1 T1, t2 T2) TResult
