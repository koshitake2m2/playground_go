# mtls_sample

## Setup

### Install

```bash
brew install openssl
```

### 証明書作成

```bash
make
```

以下の3種類の証明書を作成する。

- ルートCA
- サーバー証明書
- クライアント証明書

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
  --key client.key
```

### Tips

```bash
# 証明書の内容確認と検証
openssl verify -CAfile ca.crt server.crt client.crt
```

## 実行

### Server

```bash
make server-run
```

### Client

```bash
make client-run
```

```bash
curl https://localhost:8443 \
  --cert client.crt \
  --key client.key \
  --cacert ca.crt
```
