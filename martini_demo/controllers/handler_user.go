package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"log"
	"martini_demo/middleware"
	"martini_demo/models"
)

func HandleGetCurrentUser(r render.Render, backChannel middleware.Backchannel, db *models.DB) {
	if backChannel.UserId() == "" {
		r.JSON(200, map[string]*models.User{"user": nil})
		return
	}
	if user, err := db.GetUserWithId(backChannel.UserId()); err == nil {
		r.JSON(200, map[string]*models.User{"user": user})
	} else {
		log.Println("Error getting current user:", err.Error())
		r.JSON(500, "Sorry, an internal server error has occurred.")
	}
}
