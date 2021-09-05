package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anfelo/chillhacks_platform/api"
	"github.com/anfelo/chillhacks_platform/datastores/postgres"
)

func main() {
	dsn := "postgres://postgres:secret@localhost/postgres?sslmode=disable"

	store, err := postgres.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	csrfKey := []byte("012345678901234567890123456789")
	h := api.NewHandler(store, csrfKey)
	fmt.Println("Listening on port :8000")
	if err := http.ListenAndServe(httpPort(), h); err != nil {
		panic(err)
	}
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
