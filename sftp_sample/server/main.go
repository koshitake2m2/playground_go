package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func mustGetenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	hostKeyFile := mustGetenv("SSH_HOST_KEY", "../server_key")
	authorizedKeyFile := mustGetenv("SSH_AUTHORIZED_KEY", "../client_key.pub")
	addr := mustGetenv("LISTEN_ADDR", ":12222")
	uploadDir := mustGetenv("UPLOAD_DIR", "./uploads")

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatalf("create upload directory: %v", err)
	}

	hostKey, err := os.ReadFile(hostKeyFile)
	if err != nil {
		log.Fatalf("read host key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(hostKey)
	if err != nil {
		log.Fatalf("parse host key: %v", err)
	}

	authorizedKey, err := os.ReadFile(authorizedKeyFile)
	if err != nil {
		log.Fatalf("read authorized key: %v", err)
	}

	authorizedKeyParsed, _, _, _, err := ssh.ParseAuthorizedKey(authorizedKey)
	if err != nil {
		log.Fatalf("parse authorized key: %v", err)
	}

	config := &ssh.ServerConfig{
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			if ssh.FingerprintSHA256(key) == ssh.FingerprintSHA256(authorizedKeyParsed) {
				return &ssh.Permissions{
					Extensions: map[string]string{
						"pubkey-fp": ssh.FingerprintSHA256(key),
					},
				}, nil
			}
			return nil, fmt.Errorf("unknown public key")
		},
	}
	config.AddHostKey(signer)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}
	defer listener.Close()

	log.Printf("[server] listening on %s", addr)
	log.Printf("[server] upload directory: %s", uploadDir)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[server] accept error: %v", err)
			continue
		}
		go handleConnection(conn, config, uploadDir)
	}
}

func handleConnection(conn net.Conn, config *ssh.ServerConfig, uploadDir string) {
	defer conn.Close()

	sshConn, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Printf("[server] SSH handshake failed: %v", err)
		return
	}
	defer sshConn.Close()

	log.Printf("[server] new SSH connection from %s", sshConn.RemoteAddr())

	go ssh.DiscardRequests(reqs)

	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Printf("[server] accept channel failed: %v", err)
			continue
		}

		go func(in <-chan *ssh.Request) {
			for req := range in {
				ok := false
				if req.Type == "subsystem" && string(req.Payload[4:]) == "sftp" {
					ok = true
				}
				req.Reply(ok, nil)
			}
		}(requests)

		handler := &fileSystemHandler{
			uploadDir: uploadDir,
			fileMap:   make(map[string]string),
			mu:        &sync.Mutex{},
		}
		handlers := sftp.Handlers{
			FileGet:  handler,
			FilePut:  handler,
			FileCmd:  handler,
			FileList: handler,
		}
		server := sftp.NewRequestServer(channel, handlers)
		if server == nil {
			log.Printf("[server] create SFTP server failed")
			channel.Close()
			continue
		}

		go func() {
			if err := server.Serve(); err != nil && err != io.EOF {
				log.Printf("[server] SFTP server error: %v", err)
			}
			server.Close()
		}()
	}
}

type fileSystemHandler struct {
	uploadDir string
	fileMap   map[string]string // ファイルハンドルID -> ファイルパス
	mu        *sync.Mutex
}

func (fs *fileSystemHandler) Fileread(req *sftp.Request) (io.ReaderAt, error) {
	filePath := filepath.Join(fs.uploadDir, req.Filepath)
	return os.Open(filePath)
}

func (fs *fileSystemHandler) Filewrite(req *sftp.Request) (io.WriterAt, error) {
	filePath := filepath.Join(fs.uploadDir, req.Filepath)
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}

	fs.mu.Lock()
	fs.fileMap[req.Filepath] = filePath
	fs.mu.Unlock()

	return &fileWriter{
		file:     file,
		handler:  fs,
		filepath: req.Filepath,
	}, nil
}

func (fs *fileSystemHandler) Filecmd(req *sftp.Request) error {
	filePath := filepath.Join(fs.uploadDir, req.Filepath)
	switch req.Method {
	case "Setstat":
		return nil
	case "Rename":
		return os.Rename(filePath, filepath.Join(fs.uploadDir, req.Target))
	case "Rmdir":
		return os.RemoveAll(filePath)
	case "Remove":
		return os.Remove(filePath)
	case "Mkdir":
		return os.MkdirAll(filePath, 0755)
	case "Symlink":
		return os.Symlink(filePath, filepath.Join(fs.uploadDir, req.Target))
	default:
		return fmt.Errorf("unsupported method: %s", req.Method)
	}
}

func (fs *fileSystemHandler) Filelist(req *sftp.Request) (sftp.ListerAt, error) {
	filePath := filepath.Join(fs.uploadDir, req.Filepath)
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return &fileInfoList{[]os.FileInfo{info}}, nil
	}
	files, err := os.ReadDir(filePath)
	if err != nil {
		return nil, err
	}
	infos := make([]os.FileInfo, len(files))
	for i, file := range files {
		infos[i], _ = file.Info()
	}
	return &fileInfoList{infos}, nil
}

type fileInfoList struct {
	infos []os.FileInfo
}

func (l *fileInfoList) ListAt(f []os.FileInfo, offset int64) (int, error) {
	if offset >= int64(len(l.infos)) {
		return 0, io.EOF
	}
	n := copy(f, l.infos[offset:])
	if n < len(f) {
		return n, io.EOF
	}
	return n, nil
}

type fileWriter struct {
	file     *os.File
	handler  *fileSystemHandler
	filepath string
}

func (fw *fileWriter) WriteAt(p []byte, off int64) (int, error) {
	return fw.file.WriteAt(p, off)
}

func (fw *fileWriter) Close() error {
	err := fw.file.Close()

	fw.handler.mu.Lock()
	filePath, exists := fw.handler.fileMap[fw.filepath]
	if exists {
		delete(fw.handler.fileMap, fw.filepath)
	}
	fw.handler.mu.Unlock()

	if exists && filepath.Ext(filePath) == ".csv" {
		printCSVContent(filePath)
	}

	return err
}

func printCSVContent(filePath string) {
	log.Printf("[server] received CSV file: %s", filePath)
	fmt.Println("=== CSV File Content ===")
	fmt.Printf("File: %s\n", filePath)
	fmt.Println("---")

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("[server] failed to open file %s: %v", filePath, err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lineNum := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("[server] CSV read error: %v", err)
			break
		}
		lineNum++
		fmt.Printf("Line %d: %v\n", lineNum, record)
	}
	fmt.Println("=== End of CSV Content ===")
}
