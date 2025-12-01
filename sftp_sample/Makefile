.PHONY: all
all: ssh-keys

.PHONY: ssh-keys
ssh-keys: server-key client-key

.PHONY: server-key
server-key:
	# サーバーSSH秘密鍵
	ssh-keygen -t rsa -b 4096 -f server_key -N "" -C "sftp-server"
	# サーバーSSH公開鍵
	ssh-keygen -y -f server_key > server_key.pub

.PHONY: client-key
client-key:
	# クライアントSSH秘密鍵
	ssh-keygen -t rsa -b 4096 -f client_key -N "" -C "sftp-client"
	# クライアントSSH公開鍵
	ssh-keygen -y -f client_key > client_key.pub

.PHONY: run-server
run-server:
	cd server && SSH_HOST_KEY=../server_key SSH_AUTHORIZED_KEY=../client_key.pub go run .

.PHONY: run-client
run-client:
	cd client && SSH_PRIVATE_KEY=../client_key SSH_HOST_KEY=../server_key.pub go run .

.PHONY: clean
clean:
	rm -rf server_key server_key.pub \
		client_key client_key.pub \
		server/uploads

