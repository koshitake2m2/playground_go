package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	certFile := getenv("TLS_CLIENT_CERT", "../client.crt")
	keyFile := getenv("TLS_CLIENT_KEY", "../client.key")
	caFile := getenv("TLS_CA", "../ca.crt")
	url := getenv("SERVER_URL", "https://localhost:9443/")
	retries, _ := strconv.Atoi(getenv("RETRIES", "20"))
	sleepSecs, _ := strconv.Atoi(getenv("SLEEP_SECS", "1"))

	// クライアント証明書
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("load client cert: %v", err)
	}

	// サーバ検証用 CA
	caPEM, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("read CA: %v", err)
	}
	cp := x509.NewCertPool()
	if ok := cp.AppendCertsFromPEM(caPEM); !ok {
		log.Fatal("append CA failed")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			RootCAs:      cp,
		},
	}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}

	var resp *http.Response
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err = client.Do(req)
		cancel()
		if err == nil {
			break
		}
		log.Printf("[client] retrying... (%d/%d) err=%v", i+1, retries, err)
		time.Sleep(time.Duration(sleepSecs) * time.Second)
	}
	if err != nil {
		log.Fatalf("[client] request failed: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("[client] status=%s\n%s", resp.Status, string(body))
}
