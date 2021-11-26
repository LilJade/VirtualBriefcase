package models

import "time"

type GraboImagen struct {
	ProyectoID string    `bson:"proyectoid" json:"proyectoid,omitempty"`
	Titulo     string    `bson:"titulo" json:"titulo,omitempty"`
	Imagen     string    `bson:"imagen" json:"imagen,omitempty"`
	Fecha      time.Time `bson:"fecha" json:"fecha,omitempty"`
}
