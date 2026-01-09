package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
	
	// Determine OS from environment variable or runtime
	targetOS := os.Getenv("TARGET_OS")
	if targetOS == "" {
		targetOS = runtime.GOOS
	}
	
	var sevenZipPath string
	
	// Check if we should use system 7z command (for Dockerfile builds)
	if os.Getenv("USE_SYSTEM_7Z") == "1" {
		// Use system 7z command (p7zip in Alpine Linux)
		sevenZipPath = "7z"
	} else {
		// Use bundled binaries
		execPath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("failed to get executable path: %v", err)
		}
		
		switch targetOS {
		case "darwin":
			sevenZipPath = filepath.Join(filepath.Dir(execPath), "bin", "7zz")
			if _, err := os.Stat(sevenZipPath); os.IsNotExist(err) {
				sevenZipPath = filepath.Join(".", "bin", "7zz")
			}
		case "linux":
			// Check multiple possible paths for ko
			possiblePaths := []string{
				"/var/run/ko/bin/linux/7zzs",
				"/var/run/ko/bin/linux/7zz",
				"/ko-app/bin/linux/7zzs",
				"/ko-app/bin/linux/7zz",
				filepath.Join(filepath.Dir(execPath), "kodata", "bin", "linux", "7zzs"),
				filepath.Join(filepath.Dir(execPath), "kodata", "bin", "linux", "7zz"),
				filepath.Join(filepath.Dir(execPath), "bin", "linux", "7zzs"),
				filepath.Join(filepath.Dir(execPath), "bin", "linux", "7zz"),
				filepath.Join(".", "bin", "linux", "7zzs"),
				filepath.Join(".", "bin", "linux", "7zz"),
			}
			
			for _, path := range possiblePaths {
				if _, err := os.Stat(path); err == nil {
					sevenZipPath = path
					break
				}
			}
			
			if sevenZipPath == "" {
				return fmt.Errorf("7zz binary not found in any of the expected paths: %v", possiblePaths)
			}
		default:
			return fmt.Errorf("unsupported OS: %s", targetOS)
		}
	}
	
	
	cmd := exec.Command(sevenZipPath, "a", "-p"+password, outputFile, inputFile)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("7z command failed: %v\nOutput: %s", err, output)
	}

	return nil
}