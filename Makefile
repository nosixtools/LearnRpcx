.PHONY: hello world  client client
hello :
	cd hello_server && go run hello_server.go
world :
	cd world_server && go run world_server.go
client :
	cd client && go run client.go