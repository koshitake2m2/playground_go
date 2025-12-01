# sftp_sample

## Setup

SSH鍵ペアを作成する。

- サーバーSSH鍵ペア
- クライアントSSH鍵ペア

```bash
make
```

## 実行

### Server

```bash
make run-server
```

※ 起動したらmacで `Do you want the application “sftp-server” to accept incoming network connections?` というポップアップが表示されるので、`Deny` をクリックしてください。

サーバーは `:12222` で待機し、SFTP接続を受け付けます。
CSVファイルを受信すると、内容を標準出力に表示し、`uploads` ディレクトリに保存します。

### Client

```bash
make run-client
```

クライアントは `sample.csv` ファイルをサーバーに送信します。
