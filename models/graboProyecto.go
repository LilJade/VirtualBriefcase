package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GraboProyecto struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      string             `bson:"userid" json:"userid,omitempty"`
	Titulo      string             `bson:"titulo" json:"titulo,omitempty"`
	Descripcion string             `bson:"descripcion" json:"descripcion,omitempty"`
	Github      string             `bson:"github" json:"github,omitempty"`
	SitioWeb    string             `bson:"sitioWeb" json:"sitioWeb,omitempty"`
	Portada     string             `bson:"portada" json:"portada,omitempty"`
	Empresa     string             `bson:"empresa" json:"empresa,omitempty"`
	Fecha       time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
