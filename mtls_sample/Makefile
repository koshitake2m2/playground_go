.PHONY: all
all: certs

.PHONY: certs
certs: root-ca server-cert client-cert client2-cert

.PHONY: root-ca
root-ca:
	# CA秘密鍵
	openssl genrsa -out ca.key 4096
	# CA証明書（自己署名）
	openssl req -x509 -new -sha256 -days 3650 \
		-key ca.key -out ca.crt \
		-subj "/CN=Local Dev CA"

.PHONY: server-cert
server-cert:
	# サーバ鍵とCSR
	openssl genrsa -out server.key 2048
	openssl req -new -key server.key -out server.csr \
		-subj "/CN=localhost"
	# CAで署名（server.extでSAN/EKU指定）
	openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key \
		-CAcreateserial -out server.crt -days 825 -sha256 \
		-extfile server.ext

.PHONY: client-cert
client-cert:
	# クライアント鍵とCSR
	openssl genrsa -out client.key 2048
	openssl req -new -key client.key -out client.csr \
		-subj "/CN=Local Client"
	# CAで署名（client.extでclientAuth指定）
	openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key \
		-CAcreateserial -out client.crt -days 825 -sha256 \
		-extfile client.ext

.PHONY: client2-cert
client2-cert:
	# client2 鍵とCSR
	openssl genrsa -out client2.key 2048
	openssl req -new -key client2.key -out client2.csr \
		-subj "/CN=Local Client 2"
	# CAで署名（client2.extでclientAuth指定）
	openssl x509 -req -in client2.csr -CA ca.crt -CAkey ca.key \
		-CAcreateserial -out client2.crt -days 825 -sha256 \
		-extfile client2.ext

.PHONY: server-run
server-run:
	cd server && TLS_CERT=../server.crt TLS_KEY=../server.key TLS_CA=../ca.crt go run .

.PHONY: client-run
client-run:
	cd client && TLS_CLIENT_CERT=../client.crt TLS_CLIENT_KEY=../client.key TLS_CA=../ca.crt go run .

.PHONY: client2-run
client2-run:
	cd client2 && TLS_CLIENT_CERT=../client2.crt TLS_CLIENT_KEY=../client2.key TLS_CA=../ca.crt go run .

.PHONY: clean
clean:
	rm -f ca.key ca.crt ca.srl \
		server.key server.csr server.crt \
		client.key client.csr client.crt \
		client2.key client2.csr client2.crt
