package gostd

import "testing"

// 包 fstest 实现了对测试实现和文件系统用户的支持

/*
! MapFS (map[string]*MapFile) 是内存中的一个用于测试的简单文件系统，表示为从路径名（Open 的参数）到它们所表示的文件或目录信息的映射
! MapFile 描述为 MapFS 映射中的单个文件
	Glob
	Open
	ReadDir
	ReadFile
	Stat
	Sub

*/
func TestMapFS(t *testing.T) {

}
