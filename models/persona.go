package models

//import "github.com/jinzhu/gorm"

type Persona struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
