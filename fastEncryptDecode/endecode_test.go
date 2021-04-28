package fastEncryptDecode

import (
	"encoding/base64"
	"testing"
)

func TestMD5hash(T *testing.T) {
	str := "123456"
	md := MD5hash([]byte(str))
	T.Log("str", str)
	T.Log("str MD5 ", md)
	ok := MD5Verify(str, md)
	if !ok {
		T.Error("MD5Verify err")
	}
}

const (
	plantText    string = "tangxvhuitangxvhui"                           //加密前
	aesKey       string = "1234567890123456"                             //密钥
	base64String string = "imeIhfwP+X+a9f3Mh3jOUgmkhBtrobb0vXubnyDg/3s=" //加密后
)

func TestAesEncryptCBC(T *testing.T) {
	btString, err := AesEncryptCBC([]byte(plantText), []byte(aesKey))
	if err != nil {
		T.Error(err)
	}
	base64String := base64.StdEncoding.EncodeToString(btString)
	T.Log("base64:=", base64String)
}

func TestAesDecryptCBC(T *testing.T) {
	btstring, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		T.Error("base64 err", err)
	}
	btString, err := AesDecryptCBC([]byte(btstring), []byte(aesKey))
	if err != nil {
		T.Error(err)
	}

	if plantText != string(btString) {
		T.Error("AesDecryptCBC err ")
	}

}
