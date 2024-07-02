package handler

import (
	"bytes"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

var SvcMap = make(map[string]genericclient.Client)

func Gateway(ctx context.Context, c *app.RequestContext) {
	svcName := c.Param("svc")
	cli, ok := SvcMap[svcName]
	if !ok {
		c.JSON(http.StatusOK, NewErr(Err_BadRequest))
		return
	}

	// hlog.Infof("request service: %s, method: %s, url: %s, body: %s",
	// 	svcName, string(c.Method()), c.URI().String(), string(c.GetRawData()))

	req, err := http.NewRequest(string(c.Method()), c.URI().String(), bytes.NewBuffer(c.GetRawData()))
	if err != nil {
		hlog.Warnf("new http request failed: %v", err)
		c.JSON(http.StatusOK, NewErr(Err_RequestServerFail))
		return
	}

	newReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		hlog.Errorf("convert request failed: %v", err)
		c.JSON(http.StatusOK, NewErr(Err_ServerHandleFail))
		return
	}

	resp, err := cli.GenericCall(ctx, "", newReq)
	if err != nil {
		hlog.Errorf("GenericCall err:%v", err)
		bizErr, ok := kerrors.FromBizStatusError(err)
		if !ok {
			c.JSON(http.StatusOK, NewErr(Err_ServerHandleFail))
			return
		}
		c.JSON(http.StatusOK, utils.H{ResponseErrCode: bizErr.BizStatusCode(), ResponseErrMessage: bizErr.BizMessage()})
		return
	}

	realResp, ok := resp.(*generic.HTTPResponse)
	if !ok {
		c.JSON(http.StatusOK, NewErr(Err_ServerHandleFail))
		return
	}

	realResp.Body[ResponseErrCode] = 0
	realResp.Body[ResponseErrMessage] = "ok"
	c.JSON(http.StatusOK, realResp.Body)
}
