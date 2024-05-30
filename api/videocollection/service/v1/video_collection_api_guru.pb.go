// Code generated by protoc-gen-go-guru. DO NOT EDIT.
// versions:
// - protoc-gen-go-guru v0.1.36
// - protoc             (unknown)
// source: api/videocollection/service/v1/video_collection_api.proto

package v1

import (
	"github.com/castbox/go-guru/pkg/goguru/annotations"
	"github.com/castbox/go-guru/pkg/server"
)

func RegisterVideoCollectionGuruServer(srv VideoCollectionServer) error {
	var err error
	server.RegisterServiceDesc(&VideoCollection_ServiceDesc, srv)
	err = server.RegisterMethodDesc("v1.VideoCollection", "Count", &annotations.CacheRule{
		Cachable: false,
		Name:     "VideoCollection",
		Ttl:      "30s",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "One", &annotations.CacheRule{
		Cachable: false,
		Name:     "VideoCollection",
		Ttl:      "30s",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "List", &annotations.CacheRule{
		Cachable: false,
		Name:     "VideoCollection",
		Ttl:      "30s",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "Get", &annotations.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "Create", &annotations.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "Update", &annotations.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "Upsert", &annotations.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "Delete", &annotations.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = server.RegisterMethodDesc("v1.VideoCollection", "DeleteMulti", &annotations.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	return nil
}