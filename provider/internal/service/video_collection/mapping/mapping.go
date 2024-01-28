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

package mapping

// VideoCollectionColumnsDef defines and stores column names for table demo_video_collection.
type VideoCollectionColumnsDef struct {
	Id          string // 视频集ID，字符串格式
	Name        string // 视频集名称
	ContentType string // 内容类型
	FilterType  string // 筛选类型
	Count       string // 集合内视频数量
	IsOnline    string // 是否上线：0 未上线|1 已上线
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

var (
	VideoCollectionColumnMap = map[string]string{
		"Id":          "_id",
		"Name":        "name",
		"ContentType": "contentType",
		"FilterType":  "filterType",
		"Count":       "count",
		"IsOnline":    "isOnline",
		"CreatedAt":   "createdAt",
		"UpdatedAt":   "updatedAt",
	}
	VideoCollectionColumns = VideoCollectionColumnsDef{
		Id:          "id",
		Name:        "name",
		ContentType: "content_type",
		FilterType:  "filter_type",
		Count:       "count",
		IsOnline:    "is_online",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
	}
)
