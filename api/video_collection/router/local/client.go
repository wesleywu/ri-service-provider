package local

const (
	serviceName = "VideoCollection"
)

//
//func getVideoCollectionClient() *proto.VideoCollectionClientImpl {
//	svc := service.VideoCollection
//	return &proto.VideoCollectionClientImpl{
//		Count:  proxy.NewInvocationProxy[*proto.VideoCollectionCountReq, *proto.VideoCollectionCountRes](serviceName, "Count", svc.Count),
//		One:    proxy.NewInvocationProxy[*proto.VideoCollectionOneReq, *proto.VideoCollectionOneRes](serviceName, "One", svc.One),
//		List:   proxy.NewInvocationProxy[*proto.VideoCollectionListReq, *proto.VideoCollectionListRes](serviceName, "List", svc.List),
//		Create: proxy.NewInvocationProxy[*proto.VideoCollectionCreateReq, *proto.VideoCollectionCreateRes](serviceName, "Create", svc.Create),
//		Update: proxy.NewInvocationProxy[*proto.VideoCollectionUpdateReq, *proto.VideoCollectionUpdateRes](serviceName, "Update", svc.Update),
//		Upsert: proxy.NewInvocationProxy[*proto.VideoCollectionUpsertReq, *proto.VideoCollectionUpsertRes](serviceName, "Upsert", svc.Upsert),
//		Delete: proxy.NewInvocationProxy[*proto.VideoCollectionDeleteReq, *proto.VideoCollectionDeleteRes](serviceName, "Delete", svc.Delete),
//	}
//}
