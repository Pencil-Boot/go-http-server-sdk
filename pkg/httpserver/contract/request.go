package httpservercontract

type RequestHeaders interface {
	Get(key string) []string
	All() map[string][]string
}

type Params interface {
	Get(key string) string
	All() map[string]string
}

type QueryParams interface {
	Get(key string) []string
	All() map[string][]string
}

type Request interface {
	Method() string
	Path() string

	Headers() RequestHeaders
	Header(key string) string

	Params() Params
	Param(key string) string

	QueryParams() QueryParams
	QueryParam(key string) string

	Bind(ptr any) error
}
