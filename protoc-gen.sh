protoc -Igowing/proto --go_out=. types.proto
protoc -Igowing/proto --gotag_out=. types.proto
protoc -Irpc/proto:gowing/proto --go_out=. --go-triple_out=. ri-service-provider.proto
protoc -Irpc/proto:gowing/proto --gotag_out=. ri-service-provider.proto
