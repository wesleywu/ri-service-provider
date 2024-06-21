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
