package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/stretchr/testify/assert"
	"github.com/wesleywu/ri-service-provider/gwhttpclient"
)

var (
	ctx           = context.Background()
	client        = gwhttpclient.New(100)
	commonHeaders = &http.Header{
		"Content-Type":                  []string{"application/json"},
		"x-dubbo-http1.1-dubbo-version": []string{"1.0.0"},
		"x-dubbo-service-protocol":      []string{"triple"},
	}
)

const id = "01186883-7700"

func TestList(t *testing.T) {
	url := "http://localhost:8080/v1/video-collection/list"
	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, "{}", 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	respJson, err := gjson.DecodeToJson(resp.Body)
	assert.NoError(t, err)
	fmt.Println(resp.Body)
	assert.Equal(t, 3, respJson.Get("total").Int())
}

func TestListBySingleValue(t *testing.T) {
	url := "http://localhost:8080/v1/video-collection/list"
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

	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	respJson, err := gjson.DecodeToJson(resp.Body)
	assert.NoError(t, err)
	fmt.Println(resp.Body)
	assert.Equal(t, 1, respJson.Get("total").Int())
}

func TestListBySliceValue(t *testing.T) {
	url := "http://localhost:8080/v1/video-collection/list"
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
	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NoError(t, err)
	respJson, err := gjson.DecodeToJson(resp.Body)
	assert.NoError(t, err)
	fmt.Println(resp.Body)
	assert.Equal(t, 2, respJson.Get("total").Int())
}

func TestListByCondition(t *testing.T) {
	TestCreate(t)
	url := "http://localhost:8080/v1/video-collection/list"
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
	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NoError(t, err)
	respJson, err := gjson.DecodeToJson(resp.Body)
	assert.NoError(t, err)
	fmt.Println(resp.Body)
	assert.Equal(t, 1, respJson.Get("total").Int())
}

func TestGet(t *testing.T) {
	TestDelete(t)
	url := fmt.Sprintf("http://localhost:8080/v1/video-collection/%s", id)
	getResp, err := client.DoGetWithHeaders(ctx, url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.Equal(t, 404, getResp.StatusCode)

	TestCreate(t)
	getResp, err = client.DoGetWithHeaders(ctx, url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.Equal(t, 200, getResp.StatusCode)
	respJson, _ := gjson.DecodeToJson(getResp.Body)
	assert.Equal(t, uint32(1234), respJson.Get("count").Uint32())
	assert.Equal(t, false, respJson.Get("isOnline").Bool())
}

func TestCreate(t *testing.T) {
	TestDelete(t)
	url := "http://localhost:8080/v1/video-collection"
	data := fmt.Sprintf(`{
			"id": "%s",
			"name": "44444",
			"contentType": 9,
			"filterType": 9,
			"count": 1234,
			"isOnline": false
		}`, id)
	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdate(t *testing.T) {
	url := fmt.Sprintf("http://localhost:8080/v1/video-collection/%s", id)
	data := `{
				"name": "每日推荐视频集合",
				"isOnline": true
			}`
	resp, err := client.DoPatchWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpsert(t *testing.T) {
	TestDelete(t)
	url := fmt.Sprintf("http://localhost:8080/v1/video-collection/%s", id)
	// upsert when no record exists
	data := `{
				"name": "每日推荐视频集合",
				"contentType": 9,
				"filterType": 9,
				"count": 1234,
				"isOnline": true
			}`
	resp, err := client.DoPutWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)

	getResp, err := client.DoGetWithHeaders(ctx, url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.NotNil(t, getResp)
	respJson, _ := gjson.DecodeToJson(getResp.Body)
	assert.Equal(t, uint32(1234), respJson.Get("count").Uint32())
	assert.Equal(t, true, respJson.Get("isOnline").Bool())

	// upsert again when record exists
	data = `{
				"name": "每日推荐视频集合",
				"contentType": 10,
				"filterType": 10,
				"count": 1235,
				"isOnline": false
			}`
	resp, err = client.DoPutWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)

	getResp, err = client.DoGetWithHeaders(ctx, url, commonHeaders, 0)
	assert.NoError(t, err)
	assert.NotNil(t, getResp)
	respJson, _ = gjson.DecodeToJson(getResp.Body)
	assert.Equal(t, uint32(1235), respJson.Get("count").Uint32())
	assert.Equal(t, false, respJson.Get("isOnline").Bool())

}

func TestDelete(t *testing.T) {
	url := fmt.Sprintf("http://localhost:8080/v1/video-collection/%s", id)
	resp, err := client.DoDeleteWithHeaders(ctx, url, commonHeaders, "", 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestDeleteMulti(t *testing.T) {
	TestCreate(t)
	url := "http://localhost:8080/v1/video-collection/delete"
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
	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}
