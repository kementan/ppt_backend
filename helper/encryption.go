package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"gitlab.com/xsysproject/ppt_backend/config"
)

func Encrypt(plaintext, t string) (string, error) {
	config, err := config.LoadConfig("./.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	key := []byte(config.APPKey)

	var iv []byte

	if t == "f" {
		iv = make([]byte, aes.BlockSize)
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return "", err
		}
	} else {
		iv, err = hex.DecodeString(config.ElasticIVKey)
		if err != nil {
			return "", err
		}
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedPlaintext := padPlaintext([]byte(plaintext), aes.BlockSize)

	mode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(paddedPlaintext))
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	ciphertext = append(iv, ciphertext...)

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return encodedCiphertext, nil
}

func Decrypt(encodedCiphertext, t string) (string, error) {
	config, err := config.LoadConfig("./.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	key := []byte(config.APPKey)

	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext length")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]

	if t != "f" {
		iv, err = hex.DecodeString(config.ElasticIVKey)
		if err != nil {
			return "", err
		}
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedText := make([]byte, len(ciphertext)-aes.BlockSize)
	mode.CryptBlocks(decryptedText, ciphertext[aes.BlockSize:])

	unpaddedText := unpadPlaintext(decryptedText)

	return string(unpaddedText), nil
}

func padPlaintext(plaintext []byte, blockSize int) []byte {
	padding := blockSize - (len(plaintext) % blockSize)
	paddedPlaintext := make([]byte, len(plaintext)+padding)
	copy(paddedPlaintext, plaintext)
	for i := len(plaintext); i < len(paddedPlaintext); i++ {
		paddedPlaintext[i] = byte(padding)
	}
	return paddedPlaintext
}

func unpadPlaintext(paddedPlaintext []byte) []byte {
	padding := int(paddedPlaintext[len(paddedPlaintext)-1])
	if padding >= len(paddedPlaintext) {
		return paddedPlaintext
	}
	unpaddedText := paddedPlaintext[:len(paddedPlaintext)-padding]
	return unpaddedText
}
