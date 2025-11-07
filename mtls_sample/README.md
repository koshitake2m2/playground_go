# mtls_sample

## Install

```bash
brew install openssl
```

## Setup

以下の3種類の証明書を作成する。

- ルートCA
- サーバー証明書
- クライアント証明書

```bash
make
```

## 実行

### Server

```bash
make run-server
```

### Client

```bash
make run-client
make run-client2
```

### 簡易動作確認

```bash
# 簡易的なTLSサーバ実行
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
