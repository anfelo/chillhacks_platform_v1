package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anfelo/chillhacks_platform/api"
	"github.com/anfelo/chillhacks_platform/datastores/postgres"
)

func StartApplication() {
	dsn := "postgres://postgres:secret@localhost/postgres?sslmode=disable"

	store, err := postgres.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	h := api.NewHandler(store)
	fmt.Println("Listening on port :3000")
	if err := http.ListenAndServe(httpPort(), h); err != nil {
		panic(err)
	}
}

func httpPort() string {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
