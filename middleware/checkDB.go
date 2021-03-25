package middleware

import(
	db "github.com/LilJade/virtualBriefcase/database"
	"net/http"
)

//Revisamos que la conexion a la DB continue vigente
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Si CheckConnection nos devuelve
		//un cero, la conexión habrá fracasado
		if db.CheckConnection() == 0 {
			http.Error(w, "Se perdió la conexión con la DB...", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
