// go get "github.com/gtank/cryptopasta"
// 1. Allows AES blocks autofill
// 2. Non-repeatable blocks with repeatable payload

package main

import (
	"crypto/aes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gtank/cryptopasta"
)

func ExampleSHA256() {
	//md5, sha-1 - !NOT! USE!
	//SHA-256, SHA-512 USE!
	message := []byte("123456")
	hash := sha256.Sum256(message)
	fmt.Printf("%x\n", hash)
}

// GenerateHMAC создает HMAC подпись для сообщения
func GenerateHMAC(key, message []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC проверяет HMAC подпись
func VerifyHMAC(key, message []byte, signature string) bool {
	// Вычисляем ожидаемую подпись
	h := hmac.New(sha256.New, key)
	h.Write(message)
	expectedSignature := hex.EncodeToString(h.Sum(nil))

	// Сравниваем (используем безопасное сравнение)
	expectedBytes, err1 := hex.DecodeString(expectedSignature)
	receivedBytes, err2 := hex.DecodeString(signature)
	if err1 != nil || err2 != nil {
		return false
	}
	return hmac.Equal(expectedBytes, receivedBytes)
}

func ExampleHMAC() {
	key := []byte("secret-key-123")
	message := []byte("Hello, HMAC World!")

	// Генерация подписи
	signature := GenerateHMAC(key, message)
	fmt.Printf("Message: %s\n", string(message))
	fmt.Printf("HMAC Signature: %s\n\n", signature)

	// Проверка корректной подписи
	isValid := VerifyHMAC(key, message, signature)
	fmt.Printf("Valid signature check: %v\n", isValid)

	// Проверка неверной подписи
	isValid = VerifyHMAC(key, message, "wrong-signature-here")
	fmt.Printf("Invalid signature check: %v\n", isValid)

	// Проверка с измененным сообщением
	isValid = VerifyHMAC(key, []byte("Modified message"), signature)
	fmt.Printf("Modified message check: %v\n", isValid)
}

func ExampleAES_std() {
	key := []byte("examplekey123456") // 16, 24 или 32 байта
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return
	}
	plaintext := []byte("Hello, AES!!!!!!")
	ciphertext := make([]byte, len(plaintext))

	// ECB-режим в Go не встроен, но для примера
	// показываю блоковое шифрование без режима.
	block.Encrypt(ciphertext, plaintext)
	fmt.Printf("Зашифрованный текст: %x\n", ciphertext)
}

func ExampleAES_cryptopasta() {
	key := cryptopasta.NewEncryptionKey()      // Генерация ключа
	plaintext := []byte("Hello, cryptopasta!") // Текст для шифрования

	// Шифрование (автоматически использует GCM)
	ciphertext, err := cryptopasta.Encrypt(plaintext, key)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Encrypted: %x\n", ciphertext)

	// Расшифрование
	decrypted, err := cryptopasta.Decrypt(ciphertext, key)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}

func ExampleRSA() {
	// Генерация ключей
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	// Сообщение для шифрования
	message := []byte("Hello, RSA Encryption!")

	// Шифрование с помощью публичного ключа
	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		message,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Original: %s\n", message)
	fmt.Printf("Encrypted (base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))

	// Расшифрование с помощью приватного ключа
	plaintext, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		ciphertext,
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Decrypted: %s\n", plaintext)
}

func main() {
	ExampleSHA256()
	ExampleHMAC()
	ExampleAES_std()
	ExampleAES_cryptopasta()
	ExampleRSA()
}
