package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// CbcEncrypt aes cbc encrypt , The most common method of aes encryption
func CbcEncrypt(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = pKCS7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypt := make([]byte, len(plainText))
	blockMode.CryptBlocks(crypt, plainText)
	return []byte(base64.StdEncoding.EncodeToString(crypt)), nil
}

// CbcDecrypt aes cbc decrypt
func CbcDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	cryptByte, err := base64.StdEncoding.DecodeString(string(cipherText)) // to byte array
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key) // group secret key
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv) // encryption mode
	plainText := make([]byte, len(cryptByte))      // create array
	blockMode.CryptBlocks(plainText, cryptByte)    // decode
	return pKCS7UnPadding(plainText), nil          // to completion 去补全码
}

// PKCS7Padding 补码
func pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// PKCS7UnPadding 去码
func pKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}

// EcbEncrypt aes ecb encrypt
func EcbEncrypt(src, key []byte) (encrypted []byte, err error) {
	cipherText, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil, err
	}
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// block encryption 分组分块加密
	for bs, be := 0, cipherText.BlockSize(); bs <= len(src); bs, be = bs+cipherText.BlockSize(), be+cipherText.BlockSize() {
		cipherText.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encrypted, nil
}

// EcbDecrypt aes ecb decrypt
func EcbDecrypt(encrypted, key []byte) (decrypted []byte, err error) {
	cipherText, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil, err
	}
	decrypted = make([]byte, len(encrypted))
	for bs, be := 0, cipherText.BlockSize(); bs < len(encrypted); bs, be = bs+cipherText.BlockSize(), be+cipherText.BlockSize() {
		cipherText.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim], nil
}

// GenerateKey generate key
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// CfbEncrypt aes cfb encrypt
func CfbEncrypt(plainText, key, iv []byte) ([]byte, error) {
	if len(iv) < 16 {
		return nil, errors.New("iv length at least 16")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return cipherText, nil
}
func CfbDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	if len(iv) < 16 {
		return nil, errors.New("iv length at least 16")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText, nil
}

// Ctr aes ctr encrypt and decrypt
func Ctr(text, key, iv []byte) ([]byte, error) {
	if len(iv) < 16 {
		return nil, errors.New("iv length at least 16")
	}
	// 指定加密,解密算法为AES,返回一个AES的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 指定分组模式
	blockMode := cipher.NewCTR(block, iv)
	// 执行加密,解密操作
	message := make([]byte, len(text))
	blockMode.XORKeyStream(message, text)
	// 返回明文或密文
	return message, nil
}

// Ofb aes ofb encrypt and decrypt
func Ofb(text, key, iv []byte) ([]byte, error) {
	if len(iv) < 16 {
		return nil, errors.New("iv length at least 16")
	}
	// 指定加密,解密算法为AES,返回一个AES的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 指定分组模式
	blockMode := cipher.NewOFB(block, iv)
	// 执行加密,解密操作
	message := make([]byte, len(text))
	blockMode.XORKeyStream(message, text)
	// 返回明文或密文
	return message, nil
}
