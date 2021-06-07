package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DevuelvoProyectoSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"usuarioid,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"usuariorelacionid,omitempty"`
	Proyecto          struct {
		ID          string    `bson:"_id" json:"_id,omitempty"`
		Titulo      string    `bson:"titulo" json:"titulo,omitempty"`
		Descripcion string    `bson:"descripcion" json:"descripcion,omitempty"`
		Portada     string    `bson:"portada" json:"portada,omitempty"`
		Empresa     string    `bson:"empresa" json:"empresa,omitempty"`
		Fecha       time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}
