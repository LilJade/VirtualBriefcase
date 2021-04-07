package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/LilJade/virtualBriefcase/middleware"
	"github.com/LilJade/virtualBriefcase/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto, el Handler y pongo a escuchar al servicio
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login",middleware.CheckDB(routers.ULogin)).Methods("POST")
	router.HandleFunc("/actualizar",middleware.CheckDB(middleware.ValidoJWT(routers.ActualizarDatos))).Methods("PUT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
