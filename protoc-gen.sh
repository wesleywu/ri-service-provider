protoc -Iapp/video_collection/proto:gowing/proto --go_out=. --go-triple_out=. video_collection.proto
protoc -Iapp/video_collection/proto:gowing/proto --gotag_out=. video_collection.proto
