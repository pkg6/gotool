package crypto

import (
	"crypto/rand"
	"crypto/rsa"
)

// DefaultPrivateKey 私钥生成 openssl genrsa -out private_key.pem 1024
var DefaultPrivateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCsh4R0kNKqxY4kHKEwQg4QKOT8osBW8a7tcADYrdROXeZYvYrC
yhl4WDZ1Go9OyZA5qEycfYo/nj4+kGmtmdpTw80CVw0ejsnJzRH2n1R3GYn40jMG
xL0g5pu2Hvev+0LB9pS/5E2e2EOHtX5QhsJaGQULRO+3czO7r81wR6vfAQIDAQAB
AoGATuG3AcSlTUb98izU1cedvm20JH4VCqt9mzm2aVsw0pPEGZavttfIRWmvnGME
WrV1p6b3QCV17BhhxSEp8CGD3EA//7B/98bTFzrA6F/NiqSvyGbCHX7HeB28Muyv
qUlpHEG+PZHpZprMpB/lU3l4JJ9t6C31hfqjIBdC3W6zQfECQQDZ6u+HsS+2ECZI
Wudcfjfrf45+Yh8tTEpCq+WXjdAAR6F05JQwky4QWPi7+gPU8DM2l5hMHIRTciYL
b9rO1DobAkEAyq4GxsAG++Z1Owi1pl7SHEiRtW49vMj7xxC4BLVzo3ljDcgZPyJc
PYz/kxV6QmGzT32Mbtk/xGYTlporR2udEwJARnPSJQh/6FioR9Q74IdeBOEkbG/E
rJxxlcSFYc4TZUPDS0trLZkn11kscXmPK5TMueWg81p03ZWV/zSWhS/P6QJACN9n
conzhFGJbkUqVpcuEYjnwA6Ma1hNFWDY/XPIFS76NB8/Y7EoYpVqltDI4mEOjXtM
i4m9LebeEqi7HkxKuwJBAJS1/VL3t5mMk6V11fPcdEoac05ZAzHozYJNpQ/pATfP
Uts5VJdXKbTE2ha6HId6slHAz7pAbgT3iNStLWSPpz4=
-----END RSA PRIVATE KEY-----
`)

// DefaultPublicKey 公钥:根据私钥生成 openssl rsa -in private_key.pem -pubout -out public_key.pem
var DefaultPublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCsh4R0kNKqxY4kHKEwQg4QKOT8
osBW8a7tcADYrdROXeZYvYrCyhl4WDZ1Go9OyZA5qEycfYo/nj4+kGmtmdpTw80C
Vw0ejsnJzRH2n1R3GYn40jMGxL0g5pu2Hvev+0LB9pS/5E2e2EOHtX5QhsJaGQUL
RO+3czO7r81wR6vfAQIDAQAB
-----END PUBLIC KEY-----
`)

// RsaEncrypt rsa encrypt
func RsaEncrypt(plainText, publicKey []byte) ([]byte, error) {
	// parsing the public key
	if len(publicKey) == 0 {
		publicKey = DefaultPublicKey
	}
	pub, err := byteToPub(publicKey)
	if err != nil {
		return nil, err
	}
	// encode
	return rsa.EncryptPKCS1v15(rand.Reader, pub, plainText)
}

// RsaDecrypt rsa decrypt
func RsaDecrypt(cipherText, privateKey []byte) ([]byte, error) {
	// parse private keys in PKCS1 format
	if len(privateKey) == 0 {
		privateKey = DefaultPrivateKey
	}
	private, err := byteToPriv(privateKey)
	if err != nil {
		return nil, err
	}
	// decode
	return rsa.DecryptPKCS1v15(rand.Reader, private, cipherText)
}
