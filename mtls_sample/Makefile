.PHONY: all
all: certs

.PHONY: certs
certs: root-ca server-cert client-cert

.PHONY: root-ca
root-ca:
	openssl genrsa -out ca.key 4096
	openssl req -x509 -new -sha256 -days 3650 \
		-key ca.key -out ca.crt \
		-subj "/CN=Local Dev CA"

.PHONY: server-cert
server-cert:
	openssl genrsa -out server.key 2048
	openssl req -new -key server.key -out server.csr \
		-subj "/CN=localhost"
	openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key \
		-CAcreateserial -out server.crt -days 825 -sha256 \
		-extfile server.ext

.PHONY: client-cert
client-cert:
	openssl genrsa -out client.key 2048
	openssl req -new -key client.key -out client.csr \
		-subj "/CN=Local Client"
	openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key \
		-CAcreateserial -out client.crt -days 825 -sha256 \
		-extfile client.ext

.PHONY: server-run
server-run:
	cd server && TLS_CERT=../server.crt TLS_KEY=../server.key TLS_CA=../ca.crt go run .

.PHONY: client-run
client-run:
	cd client && TLS_CLIENT_CERT=../client.crt TLS_CLIENT_KEY=../client.key TLS_CA=../ca.crt go run .

.PHONY: clean
clean:
	rm -f ca.key ca.crt ca.srl \
		server.key server.csr server.crt \
		client.key client.csr client.crt
