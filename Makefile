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


gen-protoc:
# cd ${MAKE_PATH} && $(PROTOC_GRPC)
	$(PROTOC_GRPC)

gen-gateway:
	$(PROTOC_GATEWAY)

db:
	cd platform/db && go run main.go

gateway:
	cd services/gateway && go run main.go

auth:
	cd services/auth && go run main.go

user:
	cd services/user && go run main.go

menu:
	cd services/menu && go run main.go

food:
	cd services/food && go run main.go

.PHONY: gen-gateway gen-protoc
