package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// PKCS7Padding PKCS7填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	pad := make([]byte, padLen)
	for i := range padLen {
		pad[i] = byte(padLen)
	}
	return append(data, pad...)
}

// PKCS7UnPadding 去除PKCS7填充
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data empty")
	}
	padLen := int(data[length-1])
	if padLen > length || padLen == 0 {
		return nil, errors.New("invalid pkcs7 padding")
	}
	// 校验填充字节
	for i := range padLen {
		if data[length-1-i] != byte(padLen) {
			return nil, errors.New("invalid pkcs7 padding")
		}
	}
	return data[:length-padLen], nil
}

// AesCBCEncrypt AES-CBC PKCS7加密
// key: 密钥 AES支持16/24/32字节（AES-128/AES-192/AES-256）
// iv: 初始向量，必须16字节
func AesCBCEncrypt(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		return nil, errors.New("iv length must equal block size(16)")
	}
	blockSize := block.BlockSize()
	plainText = PKCS7Padding(plainText, blockSize)
	cipherText := make([]byte, len(plainText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

// AesCBCDecrypt AES-CBC PKCS7解密
func AesCBCDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(iv) != block.BlockSize() {
		return nil, errors.New("iv length must equal block size(16)")
	}
	if len(cipherText)%block.BlockSize() != 0 {
		return nil, errors.New("ciphertext not multiple of block size")
	}
	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)
	return PKCS7UnPadding(plainText)
}

func main() {
	// AES-128 key 16字节，AES-256是32字节
	key := []byte("1234567890123456")
	iv := []byte("abcdef1234567890")
	origin := []byte("hello aes cbc pkcs7 test")

	enc, err := AesCBCEncrypt(origin, key, iv)
	if err != nil {
		fmt.Println("encrypt err:", err)
		return
	}
	fmt.Printf("cipher bytes: %x\n", enc)

	dec, err := AesCBCDecrypt(enc, key, iv)
	if err != nil {
		fmt.Println("decrypt err:", err)
		return
	}
	fmt.Printf("plain: %s\n", dec)
}
