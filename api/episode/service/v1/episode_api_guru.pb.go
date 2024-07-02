// Code generated by protoc-gen-go-guru. DO NOT EDIT.
// versions:
//  protoc-gen-go-guru v0.2.14
// source: api/episode/service/v1/episode_api.proto

package v1

import (
	cache "github.com/castbox/go-guru/pkg/goguru/cache"
	rpcinfo "github.com/castbox/go-guru/pkg/server/rpcinfo"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func RegisterEpisodeGuruServer(srv EpisodeServer) error {
	var err error
	rpcinfo.RegisterServiceDesc(&Episode_ServiceDesc, srv)
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "Count", &cache.CacheRule{
		Cachable: false,
		Name:     "Episode",
		Ttl:      "30s",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "One", &cache.CacheRule{
		Cachable: false,
		Name:     "Episode",
		Ttl:      "30s",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "List", &cache.CacheRule{
		Cachable: false,
		Name:     "Episode",
		Ttl:      "30s",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "Get", &cache.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "Create", &cache.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "Update", &cache.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "Upsert", &cache.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "Delete", &cache.CacheRule{
		Cachable: false,
		Name:     "",
		Ttl:      "",
		Key:      "",
	})
	if err != nil {
		return err
	}
	err = rpcinfo.RegisterMethodDesc("v1.Episode", "DeleteMulti", &cache.CacheRule{
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
