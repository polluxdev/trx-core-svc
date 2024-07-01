package helper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestEncryptDecrypt(t *testing.T) {
	data := TestStruct{
		Name:  "Test",
		Value: 123,
	}

	encrypted := Encrypt(data)
	assert.NotEmpty(t, encrypted, "Encrypted string should not be empty")

	decrypted := Decrypt(encrypted)
	assert.NotEmpty(t, decrypted, "Decrypted string should not be empty")

	var result TestStruct
	err := json.Unmarshal([]byte(decrypted), &result)
	assert.NoError(t, err, "Decrypted data should unmarshal without error")

	assert.Equal(t, data, result, "Original and decrypted data should be the same")
}
