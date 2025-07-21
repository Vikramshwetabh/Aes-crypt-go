package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	fmt.Println("AES Encryption Example")
	text := []byte("very secret text")
	//generating a random key of 16 bytes
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	//creates cipher block
	c, err := aes.NewCipher(key) //creates & return new cipher block
	if err != nil {
		panic(err)
	}
	//wrapping in a gcm
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	//create a nonce(IV) used for encryption by gcm
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}

	//encrpt the text
	ciphertext := gcm.Seal(nil, nonce, text, nil)
	fmt.Println("encrypted text", hex.EncodeToString(ciphertext))
}
