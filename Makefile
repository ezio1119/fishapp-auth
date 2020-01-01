gen-proto:
	protoc \
  -I . \
  -I ${GOPATH}/src \
  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:. \
  --validate_out="lang=go:." \
	./src/controllers/profile_grpc/profile.proto

user:
	docker-compose exec user sh