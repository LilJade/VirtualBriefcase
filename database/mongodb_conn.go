package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Variable que contiene nuestra cadena de conexion a MongoDB
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:1234@cluster0.o3xiq.mongodb.net/db?retryWrites=true&w=majority")

func conectarDB() *mongo.Client {
	//Dentro de cliente guardamos la conexion y
	//si ocurre algun error se guardará en err
	cliente, err := mongo.Connect(context.TODO(), clientOptions)

	//Si la variable err no está vacía(contiene un error)
	//imprimimos el error y retornamos el cliente(la conexion (? )
	if err != nil {
		log.Fatal(err.Error())
		return cliente
	}

	//Realizamos un Ping() a modo de comprobacion (?
	//y si nos devuelve un resultado
	//lo almacenamos en la variable err
	err = cliente.Ping(context.TODO(), nil)

	//Si la variable err no está vacía(contiene un error)
	//imprimimos el error y retornamos el cliente(la conexion (? )
	if err != nil {
		log.Fatal(err.Error())
		return cliente
	}

	//Si no ocurre ningun problema imprimimos }
	//que la conexion fué exitosa y retornamos el cliente(la conexión (? )
	log.Println("Conexión exitosa")
	return cliente
}

//Para comprobar la conexion a la DB
//lo haremos realizando Ping() con la variable MongoConn
var MongoConn = conectarDB()

func CheckConnection() int {
	//Hacemos Ping y si ocurre algun error
	//se guardará en la variable err
	err := MongoConn.Ping(context.TODO(), nil)

	//Si la variable err no está vacía(contiene un error)
	//retornamos un cero
	if err != nil {
		return 0
	}

	//Si no ocurre ningun problema retornamos un uno
	return 1
}
