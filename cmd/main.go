package main

import (
	"context"
	//"encoding/hex"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Iftikhor99/gosql/cmd/app"
	"github.com/Iftikhor99/gosql/pkg/customers"
	"github.com/Iftikhor99/gosql/pkg/managers"
	"github.com/Iftikhor99/gosql/pkg/security"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/dig"
	//	"golang.org/x/crypto/bcrypt"
)

func main() {

	// password := "secret"
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Print(err)
	// 	os.Exit(1)
	// }
	// log.Print(hex.EncodeToString(hash))
	// err = bcrypt.CompareHashAndPassword(hash, []byte("password"))
	// if err != nil {
	// 	log.Print("Password is invalid")
	// 	os.Exit(1)
	// }

	host := "0.0.0.0"
	port := "9999"
	dsn := "postgres://app:pass@192.168.99.100:5432/db"

	if err := execute(host, port, dsn); err != nil {
		log.Print(err)
		os.Exit(1)
	}

}

func execute(host string, port string, dsn string) (err error) {
	deps := []interface{}{
		app.NewServer,
		mux.NewRouter,
		func() (*pgxpool.Pool, error) {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)	
			return pgxpool.Connect(ctx, dsn)
		},
		customers.NewService,
		func(server *app.Server) *http.Server {
			return &http.Server{
				Addr: net.JoinHostPort(host, port),
				Handler: server,
			}
		},
		managers.NewService,
		security.NewService,
	}

	container := dig.New()
	for _, dep := range deps {
		err = container.Provide(dep)
		if err != nil {
			return err
		}
	}
	
	err = container.Invoke(func(server *app.Server){
		server.Init() 	
	})
	if err != nil {
		return err
	}

	return container.Invoke(func(server *http.Server) error {
		return server.ListenAndServe()
	})

	
}
