package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fwojciec/gqlgen-sqlc-example/dataloaders" // update the username
	"github.com/fwojciec/gqlgen-sqlc-example/gqlgen"      // update the username
	"github.com/fwojciec/gqlgen-sqlc-example/pg"          // update the username
)

func main() {
	// initialize the db
	db, err := pg.Open("dbname=txowumysaeefrk:bcbc3b27697cc997f373f5328ae6e3fb80178d2e816175b2d0cb03d3f0af4772@ec2-44-209-186-51.compute-1.amazonaws.com:5432/ddb9oftk5ga1qi sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	// initialize the dataloaders
	dl := dataloaders.NewRetriever() // <- here we initialize the dataloader.Retriever

	// configure the server
	mux := http.NewServeMux()
	mux.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	dlMiddleware := dataloaders.Middleware(repo)     // <- here we initialize the middleware
	queryHandler := gqlgen.NewHandler(repo, dl)      // <- use dataloader.Retriever here
	mux.Handle("/query", dlMiddleware(queryHandler)) // <- use dataloader.Middleware here

	// run the server
	port := ":8080"
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost%s\n", port)
	fmt.Fprintln(os.Stderr, http.ListenAndServe(port, mux))
}
