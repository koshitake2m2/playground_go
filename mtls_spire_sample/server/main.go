package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func main() {
	ctx := context.Background()

	// Workload API ソース（証明書は自動取得＆ローテ）
	source, err := workloadapi.NewX509Source(ctx,
		workloadapi.WithClientOptions(workloadapi.WithAddr("unix://"+os.Getenv("PROJECT_ROOT")+"/spire/workload-api/spire-agent.sock")),
	)

	if err != nil {
		log.Fatalf("create x509 source: %v", err)
	}
	defer source.Close()

	// クライアントのSPIFFE IDを許可（認可）
	clientID := spiffeid.RequireFromString("spiffe://example.com/client")
	tlsCfg := tlsconfig.MTLSServerConfig(source, source, tlsconfig.AuthorizeID(clientID))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{"message":"hello world"}`))
	})

	srv := &http.Server{
		Addr:      ":8443",
		Handler:   mux,
		TLSConfig: tlsCfg,
	}

	log.Println("[server] listening on :8443 (SPIFFE mTLS)")
	if err := srv.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}
