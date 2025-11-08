package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
)

func main() {
	ctx := context.Background()

	source, err := workloadapi.NewX509Source(ctx,
		workloadapi.WithClientOptions(workloadapi.WithAddr("unix://"+os.Getenv("PROJECT_ROOT")+"/spire/workload-api/spire-agent.sock")),
	)
	if err != nil {
		log.Fatalf("create x509 source: %v", err)
	}
	defer source.Close()

	serverID := spiffeid.RequireFromString("spiffe://example.com/server")
	tlsCfg := tlsconfig.MTLSClientConfig(source, source, tlsconfig.AuthorizeID(serverID))

	client := &http.Client{
		Transport: &http.Transport{TLSClientConfig: tlsCfg},
	}

	resp, err := client.Get("https://localhost:8443/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Printf("status=%s body=%s", resp.Status, string(body))
}
