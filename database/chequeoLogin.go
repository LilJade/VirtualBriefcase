package database

import (
	"github.com/LilJade/virtualBriefcase/models"
	"golang.org/x/crypto/bcrypt"
)


// verifico  si el usuario existe en el login

func Checklogin(email string, password string)( models.Usuario, bool){
	user,verificar,_:=ChequeoYaExisteUsuario(email) //llamo a funcion para verifique si el usuario existe

	if verificar == false{
		return user,false

	}




 	passwordBytes:=[]byte(password)
 	passwordbd :=[]byte(user.Password)

 	err:=bcrypt.CompareHashAndPassword(passwordbd,passwordBytes)//compuebo la contraseña

 	if err !=nil{
		return user,false
	}

	return user,true




}
