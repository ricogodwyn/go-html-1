package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// Change this to true when in production
	app.Production = false

	// Initialize the session manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Production // Set to true in production

	app.Session = session

	// Create template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache:", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	// Initialize the repository and handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// Serve static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Set up routes
	routes := routes(&app)

	// Start the server
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
