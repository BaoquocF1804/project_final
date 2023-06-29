sqlc:
	sqlc generate

proto_user:
	rm -f pb_user/*.go
	protoc --proto_path=proto_user --go_out=pb_user --go_opt=paths=source_relative \
        --go-grpc_out=pb_user --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb_user --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=pb_user/openapiv2_user \
        proto_user/*.proto

proto_connect:
	rm -f pb_connect/*.go
	protoc --proto_path=proto_connect --go_out=pb_connect --go_opt=paths=source_relative \
        --go-grpc_out=pb_connect --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb_connect --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=pb_connect/openapiv2_connect \
        proto_connect/*.proto

proto_account:
	rm -f pb_account/*.go
	protoc --proto_path=proto_account --go_out=pb_account --go_opt=paths=source_relative \
        --go-grpc_out=pb_account --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb_account --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=pb_account/openapiv2_account \
        proto_account/*.proto

server_user:
	go run service_user/server.go

server_account:
	go run service_account/server.go

server_connect:
	go run service_connect/server.go
.PHONY: sqlc proto_user proto_connect proto_account server_user server_account server_connect