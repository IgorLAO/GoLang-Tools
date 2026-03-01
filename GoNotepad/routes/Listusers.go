package routes

import (
	"gonotepad/models"
)

func ListUsers(passw map[int]models.PassModel){
	
}

func CreateUser(model []models.PassModel, pass string) models.PassModel{
	newID := len(model) + 1
	newPass := models.PassModel{
		ID: newID,
		Passw: pass,
	}

	return newPass
}