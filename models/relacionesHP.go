package models

type RelacionesHP struct {
	ProyectoID string `bson:"proyectoid" json:"proyectoId"`
	HerramientasID    string `bson:"herramientasid" json:"herramientasId"`
}
