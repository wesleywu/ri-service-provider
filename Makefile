.PHONY: wire
wire:
	cd app/videocollection/service/cmd/server && wire
	cd test/rpc_client && wire

.PHONY: proto-gen
proto-gen:
	buf mod update
	buf generate --template buf.gen-model.yaml app/videocollection/service/enum/video_collection_enum.proto
	buf generate --template buf.gen-model.yaml app/videocollection/service/proto/video_collection.proto
	buf generate --template buf.gen-tag.yaml   app/videocollection/service/proto/video_collection.proto
	buf generate --template buf.gen.yaml       api/videocollection/service/v1/video_collection_api.proto
