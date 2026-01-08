package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ProtonMail/gopenpgp/v3/crypto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  Encrypt file: go run main.go encrypt <input_file> <public_key_file>")
		fmt.Println("  Decrypt file: go run main.go decrypt <encrypted_file> <private_key_file> <passphrase>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "encrypt":
		if len(os.Args) < 4 {
			log.Fatal("Usage: go run main.go encrypt <input_file> <public_key_file>")
		}
		encryptFile(os.Args[2], os.Args[3])
	case "decrypt":
		if len(os.Args) < 5 {
			log.Fatal("Usage: go run main.go decrypt <encrypted_file> <private_key_file> <passphrase>")
		}
		decryptFile(os.Args[2], os.Args[3], os.Args[4])
	default:
		log.Fatal("Unknown command. Use: encrypt or decrypt")
	}
}


func encryptFile(inputFile, publicKeyFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	publicKeyArmored, err := os.ReadFile(publicKeyFile)
	if err != nil {
		log.Fatalf("Error reading public key: %v", err)
	}

	publicKey, err := crypto.NewKeyFromArmored(string(publicKeyArmored))
	if err != nil {
		log.Fatalf("Error parsing public key: %v", err)
	}

	pgp := crypto.PGP()
	
	// Create encryption handle
	encHandle, err := pgp.Encryption().Recipient(publicKey).New()
	if err != nil {
		log.Fatalf("Error creating encryption handle: %v", err)
	}

	// Encrypt the message
	pgpMessage, err := encHandle.Encrypt(data)
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
	}

	// Get the armored encrypted data
	armored, err := pgpMessage.ArmorBytes()
	if err != nil {
		log.Fatalf("Error getting armored data: %v", err)
	}

	outputFile := inputFile + ".pgp"
	if err := os.WriteFile(outputFile, armored, 0644); err != nil {
		log.Fatalf("Error writing encrypted file: %v", err)
	}

	fmt.Printf("File encrypted successfully: %s\n", outputFile)
}

func decryptFile(encryptedFile, privateKeyFile, passphrase string) {
	encryptedData, err := os.ReadFile(encryptedFile)
	if err != nil {
		log.Fatalf("Error reading encrypted file: %v", err)
	}

	privateKeyArmored, err := os.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("Error reading private key: %v", err)
	}

	privateKey, err := crypto.NewPrivateKeyFromArmored(string(privateKeyArmored), []byte(passphrase))
	if err != nil {
		log.Fatalf("Error parsing private key: %v", err)
	}

	pgp := crypto.PGP()

	// Create decryption handle
	decHandle, err := pgp.Decryption().DecryptionKey(privateKey).New()
	if err != nil {
		log.Fatalf("Error creating decryption handle: %v", err)
	}

	// Decrypt the message
	decrypted, err := decHandle.Decrypt(encryptedData, crypto.Armor)
	if err != nil {
		log.Fatalf("Error decrypting data: %v", err)
	}

	outputFile := encryptedFile + ".decrypted"
	if err := os.WriteFile(outputFile, decrypted.Bytes(), 0644); err != nil {
		log.Fatalf("Error writing decrypted file: %v", err)
	}

	fmt.Printf("File decrypted successfully: %s\n", outputFile)
}