package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Code0716/graphql-study/graph"
	"github.com/Code0716/graphql-study/graph/generated"
	"github.com/Code0716/graphql-study/infrastructure/database"
	"github.com/Code0716/graphql-study/registry"
	"github.com/Code0716/graphql-study/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	code := start()
	os.Exit(code)
}

func start() int {
	env := util.Env()

	dbConn, close, err := database.NewDBConn(env)
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
		return 1
	}
	db, err := database.NewDB(dbConn, env)
	if err != nil {
		log.Fatalf("DB initialization error: %s", err)
		return 1
	}

	defer func() {
		if err := close(); err != nil {
			log.Fatal(err)
		}
	}()

	sqlHandler := new(database.SQLHandler)
	sqlHandler.Conn = db
	reg := registry.New(sqlHandler)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	e.Use(middleware.Recover())

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{Reg: reg}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	for _, r := range e.Routes() {
		fmt.Printf("[%v] %+v\n", r.Method, r.Path)
	}
	addr := util.GetAPIPath(env)

	e.Logger.Fatal(e.Start(addr))

	return 0
}
