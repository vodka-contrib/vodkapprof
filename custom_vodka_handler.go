package vodkapprof

import (
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	"github.com/insionng/vodka/engine/fasthttp"
	"net/http"
	"log"
	"sync"
)

type customVodkaHandler struct {
	httpHandler http.Handler

	wrappedHandleFunc vodka.HandlerFunc
	once sync.Once
}

func (ceh *customVodkaHandler) Handle(c vodka.Context) error {
	ceh.once.Do(func() {
		ceh.wrappedHandleFunc = ceh.mustWrapHandleFunc(c)
	})
	return ceh.wrappedHandleFunc(c)
}

func (ceh *customVodkaHandler) mustWrapHandleFunc(c vodka.Context) vodka.HandlerFunc {
	if _, ok := c.Request().(*standard.Request); ok {
		return standard.WrapHandler(ceh.httpHandler)
	} else if _, ok = c.Request().(*fasthttp.Request); ok {
		return NewFastHTTPVodkaAdaptor(ceh.httpHandler)
	}

	log.Fatal("Unknown HTTP implementation")
	return nil
}

func fromHTTPHandler(httpHandler http.Handler) *customVodkaHandler {
	return &customVodkaHandler{ httpHandler: httpHandler }
}

func fromHandlerFunc(serveHTTP func(w http.ResponseWriter, r *http.Request)) *customVodkaHandler {
	return &customVodkaHandler{ httpHandler: &customHTTPHandler{ serveHTTP: serveHTTP }}
}
