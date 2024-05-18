package gostd_testing

import (
	"flag"
	"math/rand"
	"os"
	"runtime/debug"
	"testing"
	"time"
)

// ! TestMain 为测试主入口 (可选)
func TestMain(m *testing.M) {
	//? go test [-v]
	//? Run 运行测试，返回一个可以传递给 os.Exit 的退出代码
	exitCode := m.Run()

	os.Exit(exitCode)
}

// ! TestXxx(t *testing.T) 测试
func TestXxx(t *testing.T) {
	//? go test -v -run=Test
	t.Log("Testing `TestXxx`")
}

// ! BenchmarkXxx(b *testing.B) 基准测试
func BenchmarkXxx(b *testing.B) {
	//? go test -run=NONE -bench=Benchmark [-benchtime=3s] [-benchmem]
	//? 迭代足够次数或运行足够时长来计算单次操作 body 使用的近似时间
	for i := 0; i < b.N; i++ {
		// body
		doSomething(100)
	}
}

// ! FuzzXxx(f *testing.F) 模糊测试
func FuzzXxx(f *testing.F) {
	//? go test -v -run=NONE -fuzz=^FuzzXxx$ [-parallel=8] [-fuzztime=5s] [-short]
	mlimit := debug.SetMemoryLimit(1000 * 100)
	mTreads := debug.SetMaxThreads(1000)
	f.Cleanup(func() {
		debug.SetMemoryLimit(mlimit)
		debug.SetMaxThreads(mTreads)
	})

	f.Log("Fuzzing `FuzzXxx`")
	added := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, v := range added {
		//? Add 添加种子用料
		f.Add(v)
	}
	//? Fuzz 运行模糊目标函数 ff
	f.Fuzz(func(t *testing.T, up int) {
		if up < 2 {
			t.SkipNow()
			return
		}
		// else i > 0, do something
		doSomething(up)
	})
}

// ! AllocsPerRun 函数返回调用 `f` 期间分配的平均次数 (allocs/op)。
func TestAllocsPerRun(t *testing.T) {
	//? go test -v -run=^TestAllocsPerRun$
	var allocsPerRunTests = []struct {
		name   string
		fn     func(temp *any)
		allocs float64
	}{
		{"alloc *byte", func(temp *any) { *temp = new(*byte) }, 1},
		{"alloc complex128", func(temp *any) { *temp = new(complex128) }, 1},
		{"alloc float64", func(temp *any) { *temp = new(float64) }, 1},
		{"alloc int32", func(temp *any) { *temp = new(int32) }, 1},
		{"alloc byte", func(temp *any) { *temp = new(byte) }, 1},
	}
	var temp any
	for _, tt := range allocsPerRunTests {
		if allocs := testing.AllocsPerRun(100, func() { tt.fn(&temp) }); allocs != tt.allocs {
			t.Errorf("AllocsPerRun(100, %s) = %v, want %v.", tt.name, allocs, tt.allocs)
		}
	}
}

// ! CoverMode 报告在测试的覆盖模式: set, count or atomic。
// ! Coverage 报告当前代码的覆盖率, 它不能替代 `go test -cover` 和 `go tool cover` 生成的报告。
func TestCoverage(t *testing.T) {
	//? go test -v -cover -run=^TestCoverage$ [-coverprofile='coverage.out']
	//? go tool cover -html='coverage.out'
	t.Cleanup(func() {
		// 标志 -cover 启用代码覆盖率分析, 默认为 set 模式
		coverMode := testing.CoverMode() // 未设置时返回 ""
		if coverMode == "" {
			t.Log("-cover is not enabled.")
			return
		}
		t.Logf("Cover Mode: %s, Coverage in %s is %v.", coverMode, t.Name(), testing.Coverage())
	})
	// call another file
	Perm(-100)
}

// ! Short 报告在测试中是否启用 `-test.short`。
// ! Verbose 报告在测试中是否启用 `-test.v`。
// ! Testing 报告当前代码是否在测试中运行。
func TestTesting(t *testing.T) {
	//? go test -v -short -run=^TestTesting$
	if !testing.Testing() {
		t.Errorf("%s runs only in programs created by `go test`.", t.Name())
	}
	if testing.Verbose() {
		t.Log("Enable -v mode.")
	}
	if !testing.Short() {
		t.Skipf("Skipping %s in non-short mode.", t.Name())
	}
}

// ! TB 是 `T`、`B` 和 `F` 的公共接口
func TestTBFunctions(t *testing.T) {
	//? go test -v -run=^TestTB$
	t.Helper()
	beforeTest := func(t *testing.T) {
		t.Logf("Testing >>> %s", t.Name())
		// 测试结束时按注册逆序依次调用
		t.Cleanup(func() {
			t.Logf("End Test >>> %s", t.Name())
		})
	}
	t.Run("Fail", func(t *testing.T) {
		beforeTest(t)
		t.Cleanup(func() {
			if t.Skipped() {
				t.Log("Please use command: `go test -v -run=^TestTB$`.")
			}
		})
		var _sub *testing.T
		if runflag := flag.Lookup("test.run"); runflag != nil && runflag.Value.String() == "^TestTB$" {
			// 仅运行当前测试时故意失败
			t.Run("-run=^TestTB$", func(t *testing.T) {
				_sub = t
				t.Fatal("TestTB/Fail fails deliberately.")
			})
		} else {
			t.Skip("TestTB/Fail runs only in `-run=^TestTB$` mode.")
		}
		if _sub != nil && _sub.Failed() {
			t.Logf("%s Pass", t.Name())
		}
	})
}

// ! TB.Setenv 函数调用 os.Setenv, 并在测试完成后使用 `Cleanup` 进行还原。它不能在并行测试中使用。
func TestTBSetenv(t *testing.T) {
	//? go test -v -run=^TestTBSetenv&
	env := struct {
		key string
		val string
	}{
		key: "Setenv", val: "Hello World",
	}
	if v, ok := os.LookupEnv(env.key); ok {
		t.Logf("The value of Env `Setenv` is %s", v)
	} else {
		t.Setenv(env.key, env.val)
		if os.Getenv(env.key) == "Hello World" {
			t.Log("Set environment variable `Setenv:Hello World` successfully.")
		}
	}

	panickingRecover := func(name string) {
		if got := recover(); got != nil {
			t.Logf("panicking in %s: \n%#v.", name, got)
		} else {
			t.Logf("Test: %s PASS.", name)
		}
	}
	//? 测试完成后，将恢复 Setenv 进入（子）测试之前的值
	t.Run("RestoreWhileTestCompleted", func(t *testing.T) {
		t.Setenv(env.key, "dlroW olleH")
		t.Logf("Change env %s to %s", env.key, os.Getenv(env.key))
	})
	if os.Getenv(env.key) == env.val {
		t.Log("The environment variable `Setenv:Hello World` restores.")
	}

	//? Setenv 不能用于并行测试或具有并行祖先的测试
	t.Run("ParallelAfterSetenv", func(t *testing.T) {
		defer panickingRecover(t.Name())
		t.Setenv("Setenv", "Hello")
		t.Parallel()
	})

	t.Run("ParallelBeforeSetenv", func(t *testing.T) {
		defer panickingRecover(t.Name())
		t.Parallel()
		t.Setenv("Setenv", "Hello")
	})

	t.Run("ParallelParentBeforeSetenv", func(t *testing.T) {
		t.Parallel()
		t.Run("child", func(t *testing.T) {
			defer panickingRecover(t.Name())
			t.Setenv("Setenv", "Hello")
		})
	})
}

// ! TB.TempDir 返回一个临时目录供测试使用。当（子）测试完成时自动删除。
func TestTBTempDir(t *testing.T) {
	//? go test -v -run=^TestTBTempDir$

	var dir string
	fn_checkDir := func(dir string) {
		fi, err := os.Stat(dir)
		if fi != nil {
			t.Fatalf("Directory %q from user Cleanup still exists", dir)
		}
		if !os.IsNotExist(err) {
			t.Fatalf("Unexpected error: %v", err)
		}
	}

	t.Run("CheckExist", func(t *testing.T) {
		dirCh := make(chan string, 1)
		t.Cleanup(func() {
			// 验证目录 directory 已经在测试完成时被删除
			select {
			case dir := <-dirCh:
				fn_checkDir(dir)
			default:
				if !t.Failed() {
					t.Fatal("never received dir channel")
				}
			}
		})
		dir := t.TempDir()
		t.Logf("create a tempDir=%v\n", dir)
		dirCh <- dir // 传递 tempDir
	})

	t.Run("InCleanup", func(t *testing.T) {
		t.Helper()
		t.Run("test", func(t *testing.T) {
			t.Cleanup(func() {
				dir = t.TempDir()
			})
			_ = t.TempDir()
		})
		fn_checkDir(dir)
	})

	t.Run("InBenchmark", func(t *testing.T) {
		testing.Benchmark(func(b *testing.B) {
			if !b.Run("test", func(b *testing.B) {
				// Add a loop so that the test won't fail.
				for i := 0; i < b.N; i++ {
					_ = b.TempDir()
				}
			}) {
				t.Fatal("Sub test failure in a benchmark")
			}
		})
	})
}

// ! T.

// !
// !

func doSomething(up int) {
	if up < 0 {
		return
	}
	// random generate a slice
	src := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	s := rand.Perm(src.Int() % up)
	// reverse slice
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
