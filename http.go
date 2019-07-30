package api

import (
	"context"
	"github.com/obase/center"
	"github.com/obase/conf"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var defaultClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   conf.OptiDuration("http.connectionTimeout", 30*time.Second),
			KeepAlive: conf.OptiDuration("http.connectionKeepalive", 30*time.Second),
		}).DialContext,
		MaxIdleConns:          conf.OptiInt("http.maxIdleConns", 10240),
		IdleConnTimeout:       conf.OptiDuration("http.idleConnTimeout", 90*time.Second),
		TLSHandshakeTimeout:   conf.OptiDuration("http.tlsHandshakeTimeout", 10*time.Second),
		ExpectContinueTimeout: conf.OptiDuration("http.expectContinueTimeout", 1*time.Second),
		MaxIdleConnsPerHost:   conf.OptiInt("http.maxIdleConnsPerHost", 2048),
		ResponseHeaderTimeout: conf.OptiDuration("http.responseHeaderTimeout", 5*time.Second),
	},
	Timeout: conf.OptiDuration("http.requestTimeout", 60*time.Second),
}

func HttpPost(ctx context.Context, serviceName string, uri string, header *http.Header, body io.Reader) (state int, content string, err error) {
	service, err := center.Robin(serviceName)
	if err != nil {
		return
	}
	// 拼接URL
	buf := new(strings.Builder)
	buf.WriteString("http://")
	buf.WriteString(service.Host)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(service.Port))
	buf.WriteString(uri)

	// 创建请求
	req, err := http.NewRequest(http.MethodPost, buf.String(), body)
	if err != nil {
		return
	}
	rsp, err := defaultClient.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	buf.Reset()
	_, err = io.Copy(buf, rsp.Body)
	if err != nil {
		return
	}
	return rsp.StatusCode, buf.String(), nil
}
