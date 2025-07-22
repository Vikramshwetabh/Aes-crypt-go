package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func encrypt(text, key []byte) ([]byte, []byte) {
	//creates cipher block
	c, err := aes.NewCipher(key) //creates & return new cipher block
	if err != nil {
		panic(err)
	}
	//wrapping in a gcm. use the gcm mode from cipher
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	//create a nonce(IV) used for encryption by gcm
	nonce := make([]byte, gcm.NonceSize())
	// _, err = io.ReadFull(rand.Reader, nonce)
	// if err != nil {
	// 	panic(err)
	// }
	//encrpt the text
	ciphertext := gcm.Seal(nil, nonce, text, nil)
	return ciphertext, nonce
}
func decrypt(ciphertext, nonce, key []byte) []byte {
	c, err := aes.NewCipher(key) //creates & return new cipher block
	if err != nil {
		panic(err)
	}
	//wrapping in a gcm. use the gcm mode from cipher
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
	fmt.Println("AES Encryption Example")
	text := []byte("very secret text") //plain text is converted into byte slice
	//generating a random key of 24 bytes
	key := make([]byte, 24)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	ciphertext, nonce := encrypt(text, key)
	fmt.Println("encrypted text", hex.EncodeToString(ciphertext))

	plaintext := decrypt(ciphertext, nonce, key)
	fmt.Println("decrypted text:", string(plaintext))
}
