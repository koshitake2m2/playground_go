package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	privateKeyFile := getenv("SSH_PRIVATE_KEY", "../client_key")
	hostKeyFile := getenv("SSH_HOST_KEY", "../server_key.pub")
	serverAddr := getenv("SERVER_ADDR", "localhost:12222")
	csvFile := getenv("CSV_FILE", "sample.csv")

	privateKey, err := os.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatalf("read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		log.Fatalf("parse private key: %v", err)
	}

	hostKey, err := os.ReadFile(hostKeyFile)
	if err != nil {
		log.Fatalf("read host key: %v", err)
	}

	hostKeyParsed, _, _, _, err := ssh.ParseAuthorizedKey(hostKey)
	if err != nil {
		log.Fatalf("parse host key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "sftp-user",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			if ssh.FingerprintSHA256(key) == ssh.FingerprintSHA256(hostKeyParsed) {
				return nil
			}
			return fmt.Errorf("host key verification failed")
		},
	}

	conn, err := ssh.Dial("tcp", serverAddr, config)
	if err != nil {
		log.Fatalf("dial SSH: %v", err)
	}
	defer conn.Close()

	log.Printf("[client] connected to %s", serverAddr)

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("create SFTP client: %v", err)
	}
	defer client.Close()

	log.Printf("[client] SFTP client created")

	if _, err := os.Stat(csvFile); os.IsNotExist(err) {
		log.Fatalf("CSV file not found: %s", csvFile)
	}

	remotePath := filepath.Join("/", filepath.Base(csvFile))
	if err := uploadFile(client, csvFile, remotePath); err != nil {
		log.Fatalf("upload file: %v", err)
	}

	log.Printf("[client] uploaded %s to %s", csvFile, remotePath)
}

func uploadFile(client *sftp.Client, localPath, remotePath string) error {
	localFile, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("open local file: %w", err)
	}
	defer localFile.Close()

	remoteFile, err := client.Create(remotePath)
	if err != nil {
		return fmt.Errorf("create remote file: %w", err)
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		return fmt.Errorf("copy file: %w", err)
	}

	return nil
}
