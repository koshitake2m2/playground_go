# mtls_spire_sample

Go client/server pair that performs mutual TLS using SPIFFE/SPIRE.

## Install

```bash
git clone --single-branch --branch v1.13.3 https://github.com/spiffe/spire.git
cd spire
go build ./cmd/spire-server 
go build ./cmd/spire-agent
mv spire-server spire-agent path/to/bin/
```

## Run

### Start SPIRE Server

```bash
spire-server run -config spire/server/server.conf
```

### Start SPIRE Agent

```bash
spire-server bundle show -format pem > spire/data/bundle.pem
JOIN_TOKEN=$(spire-server token generate -spiffeID spiffe://example.com/agent | awk '{print $2}')
mkdir -p ./spire/data/agent-keys
spire-agent run -config spire/agent/agent.conf -joinToken $JOIN_TOKEN
```

```bash

# spire-server agent list
# AGENT_SPIFFE_ID=spiffe://example.com/spire/agent/join_token/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
AGENT_SPIFFE_ID=$(spire-server agent list | grep SPIFFE | awk '{print $4}')

spire-server entry create \
  -parentID $AGENT_SPIFFE_ID \
  -spiffeID spiffe://example.com/server \
  -selector unix:path:$PWD/server/server
spire-server entry create \
  -parentID $AGENT_SPIFFE_ID \
  -spiffeID spiffe://example.com/client \
  -selector unix:path:$PWD/client/client
```

### Run Server

※ `go run .` では動作しない. selectorでunix:pathでバイナリを指定しているため. clientも同様.

```bash
cd server && go build . && cd .. &&PROJECT_ROOT=$PWD ./server/server
```

### Run Client

```bash
cd client && go build . && cd .. && PROJECT_ROOT=$PWD ./client/client
```

## Tips

```bash
# すでにあるunix:pathエントリを確認・削除する
spire-server entry show
spire-server entry delete -entryID xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx

SERVER_ENTRY_ID=$(spire-server entry show -output json \
  | jq -r '.entries[]
    | select(.selectors[]? | (.type == "unix" and .value == "path:'$PWD'/server/server"))
    | .id')
CLIENT_ENTRY_ID=$(spire-server entry show -output json \
  | jq -r '.entries[]
    | select(.selectors[]? | (.type == "unix" and .value == "path:'$PWD'/client/client"))
    | .id')

spire-server entry delete -entryID $SERVER_ENTRY_ID
spire-server entry delete -entryID $CLIENT_ENTRY_ID
```
