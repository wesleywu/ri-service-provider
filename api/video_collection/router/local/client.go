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

package local

import (
	"github.com/WesleyWu/gowing/rpc/proxy"
	proto "github.com/WesleyWu/ri-service-provider/proto/video_collection"
	"github.com/WesleyWu/ri-service-provider/provider/video_collection_mongo/service"
)

const (
	serviceName = "VideoCollection"
)

func getVideoCollectionClient() *proto.VideoCollectionClientImpl {
	svc := service.VideoCollection
	return &proto.VideoCollectionClientImpl{
		Count:  proxy.NewInvocationProxy[*proto.VideoCollectionCountReq, *proto.VideoCollectionCountRes](serviceName, "Count", svc.Count),
		One:    proxy.NewInvocationProxy[*proto.VideoCollectionOneReq, *proto.VideoCollectionOneRes](serviceName, "One", svc.One),
		List:   proxy.NewInvocationProxy[*proto.VideoCollectionListReq, *proto.VideoCollectionListRes](serviceName, "List", svc.List),
		Create: proxy.NewInvocationProxy[*proto.VideoCollectionCreateReq, *proto.VideoCollectionCreateRes](serviceName, "Create", svc.Create),
		Update: proxy.NewInvocationProxy[*proto.VideoCollectionUpdateReq, *proto.VideoCollectionUpdateRes](serviceName, "Update", svc.Update),
		Upsert: proxy.NewInvocationProxy[*proto.VideoCollectionUpsertReq, *proto.VideoCollectionUpsertRes](serviceName, "Upsert", svc.Upsert),
		Delete: proxy.NewInvocationProxy[*proto.VideoCollectionDeleteReq, *proto.VideoCollectionDeleteRes](serviceName, "Delete", svc.Delete),
	}
}
