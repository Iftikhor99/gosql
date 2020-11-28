package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/Iftikhor99/gosql/cmd/app"
	"github.com/Iftikhor99/gosql/pkg/customers"
	_"github.com/jackc/pgx/v4/stdlib"
)

func main() {
	host := "0.0.0.0"
	port := "9999"
	dsn := "postgres://app:pass@192.168.99.100:5432/db"

	if err := execute(host, port, dsn); err != nil {
		log.Print("We are here3")
		os.Exit(1)
	}

}

func execute(host string, port string, dsn string) (err error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Print("We are here1")
		return err
	} 

	defer func() {
		if cerr := db.Close(); cerr != nil {
			if err == nil {
				err = cerr
				log.Print("We are here2")
				return
			} 
			log.Print(err)
			
		}
	}()


	mux := http.NewServeMux()
	customersSvc := customers.NewService(db)
	server := app.NewServer(mux, customersSvc)
	server.Init()
	//	bannersSvc.Initial()
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}

	return srv.ListenAndServe()
}
