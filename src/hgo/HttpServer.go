package hgo

import (
	"context"
	"net/http"
	"time"
)

func StartHttpServer(address string) *http.Server {
	var server = &(http.Server{Addr: address})
	server.ListenAndServe()
	return server
}

func StopHttpServer(server *http.Server) error {
	var ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	return server.Shutdown(ctx)
}
