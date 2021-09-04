package main

import (
	"log"
	"net/http"

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
	http.ListenAndServe(":3000", h)
}
