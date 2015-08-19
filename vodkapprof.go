package vodkapprof

import (
	"net/http/pprof"

	"github.com/insionng/vodka"
)

func Wrap(v *vodka.Vodka) {
	e.Get("/debug/pprof/", IndexHandler)
	e.Get("/debug/pprof/heap", HeapHandler)
	e.Get("/debug/pprof/goroutine", GoroutineHandler)
	e.Get("/debug/pprof/block", BlockHandler)
	e.Get("/debug/pprof/threadcreate", ThreadCreateHandler)
	e.Get("/debug/pprof/cmdline", CmdlineHandler)
	e.Get("/debug/pprof/profile", ProfileHandler)
	e.Get("/debug/pprof/symbol", SymbolHandler)
}

var Wrapper = Wrap

func IndexHandler(ctx *vodka.Context) error {
	pprof.Index(ctx.Response(), ctx.Request())
	return nil
}

func HeapHandler(ctx *vodka.Context) error {
	pprof.Handler("heap").ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func GoroutineHandler(ctx *vodka.Context) error {
	pprof.Handler("goroutine").ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func BlockHandler(ctx *vodka.Context) error {
	pprof.Handler("block").ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func ThreadCreateHandler(ctx *vodka.Context) error {
	pprof.Handler("threadcreate").ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func CmdlineHandler(ctx *vodka.Context) error {
	pprof.Cmdline(ctx.Response(), ctx.Request())
	return nil
}

func ProfileHandler(ctx *vodka.Context) error {
	pprof.Profile(ctx.Response(), ctx.Request())
	return nil
}

func SymbolHandler(ctx *vodka.Context) error {
	pprof.Symbol(ctx.Response(), ctx.Request())
	return nil
}
