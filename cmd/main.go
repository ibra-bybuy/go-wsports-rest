package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ibra-bybuy/go-wsports-events/internal/controller/banners"
	"github.com/ibra-bybuy/go-wsports-events/internal/controller/events"
	"github.com/ibra-bybuy/go-wsports-events/internal/controller/tournaments"
	httpBanners "github.com/ibra-bybuy/go-wsports-events/internal/handler/rest/banners"
	httpEvents "github.com/ibra-bybuy/go-wsports-events/internal/handler/rest/events"
	httpTournaments "github.com/ibra-bybuy/go-wsports-events/internal/handler/rest/tournaments"
	bRepository "github.com/ibra-bybuy/go-wsports-events/internal/repository/banners"
	"github.com/ibra-bybuy/go-wsports-events/internal/repository/dotenv"
	eRepository "github.com/ibra-bybuy/go-wsports-events/internal/repository/events"
	"github.com/ibra-bybuy/go-wsports-events/internal/repository/mongodb"
	tRepository "github.com/ibra-bybuy/go-wsports-events/internal/repository/tournaments"
)

func main() {
	// Load .env vars
	dotenv.Load()

	// Mysql
	db := mongodb.New()

	// Repositories
	eventsRepository := eRepository.New(db)
	tournamentsRepository := tRepository.New(db)
	bannersRepository := bRepository.New(db)

	// Controllers
	eventsController := events.New(eventsRepository)
	tournamentsController := tournaments.New(tournamentsRepository)
	bannersController := banners.New(bannersRepository)

	// Handlers
	eventsHandler := httpEvents.New(eventsController)
	tournamentsHandler := httpTournaments.New(tournamentsController)
	bannersHandler := httpBanners.New(bannersController)

	http.Handle("/api/v2/events", http.HandlerFunc(eventsHandler.Handle))
	http.Handle("/api/v2/tournaments", http.HandlerFunc(tournamentsHandler.Handle))
	http.Handle("/api/v2/banners", http.HandlerFunc(bannersHandler.Handle))

	port := dotenv.Get("PORT")
	log.Println(fmt.Sprintf("Listening to port %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic("Error listening to server")
	}

}
