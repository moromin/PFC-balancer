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
	make db
	make user
	make auth
	make food
	make recipe
	docker compose up

gen-protoc:
# cd ${MAKE_PATH} && $(PROTOC_GRPC)
	$(PROTOC_GRPC)

gen-gateway:
	$(PROTOC_GATEWAY)

db:
	docker build -t moromin/pfc-balancer/db:latest --file platform/db/Dockerfile .

user:
	docker build -t moromin/pfc-balancer/user:latest --file services/user/Dockerfile .

gateway:
	cd services/gateway && go run main.go

auth:
	docker build -t moromin/pfc-balancer/auth:latest --file services/auth/Dockerfile .

food:
	docker build -t moromin/pfc-balancer/food:latest --file services/food/Dockerfile .

recipe:
	docker build -t moromin/pfc-balancer/recipe:latest --file services/recipe/Dockerfile .


.PHONY: gen-gateway gen-protoc
