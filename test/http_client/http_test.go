/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
					"@type":"gwtypes.StringSlice",
					"value":["每日推荐视频","日常推荐视频"]
				},
				"isOnline": {
					"@type":"gwtypes.BoolSlice",
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
	url := "http://localhost:8080/v1/video-collection/list"
	data := `{
				"name": {
					"@type":"gwtypes.Condition",
					"operator": "Like",
					"wildcard": "StartsWith",
					"value": {
						"@type":"google.protobuf.StringValue",
						"value":"每日"
					}
				},
				"count": {
					"@type":"gwtypes.Condition",
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

func TestCreate(t *testing.T) {
	url := "http://localhost:8080/v1/video-collection"
	data := `{
			"id": "01186883-7700",
			"name": "44444",
			"contentType": 9,
			"filterType": 9,
			"count": 1234,
			"isOnline": false
		}`
	resp, err := client.DoPostWithHeaders(ctx, url, commonHeaders, data, 0)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	fmt.Println(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdate(t *testing.T) {
	url := "http://localhost:8080/v1/video-collection/id/01186883-7700"
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
	url := "http://localhost:8080/v1/video-collection/id/01186883-7700"
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
}

func TestDelete(t *testing.T) {
	url := "http://localhost:8080/v1/video-collection/id/01186883-7700"
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
					"@type":"gwtypes.Condition",
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
