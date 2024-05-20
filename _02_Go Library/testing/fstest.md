<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="XXXX" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>
<a id="TOP"></a>

## [Go testing/fstest](https://pkg.go.dev/testing/fstest)

包 `fstest` 实现了对测试实现和文件系统用户的支持。

---
### MapFS

`MapFS` 是一个简单的内存中文件系统，用于测试，表示为从路径名（`MapFS.Open` 的参数）到它们所表示的文件或目录信息的映射。 


```go
type MapFS map[string]*MapFile

```