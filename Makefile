proto-gen:
	buf generate -o provider/api proto/video_collection/video_collection.proto
	cd provider/api && buf generate --template ../../buf.gen-tag.yaml ../../proto/video_collection/video_collection.proto

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: wire
wire:
	cd provider/cmd/server && wire