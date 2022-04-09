PROTOC_GRPC = protoc  --go_out=. \
    		--go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,\
			proto/*.proto

SERVICE = user
RUN_PATH = platform/db/db

protoc:
	cd services/${SERVICE} && $(PROTOC_GRPC)

run:
	cd ${RUN_PATH} && go run main.go

.PHONY:
