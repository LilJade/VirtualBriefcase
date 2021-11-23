package models

// RelacionesHU
type RelacionesHu struct {
	UsuarioID      string `bson:"usuarioid" json:"usuarioId"`
	HerramientasID string `bson:"herramientasid" json:"herramientasId"`
}
