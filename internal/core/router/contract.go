package router

import httpservercontract "go-http-server-sdk/pkg/httpserver/contract"

type RouterRegister interface {
	Register(
		method httpservercontract.HttpMethod,
		path string,
		handler httpservercontract.Handler,
		config *httpservercontract.RouterConfig,
	)
	Group(prefix string) RouterRegister
}
