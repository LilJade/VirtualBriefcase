package models

import "time"

type GraboProyecto struct {
	UserID      string    `bson:"userid" json:"userid,omitempty"`
	Titulo      string    `bson:"titulo" json:"titulo,omitempty"`
	Descripcion string    `bson:"descripcion" json:"descripcion,omitempty"`
	Portada     string    `bson:"portada" json:"portada,omitempty"`
	Empresa     string    `bson:"empresa" json:"empresa,omitempty"`
	Fecha       time.Time `bson:"fecha" json:"fecha,omitempty"`
}
