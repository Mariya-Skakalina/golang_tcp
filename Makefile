.PHONY: echo_server
echo_server:
	go run echo_server/server/main.go

.PHONY: echo_client
echo_client:
	go run echo_server/client/main.go

.PHONY: chat_server
chat_server:
	go run chat_server/server/main.go

.PHONY: chat_client
chat_client:
	go run chat_server/client/main.go