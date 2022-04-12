MAKE_PATH = .

PROTOC_GRPC = protoc -I . \
			--go_out ./${MAKE_PATH} \
    		--go-grpc_out ./${MAKE_PATH} \
			--go-grpc_opt=paths=source_relative,\
			--go-grpc_opt=require_unimplemented_servers=false,\
			./${MAKE_PATH}/proto/*.proto

PROTOC_GATEWAY = protoc -I . \
				--grpc-gateway_out ./${MAKE_PATH} \
				--grpc-gateway_opt logtostderr=true \
				--grpc-gateway_opt paths=source_relative \
				--grpc-gateway_opt generate_unbound_methods=true \
				./${MAKE_PATH}/proto/*.proto

# TODO: run all services rule
all:
	go run ***

protoc:
# cd ${MAKE_PATH} && $(PROTOC_GRPC)
	$(PROTOC_GRPC)

gateway:
	$(PROTOC_GATEWAY)

run:
	cd ${MAKE_PATH} && go run main.go

.PHONY:
