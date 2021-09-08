package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/anfelo/chillhacks_platform/api"
	"github.com/anfelo/chillhacks_platform/datastores/postgres"
)

var (
	pguser = os.Getenv("PGUSER")
	pghost = os.Getenv("PGHOST")
	pgport = os.Getenv("PGPORT")
	pgdb   = os.Getenv("PGDATABASE")
	pgpw   = os.Getenv("PGPASSWORD")
)

func StartApplication() {
	dsn := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		pguser, pgpw, pghost, pgport, pgdb,
	)

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
