package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Code0716/graphql-study/graph"
	"github.com/Code0716/graphql-study/graph/generated"
	"github.com/Code0716/graphql-study/infrastructure/database"
	"github.com/Code0716/graphql-study/registry"
	"github.com/Code0716/graphql-study/util"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	env := util.Env()

	dbConn, close, err := database.NewDBConn(env)
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
		return
	}

	db, err := database.NewDB(dbConn, env)
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
		return
	}

	defer func() {
		if err := close(); err != nil {
			log.Fatal(err)
		}
	}()
	sqlHandler := new(database.SQLHandler)
	sqlHandler.Conn = db
	reg := registry.New(sqlHandler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Reg: reg}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://127.0.0.1:%s/ for GraphQL playground", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Print(err)
	}
}
