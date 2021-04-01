package models

//RespuestaLogin tiene el token que se devolvio del login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
