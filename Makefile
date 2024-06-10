PROJECT = playable-cms
IMAGE = us-central1-docker.pkg.dev/infra-387702/saas-guru/$(PROJECT)/$(APP_NAME)
TAG = 0.0.8

.PHONY: deploy-prod
deploy-prod:
	$(MAKE) deploy APP_ENV=prod NAMESPACE=saas-guru

.PHONY: configmap-prod
configmap-prod:
	$(MAKE) configmap APP_ENV=prod NAMESPACE=saas-guru

.PHONY: wire
wire:
	cd app/videocollection/service/cmd/server && wire

.PHONY: proto-gen
proto-gen:
	buf mod update
	buf generate --template buf.gen-model.yaml app/videocollection/service/enum/video_collection_enum.proto
	buf generate --template buf.gen-model.yaml app/videocollection/service/proto/video_collection.proto
	#buf generate --template buf.gen-model.yaml app/videocollection/service/proto/video_collection_repo.proto
	buf generate --template buf.gen-tag.yaml app/videocollection/service/proto/video_collection.proto
	#buf generate --template buf.gen-tag.yaml app/videocollection/service/proto/video_collection_repo.proto
	buf generate --template buf.gen.yaml api/videocollection/service/v1/video_collection_api.proto
