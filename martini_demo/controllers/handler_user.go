package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"martini_demo/middleware"
	"martini_demo/models"
)

func HandleGetCurrentUser(r render.Render, backChannel middleware.Backchannel, db *models.DB) {
	if user, err := db.GetUserWithId(backChannel.UserId()); err == nil {
		r.JSON(200, map[string]interface{}{
			"user": user,
			"auth": backChannel,
		})
	} else {
		r.JSON(200, map[string]interface{}{
			"user": "",
			"auth": backChannel,
		})
	}
}
