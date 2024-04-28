.PHONY: echo_server
echo_server:
	go run echo_server/server/main.go

.PHONY: echo_client
echo_client:
	go run echo_server/client/main.go