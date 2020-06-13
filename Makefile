DC = docker-compose
CURRENT_DIR = $(shell pwd)
API = profile

proto:
	docker run --rm -v $(CURRENT_DIR)/pb:/pb -v $(CURRENT_DIR)/schema:/proto ezio1119/protoc \
	-I/proto \
	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:/pb \
	--validate_out="lang=go:/pb" \
	profile.proto

cli:
	docker run --rm --name grpc_cli --net=api-gateway_default namely/grpc-cli \
	call $(API):50051 $(API)_grpc.ProfileService.$(m) "$(q)" $(o)

sqldoc:
	docker run --rm --net=api-gateway_default -v $(CURRENT_DIR)/db:/work ezio1119/tbls \
	doc -f -t svg mysql://root:password@${API}-db:3306/${API}_DB ./

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
	docker logs -f --tail 100 fishapp-profile_profile_1