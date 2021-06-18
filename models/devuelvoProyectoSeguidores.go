package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoProyectoSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"usuarioid,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"usuariorelacionid,omitempty"`
	Proyectos         struct {
		ID          string    `bson:"_id" json:"_id,omitempty"`
		Titulo      string    `bson:"titulo" json:"titulo,omitempty"`
		Descripcion string    `bson:"descripcion" json:"descripcion,omitempty"`
		Portada     string    `bson:"portada" json:"portada,omitempty"`
		Empresa     string    `bson:"empresa" json:"empresa,omitempty"`
		Github      string    `bson:"github" json:"github,omitempty"`
		SitioWeb    string    `bson:"sitioWeb" json:"sitioWeb,omitempty"`
		Fecha       time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}
