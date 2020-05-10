package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	"github.com/rs/cors"


	"github.com/MieGorenk/data-api/api/models"
	"github.com/MieGorenk/data-api/api/responses"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize db connection and wire up router
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cant connect to the DB %s", DbName)
		log.Fatal("This is the error", err)

	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "main." + defaultTableName
	}
	a.DB.Debug().AutoMigrate(&models.Movie{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/movies", a.CreateMovie).Methods("POST")
	a.Router.HandleFunc("/movies", a.GetMovies).Methods("GET")
	a.Router.HandleFunc("/movies/{id:[0-9]+}", a.GetMovieByID).Methods("GET")
}

func (a *App) RunServer() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
    	AllowedHeaders: []string{"*"},
    	Debug:          true,
	})

	handler := c.Handler(a.Router)
	log.Printf("\nServer starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000", handler))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	responses.JSON(w, http.StatusOK, "Welcome To ProjectA API")
}
