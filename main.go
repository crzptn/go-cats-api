package main

import (
	"context"
	"database/sql"
	"example/cats/db"
	"example/cats/internal"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	_ "modernc.org/sqlite"
)

var port = flag.Uint("port", 6969, "port which the server to run on")
var dbFile = flag.String("db", "database.db", "sqlite database")

func main() {
	flag.Parse()
	conn, err := sql.Open("sqlite", *dbFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	application := app.New(db.New(conn), fmt.Sprintf("127.0.0.1:%d", *port))
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()
	err = application.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
