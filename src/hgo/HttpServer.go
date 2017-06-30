package hgo

import (
	"context"
	"net/http"
	"time"
)

func StartHttpServer(address string) *http.Server {
	var server = &(http.Server{Addr: address})
	go server.ListenAndServe()
	return server
}

func StopHttpServer(server *http.Server, timeout time.Duration) error {
	var ctx, _ = context.WithTimeout(context.Background(), timeout)
	return server.Shutdown(ctx)
}
