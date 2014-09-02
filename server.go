package main

import (

	// "fmt" // Commented out to avoid Go errors for unused imports. Uncomment after use.

	"github.com/albrow/negroni-json-recovery"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	// Commented out to avoid Go errors for empty imports. Uncomment after adding controllers/models.
	// "github.com/soroushjp/go-rest/controllers"
	// "github.com/soroushjp/go-rest/models"
)

func main() {

	router := mux.NewRouter()

	n := negroni.New(negroni.NewLogger())
	n.UseHandler(router)
	n.Use(recovery.JSONRecovery(true))

	n.Run(":3000")

}
