package gostd

import (
	"container/ring"
	"math"
	"testing"
)

/*
! New 创建一个 n 元素的环
! ring.Ring 是一个循环列表或环的元素。环没有开始或结束；指向任何环元素的指针用作整个环的引用。空环表示为 nil 环指针。环的零值是一个具有 nil 值的单元素环。
	Do 对环中的每个元素按前向顺序调用函数 f。如果 f 改变 *r，则 Do 的行为未定义
	Len 计算环 r 中元素的个数
	Link 连接环 r 和环 s，使得 r.Next 变成 s，并返回 r.Next() 的原始值。r 不能为空。
		如果 r 和 s 指向同一个环，连接它们会从环中删除 r 和 s 之间的元素。移除的元素形成一个子环，
		结果是对该子环的引用（如果没有元素被移除，结果仍然是 r.Next() 的原始值，而不是 nil）。
	Move 在环中向后 (n < 0) 或向前 (n >= 0) 移动 n % r.Len() 元素，并返回该环元素
	Next 返回下一个环元素
	Prev 返回前一个环元素
	Unlink 从环 r 中删除 n % r.Len() 元素，从 r.Next() 开始。如果 n % r.Len() == 0，则 r 保持不变。结果是被移除的子环。
*/

func TestRing(t *testing.T) {
	var r = ring.New(10)
	var tmp *ring.Ring = r
	for i := range r.Len() {
		tmp.Value = i
		tmp = tmp.Next()
	}
	sqrt := func(a any) {
		logfln("sqrt of %v is %.4f", a, math.Sqrt(float64(a.(int))))
	}

	log := func(a any) {
		logfln("%v", a)
	}

	r.Do(sqrt)

	r2 := r.Unlink(5)
	for i := range r2.Len() {
		r2.Value = r2.Value.(int)*(i+1)*i + i
		r2 = r2.Prev()
	}
	r2.Do(log)

	r.Link(r2)
	r.Do(sqrt)
	r = r.Move(5)
	r.Link(r2)
	r.Do(log)
}
