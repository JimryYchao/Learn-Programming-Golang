<a id="TOP"></a>

## Package crypto

<div id="top" style="z-index:99999999;position:fixed;bottom:35px;right:50px;float:right">
	<a href="./code/crypto_test.go" target="_blank"><img id="img-code" src="../_rsc/to-code.drawio.png" ></img></a>
	<!-- <a href="#TOP" ><img id="img-top" src="../_rsc/to-top.drawio.png" ></img></a>	 -->
	<a href="https://pkg.go.dev/crypto"  target="_blank"><img id="img-link" src="../_rsc/to-link.drawio.png" ></img></a>
	<a href="..\README.md"><img id="img-back" src="../_rsc/back.drawio.png"></img></a>
</div>

包 `crypto` 收集常用的加密常量。

`Hash` 标识在其他包中实现的加密哈希函数：

```go
const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b
)
```

包 `ase` 实现 AES 加密。此包中的 AES 操作不是使用恒定时间算法实现的。但运行在启用了 AES 硬件支持的系统上时，这些操作是恒定时间的.

包 `cipher` 实现了标准的分组密码模式，这些模式可以封装在低级分组密码的实现中。

包 `des` 实现 Data Encryption Standard (DES) 和 Triple Data Encryption Algorithm (TDEA) 加密。DES 在密码学上被破坏，不应用于安全应用程序。

包 `ecdh` 实现了基于 NIST 曲线和 Curve25519 的 Elliptic Curve Diffie-Hellman (ECDH)。

包 `ecdsa` 实现了在 FIPS 186-4 和 SEC 1 2.0 中定义的椭圆曲线数字签名算法。这个包生成的签名不是确定性的，但是熵 (entropy) 与私钥和消息混合，在随机性源失败的情况下实现相同的安全级别。

包 `ed25519` 实现了 [Ed25519](https://ed25519.cr.yp.to/) 签名算法。这些函数还与 RFC 8032 中定义的 “Ed 25519” 函数兼容。然而与 RFC 8032 的表述不同，该包的私钥表示包括公钥后缀，以使使用同一密钥的多个签名操作更高效。该包将 RFC 8032 私钥称为 “seed”。




---
<a id="exam" ><a>