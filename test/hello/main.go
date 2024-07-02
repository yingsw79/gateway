package main

import (
	api "gateway/test/hello/kitex_gen/api/hello"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		klog.Fatal(err)
	}

	svr := api.NewServer(
		new(HelloImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8081}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
