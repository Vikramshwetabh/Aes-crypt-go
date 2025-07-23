package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

type Input struct {
	Message string `json:"message"`
}

func encrypt(text, key []byte) ([]byte, []byte) {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}
	ciphertext := gcm.Seal(nil, nonce, text, nil)
	return ciphertext, nonce
}

func decrypt(ciphertext, nonce, key []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	return plaintext
}

func main() {
	fmt.Println("Enter JSON input (e.g. {\"message\": \"very secret text\"}):")
	var input Input
	decoder := json.NewDecoder(os.Stdin)
	if err := decoder.Decode(&input); err != nil {
		panic(err)
	}

	key := make([]byte, 24)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	ciphertext, nonce := encrypt([]byte(input.Message), key)
	fmt.Println("encrypted text:", hex.EncodeToString(ciphertext))
	fmt.Println("nonce:", hex.EncodeToString(nonce))
	fmt.Println("key:", hex.EncodeToString(key))

	plaintext := decrypt(ciphertext, nonce, key)
	fmt.Println("decrypted text:", string(plaintext))
}
