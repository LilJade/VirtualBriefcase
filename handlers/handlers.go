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
	router.HandleFunc("/login", middleware.CheckDB(routers.ULogin)).Methods("POST")
	router.HandleFunc("/usuario", middleware.CheckDB(middleware.ValidoJWT(routers.VerUsuario))).Methods("GET")
	router.HandleFunc("/actualizar", middleware.CheckDB(middleware.ValidoJWT(routers.ActualizarDatos))).Methods("PUT")

	router.HandleFunc("/subirAvatar", middleware.CheckDB(middleware.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleware.CheckDB(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/mostrarDatos", middleware.CheckDB(middleware.ValidoJWT(routers.Listausuarios))).Methods("GET")

	router.HandleFunc("/createProject", middleware.CheckDB(middleware.ValidoJWT(routers.InsertoProyecto))).Methods("POST")
	router.HandleFunc("/updateProject", middleware.CheckDB(middleware.ValidoJWT(routers.ActualizarProyecto))).Methods("PUT")
	router.HandleFunc("/updatePortada", middleware.CheckDB(middleware.ValidoJWT(routers.SubirPortada))).Methods("GET")
	router.HandleFunc("/getPortada", middleware.CheckDB(routers.ObtenerPortada)).Methods("GET")
	router.HandleFunc("/delProject", middleware.CheckDB(middleware.ValidoJWT(routers.BorroProyecto))).Methods("DELETE")
	router.HandleFunc("/seeProject", middleware.CheckDB(middleware.ValidoJWT(routers.VeoProjects))).Methods("GET")
	router.HandleFunc("/seeFProjects", middleware.CheckDB(middleware.ValidoJWT(routers.VeoProyectosSeguidores))).Methods("GET")

	router.HandleFunc("/seguidor", middleware.CheckDB(middleware.ValidoJWT(routers.Useguidor))).Methods("POST")
	router.HandleFunc("/consultaRelacion", middleware.CheckDB(middleware.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/bajaRelacion", middleware.CheckDB(middleware.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
