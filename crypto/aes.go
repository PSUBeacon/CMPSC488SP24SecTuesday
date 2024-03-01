package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// function to generate an encryption key if it doesn't exist
func generateRandomKey() string {
	key := make([]byte, 16) // 16-byte key
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generating random key:", err)
	}
	return fmt.Sprintf("%x", key)
}

// function to get an encryption key from an environment variable
func getEncryptionKey() []byte {
	// get the encryption key from an environment variable
	key := os.Getenv("AES_ENCRYPTION_KEY")

	// generate a random key if the key is not available in an environment variable
	if key == "" {
		key = generateRandomKey()
		err := os.Setenv("AES_ENCRYPTION_KEY", key)
		if err != nil {
			fmt.Println("Error setting environment variable:", err)
		}
	}

	return []byte(key)
}

// function to encrypt a message
func Encrypt(plaintext []byte) ([]byte, error) {
	// get the encryption key from the environment variable
	key := getEncryptionKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// function to decrypt a message
func Decrypt(ciphertext []byte) ([]byte, error) {
	// get the encryption key from the environment variable
	key := getEncryptionKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

//func main() {
//	// load .env file which is in gitignore
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	// message to encrypt, and later decrypt
//	plaintext := []byte("Hello, Golang Encryption!")
//
//	ciphertext, err := encrypt(plaintext)
//	if err != nil {
//		fmt.Println("Encryption error:", err)
//		return
//	}
//
//	decrypted, err := decrypt(ciphertext)
//	if err != nil {
//		fmt.Println("Decryption error:", err)
//		return
//	}
//
//	fmt.Printf("Encrypted: %x\n", ciphertext)
//	fmt.Println("Decrypted:", string(decrypted))
//	fmt.Println("Encryption key:", string(getEncryptionKey()))
//}
