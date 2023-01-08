package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/eniabiola/bookings/pkg/config"
	"github.com/eniabiola/bookings/pkg/handler"
	"github.com/eniabiola/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//get the template cache from the app config
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)
	/*http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)*/

	fmt.Printf("Starting application on %s", portNumber)

	/*_ = http.ListenAndServe(
		portNumber,
		nil,
	)*/

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
