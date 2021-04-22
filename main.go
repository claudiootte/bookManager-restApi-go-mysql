package main

import (
	"github.com/claudiootte/bookManager-restApi-go-mysql/controller"
	"github.com/claudiootte/bookManager-restApi-go-mysql/model"
)

func main() {
	model.InitialMigration()
	controller.InitializeRouter()
}
