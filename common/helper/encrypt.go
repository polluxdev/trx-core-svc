package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/polluxdev/trx-core-svc/application/config"
	"github.com/polluxdev/trx-core-svc/common/utils"
)

func Encrypt(value interface{}) string {
	plainText, err := json.Marshal(value)
	if err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	block, err := aes.NewCipher([]byte(config.Aes.Key))
	if err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.URLEncoding.EncodeToString(cipherText)
}

func Decrypt(encryptedText string) string {
	cipherText, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	block, err := aes.NewCipher([]byte(config.Aes.Key))
	if err != nil {
		panic(utils.InvariantError(err.Error(), err))
	}

	if len(cipherText) < aes.BlockSize {
		panic(utils.InvariantError("ciphertext too short", nil))
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
