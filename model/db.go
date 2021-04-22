package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// Inicializa-se a comunicação entre a aplicaçãoe o banco de dados e fazendo a migração dos dados

const DNS = "root:admin@tcp(localhost:3306)/godb001?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Book{})
}
