package models

import "time"

// Imagen model
type Imagen struct {
	Titulo string    `bson:"titulo" json:"titulo,omitempty"`
	Imagen string    `bson:"imagen" json:"imagen,omitempty"`
	Fecha  time.Time `bson:"fecha" json:"fecha,omitempty"`
}
