package main

import (
	"gonotepad/routes"
	"gonotepad/models"
)

var passw = make(map[int]models.PassModel)


func main(){
	routes.ListUsers(passw)
}