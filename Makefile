PROTOC_GRPC = protoc  --go_out=. \
    		--go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,\
			proto/*.proto

SERVICE = user

protoc:
	cd services/${SERVICE} && $(PROTOC_GRPC)

.PHONY:
