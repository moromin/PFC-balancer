PROTOC_GRPC = protoc  --go_out=. \
    		--go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,\
			proto/*.proto

MAKE_PATH = platform/db/db

protoc:
	cd ${MAKE_PATH} && $(PROTOC_GRPC)

run:
	cd ${MAKE_PATH} && go run main.go

.PHONY:
