package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario es el modelo de usuarios de la bade MongoDB
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempty"`
	Direccion       string             `bson:"direccion" json:"direccion,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	Profesion       string             `bson:"profesion" json:"profesion,omitempty"`
	Creado          time.Time          `bson:"creado" json:"creado,omitempty"`
	Instagram       string             `bson:"instagram" json:"instagram,omitempty"`
	Facebook        string             `bson:"facebook" json:"facebook,omitempty"`
	Twitter         string             `bson:"twitter" json:"twitter,omitempty"`
}
