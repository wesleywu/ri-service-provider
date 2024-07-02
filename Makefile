.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/castbox/protoc-gen-go-guru/cmd/protoc-gen-go-guru@latest

.PHONY: wire
wire:
	cd app/episode/service/cmd/server && wire
	cd test/rpc_client && wire

.PHONY: proto-gen
proto-gen:
	buf mod update
	buf generate --template buf.gen-model.yaml app/episode/service/enum/episode_enum.proto
	buf generate --template buf.gen-model.yaml app/episode/service/proto/episode.proto
	buf generate --template buf.gen-tag.yaml   app/episode/service/proto/episode.proto
	buf generate --template buf.gen.yaml       api/episode/service/v1/episode_api.proto
