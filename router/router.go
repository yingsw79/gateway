package router

import (
	"context"
	"gateway/handler"
	"net/http"
	"os"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Register(r *server.Hertz) {
	registerGateway(r)
}

func registerGateway(r *server.Hertz) {
	if handler.SvcMap == nil {
		handler.SvcMap = make(map[string]genericclient.Client)
	}

	idlPath := "idl"
	entries, err := os.ReadDir(idlPath)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
	}

	etcdResolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		hlog.Fatalf("err:%v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == "common.thrift" {
			continue
		}
		svcName := strings.Split(entry.Name(), ".")[0]
		provider, err := generic.NewThriftFileProvider(entry.Name(), idlPath)
		if err != nil {
			hlog.Fatalf("new thrift file provider failed: %v", err)
			break
		}
		g, err := generic.HTTPThriftGeneric(provider)
		if err != nil {
			hlog.Fatal(err)
		}
		cli, err := genericclient.NewClient(
			svcName,
			g,
			client.WithResolver(etcdResolver),
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
			// client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
		)
		if err != nil {
			hlog.Fatal(err)
		}
		handler.SvcMap[svcName] = cli
	}

	g := r.Group("/gateway")
	g.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "gateway is running")
	})
	g.Any("/:svc/*_", handler.Gateway)
}
