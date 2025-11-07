package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

func mustGetenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	certFile := mustGetenv("TLS_CERT", "../server.crt")
	keyFile := mustGetenv("TLS_KEY", "../server.key")
	caFile := mustGetenv("TLS_CA", "../ca.crt")
	addr := mustGetenv("LISTEN_ADDR", ":9443")

	// サーバ証明書
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("load server cert: %v", err)
	}

	// クライアント検証用 CA
	caPEM, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("read CA: %v", err)
	}
	cp := x509.NewCertPool()
	if ok := cp.AppendCertsFromPEM(caPEM); !ok {
		log.Fatal("append CA failed")
	}

	tlsCfg := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    cp,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cn := "(none)"
		if r.TLS != nil && len(r.TLS.PeerCertificates) > 0 {
			cn = r.TLS.PeerCertificates[0].Subject.CommonName
		}
		fmt.Fprintf(w, "mTLS OK. Hello, %s!\n", cn)
	})

	srv := &http.Server{
		Addr:      addr,
		Handler:   mux,
		TLSConfig: tlsCfg,
	}

	log.Printf("[server] listening on %s (mTLS required)", addr)
	if err := srv.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
