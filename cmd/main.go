package main

import (
	"net"
	"net/http"
	"os"

	"github.com/Iftikhor99/gosql/cmd/app"
	"github.com/Iftikhor99/gosql/pkg/banners"
)

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}

}

func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)
	server.Init()
	//	bannersSvc.Initial()
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}

	return srv.ListenAndServe()
}
