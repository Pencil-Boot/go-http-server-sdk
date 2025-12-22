package httpservercontract

type ResponseHeaders interface {
	Get(key string) []string
	Set(key, value string)
	Add(key, value string)
	All() map[string][]string
}

type Response interface {
	Status() int
	Headers() ResponseHeaders
	Body() any
}
