DC = docker-compose
CURRENT_DIR = $(shell pwd)
API = profile

proto:
	docker run --rm -w $(CURRENT_DIR) \
	-v $(CURRENT_DIR)/schema/$(API):/schema \
	-v $(CURRENT_DIR)/src/controllers/$(API)_grpc:$(CURRENT_DIR) \
	thethingsindustries/protoc \
	-I/schema \
	-I/usr/include/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:. \
	--validate_out="lang=go:." \
	--doc_out=markdown,README.md:/schema \
	$(API).proto

cli:
	docker run --rm --name grpc_cli --net=api-gateway_default namely/grpc-cli \
	call $(API):50051 $(API)_grpc.ProfileService.$(m) $(q) $(o)

migrate:
	docker run --rm -it --name migrate --net=api-gateway_default \
	-v $(CURRENT_DIR)/db/sql:/sql migrate/migrate:latest \
	-path /sql/ -database "mysql://root:password@tcp($(API)-db:3306)/$(API)_DB" up

up:
	$(DC) up -d

ps:
	$(DC) ps

build:
	$(DC) build

down:
	$(DC) down

exec:
	$(DC) exec $(API) sh

logs:
	$(DC) logs -f --tail 100 $(API)

redis:
	$(DC) exec $(API)-kvs sh