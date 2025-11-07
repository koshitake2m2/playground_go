# mtls_sample

## Setup

### Install

```bash
brew install openssl
```

### ルートCA

```bash
# CA秘密鍵
openssl genrsa -out ca.key 4096

# CA証明書（自己署名）
openssl req -x509 -new -sha256 -days 3650 \
  -key ca.key -out ca.crt \
  -subj "/CN=Local Dev CA"
```

### サーバー証明書

```bash
# サーバ鍵とCSR
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr \
  -subj "/CN=localhost"

# 拡張設定（SAN と EKU）
cat > server.ext <<'EOF'
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = localhost
EOF

# CAで署名
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key \
  -CAcreateserial -out server.crt -days 825 -sha256 \
  -extfile server.ext
```

### クライアント証明書

```bash
# クライアント鍵とCSR
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr \
  -subj "/CN=Local Client"

# 拡張設定（clientAuth）
cat > client.ext <<'EOF'
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth
EOF

# CAで署名
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key \
  -CAcreateserial -out client.crt -days 825 -sha256 \
  -extfile client.ext
```

### 動作確認

```bash
# 簡易的なTLSサーバ実行（8443で待受）
openssl s_server -accept 8443 -www \
  -cert server.crt -key server.key \
  -CAfile ca.crt -Verify 1
```

```bash
# サーバ証明書の検証に使うCAを指定し、
# クライアント証明書と鍵を提示する
curl https://localhost:8443/ \
  --cacert ca.crt \
  --cert client.crt \
  --key client.key -v
```

```bash
# 証明書の内容確認と検証
openssl verify -CAfile ca.crt server.crt client.crt
```

## 実行

### Server

```bash
cd server
go run .
```

### Client

```bash
cd client
go run.
```

```bash
curl https://localhost:8443 \
  --cert client.crt \
  --key client.key \
  --cacert ca.crt
```
