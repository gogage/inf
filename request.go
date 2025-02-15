package inf

type HttpCookie struct {
	Name     string
	Path     string
	Age      string
	IsSecure bool
	HttpOnly string
	Domain   string
}

type Value interface {
	String(defaultValue ...string) string
	Strings() []string
	Int(defaultValue ...int) int
	Ints() []int
	Int32(defaultValue ...int32) int32
	Int32s() []int32
	Int64(defaultValue ...int64) int64
	Int64s() []int64
	Float(defaultValue ...float32) float32
	Floats() []float32
	Float64(defaultValue ...float64) float64
	Float64s() []float64
	Bool(defaultValue ...bool) bool
	Bools() []bool
	Object(defaultValue ...Getter) Getter
	Objects() []Getter
	Marshal(output interface{}) error
}

type Getter interface {
	Get(name string) Value
}

type HttpRequestPath interface {
	Path() string
	Params() Getter
}

type HttpRequest interface {
	Path() HttpRequestPath
	Query() Getter
	Body() Value
}

type HttpResponse interface {
	ContentType(contentType string) HttpResponse
	Body(data interface{}) HttpResponse
	Cookie(cookie HttpCookie, cookies ...HttpCookie) HttpResponse
	Redirect(location string) HttpResponse
	Status(code int, text ...string) HttpResponse
}

type HttpHandler interface {
	Route() string
	Handle(in HttpRequest, out HttpResponse)
}

type Module interface {
	ID() string
	HttpHandlers() []HttpHandler
}
