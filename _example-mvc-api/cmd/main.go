package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/mkosakana/go-iris-sample/_example-mvc-api/configs"
)

func main() {
	app := iris.New()

	// middleware
	app.Use(iris.Compression)
	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

	// routes
	err := configs.SetRoutes(app)
	if err != nil {
		log.Fatalf("func %s failed: %v", "SetRoutes", err)
	}

	// port
	err = configs.SetPort(app)
	if err != nil {
		log.Fatalf("func %s failed: %v", "SetPort", err)
	}
}
