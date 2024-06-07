/*
package reflect/helper 从包 reflect 的 Value 和 Type 分离出不同类型的反射接口转发。

TypeTo 从 [reflect.Type] 尝试构建指定类型的 TypeHelper，该反射接口转发仅包含适用于特定类型的反射方法集

	if t, err := TypeTo[*SliceType](tp); err == nil {
	  do slice by t
	}

TypeFor 从 T 的类型构造一个 TypeHelper

	t := TypeFor[T]()
	if t.Kind() == reflect.Slice {
	  st := t.(*SliceType)
	  do slice by t
	}

TypeOf 从 v 中提取反射信息，并包装为其类型特定的反射接口转发。

	  t := TypeOf(i)
		if t.Kind() == reflect.Slice {
		  st := t.(*SliceType)
		  do slice by t
		}

TypeWrap 包装一个 [reflect.Type]，根据 Kind 可以显式转换为特定的反射接口转发。
*/
package helper
