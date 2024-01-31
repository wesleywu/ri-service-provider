proto-gen:
	buf generate -o api proto/video_collection/video_collection.proto
	cd api && buf generate --template ../buf.gen-tag.yaml ../proto/video_collection/video_collection.proto
	buf generate --template buf.gen-error.yaml -o api proto/errors/errors.proto

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: wire
wire:
	cd provider/cmd/server && wire