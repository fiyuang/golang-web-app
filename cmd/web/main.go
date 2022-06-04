package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fiyuang/golang-web-app/pkg/config"
	"github.com/fiyuang/golang-web-app/pkg/handlers"
	"github.com/fiyuang/golang-web-app/pkg/render"
)

const portNumber = ":9090"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
