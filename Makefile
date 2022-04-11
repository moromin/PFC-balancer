PROTOC_GRPC = protoc  --go_out=. \
    		--go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,\
			proto/*.proto

PROTOC_GATEWAY = protoc -I . \
				--grpc-gateway_out . \
				--grpc-gateway_opt logtostderr=true \
				--grpc-gateway_opt paths=source_relative \
				--grpc-gateway_opt generate_unbound_methods=true \
				proto/*.proto

MAKE_PATH = platform/db/db

protoc:
	cd ${MAKE_PATH} && $(PROTOC_GRPC)

gateway:
	cd ${MAKE_PATH} && $(PROTOC_GATEWAY)

run:
	cd ${MAKE_PATH} && go run main.go

.PHONY:
