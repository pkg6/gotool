package crypto

import "encoding/base64"

// Base64Encrypt base64 encrypt
func Base64Encrypt(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// Base64Decrypt base64 decrypt
func Base64Decrypt(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
