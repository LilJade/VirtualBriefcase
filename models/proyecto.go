package models

type Proyecto struct {
	Titulo      string `bson:"titulo" json:"titulo"`
	Descripcion string `bson:"descripcion" json:"descripcion"`
	Portada     string `bson:"portada" json:"portada"`
	Empresa     string `bson:"empresa" json:"empresa"`
	Github      string `bson:"github" json:"github"`
	SitioWeb    string `bson:"sitioWeb" json:"sitioWeb"`
}
