package model

// Aqui indica as informações relativa ao banco de dados. Modificar DIRETO no arquivo db.go

var User = "root"
var Password = "admin"
var TCP = "@tcp(localhost:3306)/"
var DBName = "godb001"
var Others = "?charset=utf8mb4&parseTime=True&loc=Local"

// var DNS = fmt.Sprintln(User + ":" + Password + TCP + DBName + Others)
