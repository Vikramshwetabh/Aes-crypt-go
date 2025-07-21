package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("AES Encryption Example")
	text := []byte("very secret text") //plain text is converted into byte slice
	//generating a random key of 16 bytes
	key := make([]byte, 24)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
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
	// }s

	//encrpt the text
	ciphertext := gcm.Seal(nil, nonce, text, nil)
	fmt.Println("encrypted text", hex.EncodeToString(ciphertext))

	normalText, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("decrypted text:", string(normalText))

}
