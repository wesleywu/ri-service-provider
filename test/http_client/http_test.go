package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/castbox/go-guru/pkg/server"
	"github.com/castbox/go-guru/pkg/util/gjson"
	"github.com/castbox/go-guru/pkg/util/httpclient"
	"github.com/castbox/go-guru/pkg/util/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/enum"
	"github.com/wesleywu/ri-service-provider/app/videocollection/service/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	ctx       = context.Background()
	baseUrl   = "http://localhost:8200/v1/video-collection"
	log       = logger.NewConsoleLogger(zerolog.InfoLevel)
	logHelper = logger.NewLoggerHelper(ctx, log)
	client    = httpclient.NewHttpClient(&conf.HttpClient{
		Timeout: durationpb.New(10 * time.Second),
	},
		logHelper)
	commonHeaders = &http.Header{
		"Content-Type":                  []string{"application/json"},
		"x-dubbo-http1.1-dubbo-version": []string{"1.0.0"},
		"x-dubbo-service-protocol":      []string{"triple"},
	}
)

const id = "01186883-7700"

func TestList(t *testing.T) {
	url := baseUrl + "/list"
	resp, err := client.PostWithHeaders(url, commonHeaders, "{}", 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	res := (*proto.VideoCollectionListRes)(nil)
	err = gjson.DecodeTo(resp.Body, res)
	require.NoError(t, err)
	require.Equal(t, int64(3), res.PageInfo.TotalElements)
}

func TestListBySingleValue(t *testing.T) {
	url := baseUrl + "/list"
	data := `{
				"name": {
					"@type":"google.protobuf.StringValue",
					"value":"每日推荐视频"
				},
				"isOnline": {
					"@type":"google.protobuf.BoolValue",
					"value":false
				}
			}`

	resp, err := client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	res := (*proto.VideoCollectionListRes)(nil)
	err = gjson.DecodeTo(resp.Body, res)
	require.NoError(t, err)
	require.Equal(t, int64(1), res.PageInfo.TotalElements)
}

func TestListBySliceValue(t *testing.T) {
	url := baseUrl + "/list"
	data := `{
				"name": {
					"@type":"goguru.types.StringSlice",
					"value":["每日推荐视频","日常推荐视频"]
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				}
			}`
	resp, err := client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.NoError(t, err)
	res := (*proto.VideoCollectionListRes)(nil)
	err = gjson.DecodeTo(resp.Body, res)
	require.NoError(t, err)
	require.Equal(t, int64(2), res.PageInfo.TotalElements)

}

func TestListByCondition(t *testing.T) {
	TestCreate(t)
	url := baseUrl + "/list"
	data := `{
				"name": {
					"@type":"goguru.types.Condition",
					"operator": "Like",
					"wildcard": "StartsWith",
					"value": {
						"@type":"google.protobuf.StringValue",
						"value":"每日"
					}
				},
				"count": {
					"@type":"goguru.types.Condition",
					"operator": "GT",
					"value": {
						"@type":"google.protobuf.Int32Value",
						"value":1
					}
				}
			}`
	resp, err := client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.NoError(t, err)
	res := (*proto.VideoCollectionCountRes)(nil)
	err = gjson.DecodeTo(resp.Body, &res)
	require.NoError(t, err)
	fmt.Println(resp.Body)
	require.Equal(t, int64(1), res.TotalElements)
}

func TestGet(t *testing.T) {
	TestDelete(t)
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	getResp, err := client.GetWithHeaders(url, commonHeaders, 0)
	require.NoError(t, err)
	require.Equal(t, 404, getResp.StatusCode)

	TestCreate(t)
	getResp, err = client.GetWithHeaders(url, commonHeaders, 0)
	require.NoError(t, err)
	require.Equal(t, 200, getResp.StatusCode)
	res := (*proto.VideoCollectionGetRes)(nil)
	err = gjson.DecodeTo(getResp.Body, &res)
	require.Equal(t, uint32(1234), *res.Count)
	require.Equal(t, false, *res.IsOnline)
}

func TestCreate(t *testing.T) {
	url := baseUrl
	data := fmt.Sprintf(`{
			"name": "44444",
			"contentType": "%s",
			"filterType": "%s",
			"count": 1234,
			"isOnline": false
		}`, enum.ContentType_name[int32(enum.ContentType_PortraitVideo)], enum.FilterType_name[int32(enum.FilterType_Manual)])
	httpRes, err := client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	createRes := (*server.ResponseWrapper[proto.VideoCollectionCreateRes])(nil)
	err = gjson.DecodeTo(httpRes.Body, &createRes)
	require.NoError(t, err)
	require.NotNil(t, createRes.Data.InsertedID)
	require.Equal(t, int64(1), createRes.Data.InsertedCount)
	insertedID := *createRes.Data.InsertedID

	url = fmt.Sprintf("%s/%s", baseUrl, insertedID)
	httpRes, err = client.DeleteWithHeaders(url, commonHeaders, "", 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	deleteRes := (*server.ResponseWrapper[proto.VideoCollectionDeleteRes])(nil)
	err = gjson.DecodeTo(httpRes.Body, &deleteRes)
	require.NoError(t, err)
	require.Equal(t, int64(1), deleteRes.Data.DeletedCount)
}

func TestUpdate(t *testing.T) {
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	data := `{
				"name": "每日推荐视频集合",
				"isOnline": true
			}`
	resp, err := client.PatchWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	fmt.Println(resp.Body)
	require.Equal(t, 200, resp.StatusCode)
}

func TestUpsert(t *testing.T) {
	TestDelete(t)
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	// upsert when no record exists
	data := `{
				"name": "每日推荐视频集合",
				"contentType": 9,
				"filterType": 9,
				"count": 1234,
				"isOnline": true
			}`
	resp, err := client.PutWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	fmt.Println(resp.Body)
	require.Equal(t, 200, resp.StatusCode)

	getResp, err := client.GetWithHeaders(url, commonHeaders, 0)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	res := (*proto.VideoCollectionGetRes)(nil)
	err = gjson.DecodeTo(getResp.Body, &res)
	require.Equal(t, uint32(1234), *res.Count)
	require.Equal(t, true, *res.IsOnline)

	// upsert again when record exists
	data = `{
				"name": "每日推荐视频集合",
				"contentType": 10,
				"filterType": 10,
				"count": 1235,
				"isOnline": false
			}`
	resp, err = client.PutWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	fmt.Println(resp.Body)
	require.Equal(t, 200, resp.StatusCode)

	getResp, err = client.GetWithHeaders(url, commonHeaders, 0)
	require.NoError(t, err)
	require.NotNil(t, getResp)
	err = gjson.DecodeTo(getResp.Body, &res)
	require.Equal(t, uint32(1235), *res.Count)
	require.Equal(t, false, *res.IsOnline)

}

func TestDelete(t *testing.T) {
	url := fmt.Sprintf("%s/%s", baseUrl, id)
	resp, err := client.DeleteWithHeaders(url, commonHeaders, "", 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	fmt.Println(resp.Body)
	require.Equal(t, 200, resp.StatusCode)
}

func TestDeleteMulti(t *testing.T) {
	TestCreate(t)
	url := baseUrl + "/delete"
	data := `{
				"name": {
					"@type":"goguru.types.Condition",
					"operator": "Equals",
					"value": {
						"@type":"google.protobuf.StringValue",
						"value":"44444"
					}
				}
			}`
	resp, err := client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, resp)
	fmt.Println(resp.Body)
	require.Equal(t, 200, resp.StatusCode)
}

func TestVideoCollection_Create_One_Count_List_DeleteMulti(t *testing.T) {
	var (
		url         string
		data        string
		httpRes     *httpclient.HttpResponse
		createRes   *proto.VideoCollectionCreateRes
		oneRes      *proto.VideoCollectionOneRes
		countRes    *proto.VideoCollectionCountRes
		listRes     *proto.VideoCollectionListRes
		deleteRes   *proto.VideoCollectionDeleteMultiRes
		insertedID1 string
		insertedID2 string
		err         error
	)
	// test Create twice
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
	err = gjson.DecodeTo(httpRes.Body, &createRes)
	assert.NoError(t, err)
	assert.NotNil(t, createRes.InsertedID)
	assert.Equal(t, int64(1), createRes.InsertedCount)
	insertedID1 = *createRes.InsertedID

	data = fmt.Sprintf(`{
			"name": "测试视频集02",
			"contentType": "%s",
			"filterType": "%s",
			"count": 2345,
			"isOnline": true
		}`, enum.ContentType_name[int32(enum.ContentType_LandscapeVideo)], enum.FilterType_name[int32(enum.FilterType_Ruled)])
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	createRes = (*proto.VideoCollectionCreateRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &createRes)
	assert.NoError(t, err)
	assert.NotNil(t, createRes.InsertedID)
	assert.Equal(t, int64(1), createRes.InsertedCount)
	insertedID2 = *createRes.InsertedID

	// test One 第1次，命中1条记录
	url = baseUrl + "/one"
	data = `{
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
				}
			}`
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	oneRes = (*proto.VideoCollectionOneRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &oneRes)
	assert.NoError(t, err)
	assert.Equal(t, true, oneRes.Found)
	assert.Equal(t, uint32(1234), *oneRes.Item.Count)

	// test One 第2次，无命中记录
	data = `{
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
				}
			}`
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	oneRes = (*proto.VideoCollectionOneRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &oneRes)
	assert.NoError(t, err)
	assert.Equal(t, false, oneRes.Found)

	// test Count 第1次，共2条满足条件的记录
	url = baseUrl + "/count"
	data = `{
				"contentType": {
					"@type":"goguru.types.StringSlice",
					"value":["LandscapeVideo", "PortraitVideo"]
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				}
			}`
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	countRes = (*proto.VideoCollectionCountRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &countRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), countRes.TotalElements)

	// test Count 第2次，共1条满足条件的记录
	url = baseUrl + "/count"
	data = `{
				"name": {
					"@type":"google.protobuf.StringValue",
					"value":"测试视频集01"
				}
			}`
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	countRes = (*proto.VideoCollectionCountRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &countRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), countRes.TotalElements)

	// test List 第1次，返回第2页，每页1条记录，当页有1条记录，为满足条件的第2条记录，其 Name 为 "TemplateName456"
	url = baseUrl + "/list"
	data = `{
				"contentType": {
					"@type":"goguru.types.StringSlice",
					"value":["LandscapeVideo", "PortraitVideo"]
				},
				"isOnline": {
					"@type":"goguru.types.BoolSlice",
					"value":[true, false]
				},
				"pageRequest": {
					"number": 2,
					"size": 1,
					"sorts": [{
						"property": "name",
						"direction": "Asc"
					}]
				}
			}`
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, httpRes)
	assert.Equal(t, 200, httpRes.StatusCode)
	listRes = (*proto.VideoCollectionListRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &listRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), listRes.PageInfo.Number)
	assert.Equal(t, int64(2), listRes.PageInfo.TotalElements)
	assert.Equal(t, int64(2), listRes.PageInfo.TotalPages)
	assert.Equal(t, int64(1), listRes.PageInfo.NumberOfElements)
	assert.Equal(t, false, listRes.PageInfo.First)
	assert.Equal(t, true, listRes.PageInfo.Last)

	// test List 第2次，返回第2页，每页1条记录，当页有1条记录，为满足条件的第2条记录，其 Name 为 "TemplateName123"
	//resList, err = templateRepo.List(ctx, &proto2.TemplateListReq{
	//	Name: types.AnyCondition(q.NewCondition(
	//		types.AnyString("TemplateName"), q.WithOperator(q.OperatorType_Like), q.WithWildcard(q.WildcardType_StartsWith))),
	//	WrapperType: types.AnyStringSlice([]string{model.WrapperType_name[int32(model.WrapperType_Applovin)], model.WrapperType_name[int32(model.WrapperType_None)]}),
	//	MediaElementCount: types.AnyCondition(q.NewCondition(
	//		types.AnyInt32(0), q.WithOperator(q.OperatorType_GT))),
	//	PageRequest: &q.PageRequest{
	//		Number: 2,
	//		Size:   1,
	//		Sorts: []*q.SortParam{
	//			{
	//				Property:  "name",
	//				Direction: q.SortDirection_Desc,
	//			},
	//		},
	//	},
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, resList.PageInfo)
	//assert.NotNil(t, resList.Items)
	//assert.Equal(t, 1, len(resList.Items))
	//assert.Equal(t, "TemplateName123", *resList.Items[0].Name)
	//assert.Equal(t, int64(2), resList.PageInfo.Number)
	//assert.Equal(t, int64(2), resList.PageInfo.TotalPages)
	//assert.Equal(t, int64(2), resList.PageInfo.TotalElements)
	//assert.Equal(t, int64(1), resList.PageInfo.NumberOfElements)
	//assert.Equal(t, false, resList.PageInfo.First)
	//assert.Equal(t, true, resList.PageInfo.Last)
	//
	//// test List 第3次，返回第2页，每页1条记录，当页有0条记录，因为满足条件的只有1条记录，存在于第1页
	//resList, err = templateRepo.List(ctx, &proto2.TemplateListReq{
	//	ExtraFilters: []*q.PropertyFilter{
	//		// 条件1 mediaElements 数组的第0个元素的 name 属性  == "a_001"
	//		{
	//			Property:  "mediaElements.0.name",
	//			Condition: q.NewCondition(types.AnyString("a_001")),
	//		},
	//	},
	//	PageRequest: &q.PageRequest{
	//		Number: 2,
	//		Size:   1,
	//		Sorts: []*q.SortParam{
	//			{
	//				Property:  "name",
	//				Direction: q.SortDirection_Desc,
	//			},
	//		},
	//	},
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, resList.PageInfo)
	//assert.NotNil(t, resList.Items)
	//assert.Equal(t, 0, len(resList.Items))
	//
	//// test List 第2次，返回第2页，每页1条记录，当页有1条记录，为满足条件的第2条记录，其 Name 为 "TemplateName123"
	//resList, err = templateRepo.List(ctx, &proto2.TemplateListReq{
	//	ExtraFilters: []*q.PropertyFilter{
	//		// 条件1 mediaElements 数组的第0个元素的 name 属性以 "_001" 结尾
	//		{
	//			Property:  "mediaElements.0.name",
	//			Condition: q.NewCondition(types.AnyString("_001"), q.WithOperator(q.OperatorType_Like), q.WithWildcard(q.WildcardType_EndsWith)),
	//		},
	//	},
	//	PageRequest: &q.PageRequest{
	//		Number: 2,
	//		Size:   1,
	//		Sorts: []*q.SortParam{
	//			{
	//				Property:  "name",
	//				Direction: q.SortDirection_Desc,
	//			},
	//		},
	//	},
	//})
	//assert.NoError(t, err)
	//assert.NotNil(t, resList.PageInfo)
	//assert.NotNil(t, resList.Items)
	//assert.Equal(t, 1, len(resList.Items))
	//assert.Equal(t, "TemplateName123", *resList.Items[0].Name)
	//assert.Equal(t, int64(2), resList.PageInfo.Number)
	//assert.Equal(t, int64(2), resList.PageInfo.TotalPages)
	//assert.Equal(t, int64(2), resList.PageInfo.TotalElements)
	//assert.Equal(t, int64(1), resList.PageInfo.NumberOfElements)
	//assert.Equal(t, false, resList.PageInfo.First)
	//assert.Equal(t, true, resList.PageInfo.Last)

	// test DeleteMulti 删除2条记录
	url = baseUrl + "/delete"
	data = fmt.Sprintf(`{
				"id": {
					"@type":"goguru.types.ObjectIDSlice",
					"value": [
						{
							"value": "%s"
						},
						{
							"value": "%s"
						}
					]
				}
			}`, insertedID1, insertedID2)
	httpRes, err = client.PostWithHeaders(url, commonHeaders, data, 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	deleteRes = (*proto.VideoCollectionDeleteMultiRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &deleteRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), deleteRes.DeletedCount)

	// test Delete 删除0条记录，因为之前的 deleteMulti 已经删除过了
	url = baseUrl + "/" + insertedID1
	//data = fmt.Sprintf(`{
	//			"id": {
	//				"@type":"goguru.types.ObjectID",
	//				"value": {
	//					"value": "%s"
	//				}
	//			}
	//		}`, insertedID1)
	httpRes, err = client.DeleteWithHeaders(url, commonHeaders, "", 0)
	require.NoError(t, err)
	require.NotNil(t, httpRes)
	require.Equal(t, 200, httpRes.StatusCode)
	deleteRes = (*proto.VideoCollectionDeleteMultiRes)(nil)
	err = gjson.DecodeTo(httpRes.Body, &deleteRes)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), deleteRes.DeletedCount)

}
