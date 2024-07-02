package main

import (
	"context"
	api "gateway/test/hello/kitex_gen/api"

	"github.com/cloudwego/kitex/pkg/klog"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	klog.Infof("Echo: %s", req.Message)
	resp = &api.Response{Message: req.Message}
	return
}
