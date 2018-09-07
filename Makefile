gen-proto:
	protoc --go_out=. protos/*.proto

run:
	go run cli/main.go node -m -p 13001