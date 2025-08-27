package main

import (
	"log"

	"github.com/truyentan/backend/internal/app"
)

func main() {
	srv := app.NewServer()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
