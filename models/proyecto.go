package models

type Proyecto struct {
	Titulo      string `bson:"titulo" json:"titulo,omitempty"`
	Descripcion string `bson:"descripcion" json:"descripcion,omitempty"`
	Portada     string `bson:"portada" json:"portada,omitempty"`
	Empresa     string `bson:"empresa" json:"empresa,omitempty"`
}
