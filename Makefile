proto: 
	# rm api/shorter/*.go
	protoc -I api/proto --go_out=plugins=grpc:api/shorter api/proto/shorter.proto

run: 
	go run cmd/server/main.go
