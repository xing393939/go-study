package controllers

import (
	"github.com/codegangsta/martini-contrib/render"
	"martini_demo/models"
	"net/http"
)

func HandleGetCurrentUser(r render.Render, req *http.Request, db *models.DB) {
	id := req.FormValue("id")
	if user, err := db.GetUserWithId(id); err == nil {
		r.JSON(200, map[string]interface{}{
			"user": user,
		})
	} else {
		r.JSON(200, map[string]interface{}{
			"user": "",
		})
	}
}
