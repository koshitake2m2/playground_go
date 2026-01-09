package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <password>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	password := os.Args[2]

	if err := encryptFile(inputFile, password); err != nil {
		log.Fatalf("Error encrypting file: %v", err)
	}

	fmt.Printf("Successfully encrypted %s\n", inputFile)
}

func encryptFile(inputFile, password string) error {
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return fmt.Errorf("input file %s does not exist", inputFile)
	}

	outputFile := inputFile + ".7z"
	
	// Use system 7z command (p7zip in Alpine Linux)
	sevenZipPath := "7z"
	
	cmd := exec.Command(sevenZipPath, "a", "-p"+password, outputFile, inputFile)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("7z command failed: %v\nOutput: %s", err, output)
	}

	return nil
}