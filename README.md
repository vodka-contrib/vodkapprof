# vodkapprof
A wrapper for golang web framework vodka to use net/http/pprof easily.
# install
First install vodkapprof to your GOPATH using go get:
```
go get github.com/vodka-contrib/vodkapprof
```
# Usage
```
package main

import (
    "github.com/insionng/vodka"
    "github.com/insionng/vodka/engine/standard"
    "github.com/vodka-contrib/vodkapprof"
)

func hello() vodka.HandlerFunc {
	return func(c vodka.Context) error {
		return c.String(200, "Hello, World!\n")
	}
}

func main() {
    e := vodka.New()
    e.Get("/", hello())

    // automatically add routers for net/http/pprof
    // e.g. /debug/pprof, /debug/pprof/heap, etc.
    vodkapprof.Wrapper(e)
    e.Run(standard.New(":8080"))
}
```
Start this server, and now visit http://127.0.0.1:8080/debug/pprof/ and you'll see what you want.



