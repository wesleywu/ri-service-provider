package test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	httpclient "github.com/castbox/go-guru/pkg/client/http"
	"github.com/castbox/go-guru/pkg/infra/appinfo"
	"github.com/castbox/go-guru/pkg/infra/logger"
	"github.com/castbox/go-guru/pkg/infra/otlp"
	"github.com/castbox/go-guru/pkg/util/codec"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/enum"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
)

const (
	otlpHttpEndpoint   = "34.120.15.175"
	otlpBasicAuthToken = "bWVuZ3llLnd1QGNhc3Rib3guZm06VUxGV3BnQ3ZJNUdDWDhCTA=="
	otlpInsecure       = true
)

var (
	ctx           = context.Background()
	baseUrl       = "http://localhost:8200/v1/video-collection"
	log           = logger.NewConsoleLogger(zerolog.InfoLevel)
	logHelper     = logger.NewLoggerHelper(ctx, log)
	appMetadata   = newAppMetadata()
	otlpConfigs   = otlp.NewHttpTpConfigs(appMetadata, otlpHttpEndpoint, otlpBasicAuthToken, otlpInsecure)
	tp, _         = otlp.NewTracerProvider(ctx, appMetadata, otlpConfigs, logHelper)
	client        = httpclient.NewHttpClient(httpclient.NewConfigs(), tp, logHelper)
	commonHeaders = &http.Header{
		"Content-Type": []string{"application/json"},
	}
	jsonCodec = encoding.GetCodec(codec.Name)
)

func newAppMetadata() *appinfo.AppMetadata {
	hostname, _ := os.Hostname()
	return &appinfo.AppMetadata{
		AppName:    "http_client_test",
		AppVersion: "v0.0.1",
		HostName:   hostname,
	}
}

func TestVideoCollectionRepo_All(t *testing.T) {
	var (
		url            string
		data           string
		httpRes        *httpclient.Response
		createRes      *proto.VideoCollectionCreateRes
		upsertRes      *proto.VideoCollectionUpsertRes
		updateRes      *proto.VideoCollectionUpdateRes
		oneRes         *proto.VideoCollectionOneRes
		countRes       *proto.VideoCollectionCountRes
		getRes         *proto.VideoCollectionGetRes
		listRes        *proto.VideoCollectionListRes
		deleteRes      *proto.VideoCollectionDeleteRes
		deleteMultiRes *proto.VideoCollectionDeleteMultiRes
		insertedID1    string
		insertedID2    = "qiihWlTCtVz72T9znB9"
		err            error
	)
	tracer := tp.Tracer("ri-service-provider-test")
	_, span := tracer.Start(ctx, "http_client_test.TestVideoCollectionRepo_All")
	defer span.End()

	// test Delete 删除1条可能之前存在的记录
	url = baseUrl + "/" + insertedID2
	httpRes, err = client.DeleteWithHeaders(url, commonHeaders, "", 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	deleteRes = (*proto.VideoCollectionDeleteRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &deleteRes)
	assert.NoError(t, err)

	dateStarted := time.Now()

	// test Create 会插入一条记录
	url = baseUrl
	data = fmt.Sprintf(`{
			"name": "测试视频集01",
			"contentType": "%s",
			"filterType": "%s",
			"count": 1234,
			"isOnline": false
		}`, enum.ContentType_name[int32(enum.ContentType_PortraitVideo)], enum.FilterType_name[int32(enum.FilterType_Manual)])
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	createRes = (*proto.VideoCollectionCreateRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &createRes)
	//err = jsonCodec.Unmarshal(httpRes.Body, &createRes)
	assert.NoError(t, err)
	assert.NotNil(t, createRes.InsertedID)
	assert.Equal(t, int64(1), createRes.InsertedCount)
	insertedID1 = *createRes.InsertedID

	// test Upsert 会插入第二条记录
	url = baseUrl + "/" + insertedID2
	data = fmt.Sprintf(`{
			"name": "测试视频集02",
			"contentType": "%s",
			"filterType": "%s",
			"count": 2345,
			"isOnline": true
		}`, enum.ContentType_name[int32(enum.ContentType_LandscapeVideo)], enum.FilterType_name[int32(enum.FilterType_Ruled)])
	httpRes, err = client.PutWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	upsertRes = (*proto.VideoCollectionUpsertRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &upsertRes)
	assert.NoError(t, err)
	assert.NotNil(t, upsertRes.UpsertedID)
	assert.Equal(t, insertedID2, *upsertRes.UpsertedID)
	assert.Equal(t, int64(1), upsertRes.UpsertedCount)

	// test One 第1次，命中1条记录
	url = baseUrl + "/one"
	data = fmt.Sprintf(`{
				"name": {
					"@type":"google.protobuf.StringValue",
					"value":"测试视频集01"
				},
				"contentType": {
					"@type":"goguru.types.StringSlice",
					"value":["LandscapeVideo", "PortraitVideo"]
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				},
				"createdAt": {
					"@type":"goguru.orm.Condition",
					"operator": "GTE",
					"value": {
						"@type":"google.protobuf.Timestamp",
						"value":"%s"
					}
				}
			}`, dateStarted.Format(time.RFC3339Nano))
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	oneRes = (*proto.VideoCollectionOneRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &oneRes)
	assert.NoError(t, err)
	assert.Equal(t, true, oneRes.Found)
	assert.Equal(t, int32(1234), *oneRes.Item.Count)

	// test One 第2次，无命中记录
	data = fmt.Sprintf(`{
				"name": {
					"@type":"google.protobuf.StringValue",
					"value":"测试视频集01"
				},
				"contentType": {
					"@type":"google.protobuf.StringValue",
					"value":"LandscapeVideo"
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				},
				"createdAt": {
					"@type":"goguru.orm.Condition",
					"operator": "GTE",
					"value": {
						"@type":"google.protobuf.Timestamp",
						"value":"%s"
					}
				}
			}`, dateStarted.Format(time.RFC3339Nano))
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	oneRes = (*proto.VideoCollectionOneRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &oneRes)
	assert.NoError(t, err)
	assert.Equal(t, false, oneRes.Found)

	// test Count 第1次，共2条满足条件的记录
	url = baseUrl + "/count"
	data = fmt.Sprintf(`{
				"contentType": {
					"@type":"goguru.types.StringSlice",
					"value":["LandscapeVideo", "PortraitVideo"]
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				},
				"createdAt": {
					"@type":"goguru.orm.Condition",
					"operator": "GTE",
					"value": {
						"@type":"google.protobuf.Timestamp",
						"value":"%s"
					}
				}
			}`, dateStarted.Format(time.RFC3339Nano))
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	countRes = (*proto.VideoCollectionCountRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &countRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), countRes.TotalElements)

	// test Count 第2次，共1条满足条件的记录
	url = baseUrl + "/count"
	data = fmt.Sprintf(`{
				"name": {
					"@type":"google.protobuf.StringValue",
					"value":"测试视频集01"
				},
				"createdAt": {
					"@type":"goguru.orm.Condition",
					"operator": "GTE",
					"value": {
						"@type":"google.protobuf.Timestamp",
						"value":"%s"
					}
				}
			}`, dateStarted.Format(time.RFC3339Nano))
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	countRes = (*proto.VideoCollectionCountRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &countRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), countRes.TotalElements)

	// test Count 第3次，共2条满足条件的记录
	url = baseUrl + "/count"
	nextDate := dateStarted.AddDate(0, 0, 1)
	data = fmt.Sprintf(`{
				"createdAt": {
					"@type":"goguru.orm.Condition",
					"multi": "Between",
					"value": {
						"@type":"goguru.types.TimestampSlice",
						"value": ["%s","%s"]
					}
				}
			}`, dateStarted.Format(time.RFC3339Nano), nextDate.Format(time.RFC3339Nano))
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	countRes = (*proto.VideoCollectionCountRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &countRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), countRes.TotalElements)

	// test List 第1次，返回第2页，每页1条记录，当页有1条记录，为满足条件的第2条记录，其 Name 为 "TemplateName456"
	url = baseUrl + "/list"
	data = fmt.Sprintf(`{
				"contentType": {
					"@type":"goguru.types.StringSlice",
					"value":["LandscapeVideo", "PortraitVideo"]
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				},
				"createdAt": {
					"@type":"goguru.orm.Condition",
					"operator": "GTE",
					"value": {
						"@type":"google.protobuf.Timestamp",
						"value":"%s"
					}
				},
				"pageRequest": {
					"number": 2,
					"size": 1,
					"sorts": [{
						"property": "name",
						"direction": "Asc"
					}]
				}
			}`, dateStarted.Format(time.RFC3339Nano))
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	listRes = (*proto.VideoCollectionListRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &listRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), listRes.PageInfo.Number)
	assert.Equal(t, int64(2), listRes.PageInfo.TotalElements)
	assert.Equal(t, int64(2), listRes.PageInfo.TotalPages)
	assert.Equal(t, int64(1), listRes.PageInfo.NumberOfElements)
	assert.Equal(t, false, listRes.PageInfo.First)
	assert.Equal(t, true, listRes.PageInfo.Last)

	// test Get 返回第一条记录
	url = baseUrl + "/" + insertedID1
	httpRes, err = client.GetWithHeaders(url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	getRes = (*proto.VideoCollectionGetRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &getRes)
	assert.NotNil(t, getRes.Name)
	assert.Equal(t, "测试视频集01", *getRes.Name)

	// test Update 修改第一条记录
	url = baseUrl + "/" + insertedID1
	data = fmt.Sprintf(`{
			"name": "测试视频集03",
			"count": 3456,
			"isOnline": false
		}`)
	httpRes, err = client.PatchWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	updateRes = (*proto.VideoCollectionUpdateRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &updateRes)
	assert.Equal(t, int64(1), updateRes.ModifiedCount)

	// test Get 再次验证第一条记录
	url = baseUrl + "/" + insertedID1
	httpRes, err = client.GetWithHeaders(url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	getRes = (*proto.VideoCollectionGetRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &getRes)
	assert.NotNil(t, getRes.Name)
	assert.NotNil(t, getRes.Count)
	assert.NotNil(t, getRes.IsOnline)
	assert.Equal(t, "测试视频集03", *getRes.Name)
	assert.Equal(t, int32(3456), *getRes.Count)
	assert.Equal(t, false, *getRes.IsOnline)

	// test Upsert 修改第一条记录
	url = baseUrl + "/" + insertedID1
	data = fmt.Sprintf(`{
			"name": "测试视频集04",
			"count": 4567,
			"isOnline": true
		}`)
	httpRes, err = client.PutWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	updateRes = (*proto.VideoCollectionUpdateRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &updateRes)
	assert.Equal(t, int64(1), updateRes.ModifiedCount)

	// test Get 再次验证第一条记录
	url = baseUrl + "/" + insertedID1
	httpRes, err = client.GetWithHeaders(url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	getRes = (*proto.VideoCollectionGetRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &getRes)
	assert.NotNil(t, getRes.Name)
	assert.NotNil(t, getRes.Count)
	assert.NotNil(t, getRes.IsOnline)
	assert.Equal(t, "测试视频集04", *getRes.Name)
	assert.Equal(t, int32(4567), *getRes.Count)
	assert.Equal(t, true, *getRes.IsOnline)

	// test DeleteMulti 删除2条记录
	url = baseUrl + "/delete"
	data = fmt.Sprintf(`{
				"id": {
					"@type":"goguru.types.StringSlice",
					"value": ["%s","%s"]
				}
			}`, insertedID1, insertedID2)
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	deleteMultiRes = (*proto.VideoCollectionDeleteMultiRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &deleteMultiRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), deleteMultiRes.DeletedCount)

	// test Delete 删除0条记录，因为之前的 deleteMulti 已经删除过了
	url = baseUrl + "/" + insertedID1
	httpRes, err = client.DeleteWithHeaders(url, commonHeaders, "", 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	deleteRes = (*proto.VideoCollectionDeleteRes)(nil)
	err = jsonCodec.Unmarshal(httpRes.Body, &deleteRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), deleteRes.DeletedCount)
}
