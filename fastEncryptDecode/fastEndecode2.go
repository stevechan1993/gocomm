package fastEncryptDecode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

//md5 相关
//string to md5
func MD5hash(code []byte) string {
	h := md5.New()
	h.Write(code)
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//md5校验
func MD5Verify(code string, md5Str string) bool {
	codeMD5 := MD5hash([]byte(code))
	return 0 == strings.Compare(codeMD5, md5Str)
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "fastEncryptDecode/fastEnDeCode: invalid key size " + strconv.Itoa(int(k)) + " | key size must be 16"
}

func checkKeySize(key []byte) error {
	len := len(key)
	if len != 16 {
		return KeySizeError(len)
	}
	return nil
}

// AES encrypt pkcs7padding CBC, key for choose algorithm
func AesEncryptCBC(plantText, key []byte) ([]byte, error) {
	err := checkKeySize(key)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())
	//偏转向量iv长度等于密钥key块大小
	iv := key[:block.BlockSize()]
	blockModel := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(plantText))

	blockModel.CryptBlocks(cipherText, plantText)
	return cipherText, nil
}

func AesDecryptCBC(cipherText, key []byte) ([]byte, error) {
	err := checkKeySize(key)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//偏转向量iv长度等于密钥key块大小
	iv := key[:block.BlockSize()]
	blockModel := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plantText, cipherText)
	plantText = PKCS7UnPadding(plantText, block.BlockSize())
	return plantText, nil
}

//AES Decrypt pkcs7padding CBC, key for choose algorithm
func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	return plantText[:(length - unPadding)]
}

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
