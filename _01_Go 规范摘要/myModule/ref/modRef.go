package modRef

import "rsc.io/quote"

func init() {
	println("init modRef")

	quote.Hello()

}

func Export() {
	println("Exec func")

}

// 工作区的 go 行必须声明一个大于或等于 use 语句中列出的每个模块所声明的 go 版本的版本。
