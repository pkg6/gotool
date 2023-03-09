package crypto

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"hash"
	"hash/crc32"
	"os"
)

func HashHmacSumNil(hash hash.Hash, w []byte) []byte {
	hash.Write(w)
	return hash.Sum(nil)
}

func Md5Byte(p []byte) []byte {
	return []byte(hex.EncodeToString(HashHmacSumNil(md5.New(), p)))
}

func Sha1Byte(p []byte) []byte {
	return []byte(hex.EncodeToString(HashHmacSumNil(sha1.New(), p)))
}

func Sha256Byte(p []byte) []byte {
	return []byte(hex.EncodeToString(HashHmacSumNil(sha256.New(), p)))
}

func Sha512Byte(p []byte) []byte {
	return []byte(hex.EncodeToString(HashHmacSumNil(sha512.New(), p)))

}

func Md5String(s string) string {
	return hex.EncodeToString(HashHmacSumNil(md5.New(), []byte(s)))
}

func Sha1String(s string) string {
	return hex.EncodeToString(HashHmacSumNil(sha1.New(), []byte(s)))
}

func Sha256String(s string) string {
	return hex.EncodeToString(HashHmacSumNil(sha256.New(), []byte(s)))
}

func Sha512String(s string) string {
	return hex.EncodeToString(HashHmacSumNil(sha512.New(), []byte(s)))
}

func Crc32(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

func HashHmac(algo func() hash.Hash, data, key []byte) string {
	return hex.EncodeToString(HashHmacSumNil(hmac.New(algo, key), data))
}

func Md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	var size int64 = 1048576 // 1M
	h := md5.New()
	if fi.Size() < size {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		h.Write(data)
	} else {
		b := make([]byte, size)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}
			h.Write(b[:n])
		}
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func Sha1File(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func HashPriv(hash crypto.Hash, hashed, privateKey []byte) ([]byte, error) {
	priv, err := byteToPriv(privateKey)
	if err != nil {
		return nil, err
	}
	if hash.Size() != len(hashed) {
		hashed, _ = hex.DecodeString(string(hashed))
	}
	signByte, err := rsa.SignPKCS1v15(rand.Reader, priv, hash, hashed)
	return []byte(base64.StdEncoding.EncodeToString(signByte)), err
}

func VerifyHashPub(hash crypto.Hash, hashed, data, privateKey []byte) error {
	sign, err := base64.StdEncoding.DecodeString(string(data))
	pub, err := byteToPub(privateKey)
	if err != nil {
		return err
	}
	if hash.Size() != len(hashed) {
		hashed, _ = hex.DecodeString(string(hashed))
	}
	return rsa.VerifyPKCS1v15(pub, hash, hashed, sign)
}

func byteToPriv(privateKey []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	// parse private keys in PKCS1 format
	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return pri, nil
	}
	pri2, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pri2.(*rsa.PrivateKey), nil
}

func byteToPub(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		pub, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	return pub.(*rsa.PublicKey), nil
}
