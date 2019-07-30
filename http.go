package api

import (
	"context"
	"github.com/obase/conf"
	"github.com/obase/center"
	"io"
	"net"
	"net/http"
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

func HttpPost(ctx context.Context, serviceName string, uri string, header *http.Header, body io.Reader) (state int, rsp string, err error) {

}
