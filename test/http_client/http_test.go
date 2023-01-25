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
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	url := "http://localhost:8888/GowingService/com.github.wesleywu.ri_service_provider.VideoCollectionRepo/Create"
	data := `{
			"id": "01186883-7698",
			"name": "44444",
			"contentType": 9,
			"filterType": 9,
			"count": 1234,
			"isOnline": "false"
		}`
	client := &http.Client{Timeout: 500 * time.Second}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("x-dubbo-http1.1-dubbo-version", "1.0.0")
	req.Header.Set("x-dubbo-service-protocol", "triple")
	req.Header.Set("x-dubbo-service-version", "1.0.0")
	req.Header.Set("x-dubbo-service-group", "test")
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUpdate(t *testing.T) {
	url := "http://localhost:8888/GowingService/com.github.wesleywu.ri_service_provider.VideoCollectionRepo/Update"
	data := `{
				"id": "01186883-7698",
				"name": "你好啊",
				"isOnline": "true"
			}`
	client := &http.Client{Timeout: 500 * time.Second}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("x-dubbo-http1.1-dubbo-version", "1.0.0")
	req.Header.Set("x-dubbo-service-protocol", "triple")
	req.Header.Set("x-dubbo-service-version", "1.0.0")
	req.Header.Set("x-dubbo-service-group", "test")
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
