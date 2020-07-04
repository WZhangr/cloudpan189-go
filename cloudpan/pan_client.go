package cloudpan

import (
	"github.com/tickstep/cloudpan189-go/library/requester"
	"net/http"
	"net/url"
)

const (
	// PathSeparator 路径分隔符
	PathSeparator = "/"
)

var (
	cloudpanDomainUrl = &url.URL{
		Scheme: "http",
		Host:   ".cloud.189.cn",
	}
)

type PanClient struct {
	client     *requester.HTTPClient // http 客户端
}

func NewPanClient(cookieLoginUser string) *PanClient {
	client := requester.NewHTTPClient()
	client.ResetCookiejar()
	client.Jar.SetCookies(cloudpanDomainUrl, []*http.Cookie{
		&http.Cookie{
			Name:   "COOKIE_LOGIN_USER",
			Value:  cookieLoginUser,
			Domain: "cloud.189.cn",
			Path: "/",
		},
	})

	// debug
	client.SetProxy("http://127.0.0.1:8888")

	return &PanClient{
		client: client,
	}
}