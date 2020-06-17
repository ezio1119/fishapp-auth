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
	docker run --rm --name grpc_cli --net=fishapp-net znly/grpc_cli \
	call $(API):50051 $(API).ProfileService.$(m) "$(q)"

sqldoc:
	docker run --rm --net=fishapp-net -v $(CURRENT_DIR)/db:/work ezio1119/tbls \
	doc -f -t svg mysql://root:password@${API}-db:3306/${API}_DB ./

migrate:
	docker run --rm -it --name migrate --net=fishapp-net \
	-v $(CURRENT_DIR)/db/sql:/sql migrate/migrate:latest \
	-path /sql/ -database "mysql://root:password@tcp($(API)-db:3306)/$(API)_DB" ${a}

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