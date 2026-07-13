package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(password, salt string, keyLen int) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), 100000, keyLen, sha256.New)
}

func GenerateSalt() (string, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashPassword(password, salt string) string {
	key := DeriveKey(password, salt, 32)
	return hex.EncodeToString(key)
}

func EncryptFileName(fileName, password, salt string) string {
	key := DeriveKey(password, salt, 32)
	hash := sha256.Sum256(append([]byte(fileName), key...))
	return hex.EncodeToString(hash[:])
}

func EncryptFile(inputPath, outputPath, password, salt string) error {
	key := DeriveKey(password, salt, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %w", err)
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(nonce); err != nil {
		return fmt.Errorf("failed to write nonce: %w", err)
	}

	// Read the input file content
	inputContent, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Encrypt the content
	ciphertext := gcm.Seal(nil, nonce, inputContent, nil)

	// Write the encrypted content
	if _, err := outputFile.Write(ciphertext); err != nil {
		return fmt.Errorf("failed to write encrypted content: %w", err)
	}

	return nil
}

func DecryptFile(inputPath, outputPath, password, salt string) error {
	key := DeriveKey(password, salt, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inputFile.Close()

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(inputFile, nonce); err != nil {
		return fmt.Errorf("failed to read nonce: %w", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	ciphertext, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read ciphertext: %w", err)
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("failed to decrypt: %w", err)
	}

	if _, err := outputFile.Write(plaintext); err != nil {
		return fmt.Errorf("failed to write plaintext: %w", err)
	}

	return nil
}

func GetStoragePath(projectID string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, "项目文档管理系统", projectID)
}
