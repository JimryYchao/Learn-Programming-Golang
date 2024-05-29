### 标准 Go Library

- Go Version = 1.22.1
- ✔️ = Support in Go 1.22.1
- ❌ = Deprecated in Go 1.22.1
- ✖️ = May be deprecated in the future
<!-- ✔️/✖️/❌  -->

>---
|      | Name                            | Synopsis                                                    | Ref                                             |
| :--- | :------------------------------ | :---------------------------------------------------------- | :---------------------------------------------- |
| ✔️    | [io](./io/io.md)                | 包 `io` 提供 I/O 原语的基本接口。                           | $[[↗]](https://golang.google.cn/pkg/io/)        |
| ✔️    | [io/fs](./io/io_fs.md)          | 包 `fs` 定义了文件系统的基本接口。                          | $[[↗]](https://golang.google.cn/pkg/io/fs/)     |
| ❌    | [io/ioutil](./io/io_ioutil.md)  | 包 `ioutil` 实现了一些 I/O 实用程序功能。在 1.16 之后弃用。 | $[[↗]](https://golang.google.cn/pkg/io/ioutil/) |
| ✔️    | [buildin](./buildin/buildin.md) | 包 `buildin` 提供了 Go 语言预先声明标识符的文档。           | $[[↗]](https://golang.google.cn/pkg/builtin/)   |
|      | bufio                           |                                                             |

<!--
 fmt    

        
archive/zip	
arena

bytes	
cmd/addr2line	
cmd/asm	
cmd/buildid	
cmd/cgo	Cgo 
cmd/compile	
cmd/covdata	
cmd/cover	
cmd/dist	
cmd/distpack	
cmd/doc	
cmd/fix	
cmd/go	
cmd/gofmt	   
cmd/link	   
cmd/nm	       
cmd/objdump    
cmd/pack	   
cmd/pprof	   
cmd/test2json	
cmd/tools	
cmd/trace	   
cmd/trace/v2
cmd/vet	       
cmp            
compress/bzip2	       
compress/flate	       
compress/gzip	       
compress/lzw	
compress/zlib
container/heap
container/list
container/ring	
context	            
crypto	            
crypto/aes	        
crypto/boring	    
crypto/cipher	    
crypto/des	        
crypto/dsa	        
crypto/ecdh	        
crypto/ecdsa	    
crypto/ed25519	    
crypto/elliptic	    
crypto/hmac	
crypto/md5	    
crypto/rand	    
crypto/rc4	    
crypto/rsa	    
crypto/sha1	    
crypto/sha256	    
crypto/sha512	    
crypto/subtle	        
crypto/tls	       
crypto/tls/fipso    
crypto/x509	           
crypto/x509/pkix	        
database	
database/sql	           
database/sql/driver	        
debug	
buildinfo	            
dwarf	            
elf	                
gosym	            
macho	            
pe	              
plan9obj            
embed	            
encoding            
ascii85	        
asn1	        
base32	        
base64	        
binary	        
csv	       
gob	       
hex	       
json	        
pem	           
xml	           
errors	        
expvar	        
flag	        
go	
ast	    
build   
constraint  
constant	    
doc	Package     
comment	             
format	        
importer	    
parser	        
printer	    
scanner	    
token	    
types	    
version	    
hash	    
adler32	    
crc32	
crc64	
fnv	        
maphash	
html	
template    
image	
color	
palette	
draw	
gif         
jpeg	    
png	        
index	
suffixarray 
fs	    
iter	    
log	        
slog	        
syslog	    
maps	    
math	    
big	        
bits	    
cmplx	    
rand	    
v2	          
mime	    
multipart	            
quotedprintable 	    
net	                    
http	                
cgi	                    
cookiejar	            
fcgi	                
httptest	            
httptrace	            
httputil	            
pprof	                
mail	                
netip	                
rpc	                    
jsonrpc	                
smtp	                
textproto	    
url	            
os	            
exec	        
signal	        
user	        
path	        
filepath	    
plugin	        
reflect	        
regexp	        
syntax	        
runtime	        
asan	
cgo	        
coverage	    
debug	    
metrics	    
msan	
pprof	    
race	    
trace	    
slices	    
sort	    
strconv	    
strings	    
sync	    
atomic	    
syscall	    
js	                
testing	    
fstest	    
iotest	    
quick	    
slogtest	    
text	
scanner	    
tabwriter	    
template	    
parse	    
time	    
tzdata	    
unicode	    
utf16	    
utf8	    
unsafe	    

