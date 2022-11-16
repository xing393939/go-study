package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"martini_demo/controllers"
	"martini_demo/models"
)

var a int = 1

const name *int = a

func readEnv(key string) string {
	file, err := ioutil.ReadFile("./env.json")
	if err != nil {
		println(err.Error())
		return ""
	}
	var env map[string]string
	if err = json.Unmarshal(file, &env); err != nil {
		println(err.Error())
		return ""
	}
	return env[key]
}

func main() {
	db, _ := models.OpenDB("mysql", readEnv("DB_URI"))

	m := martini.Classic()
	m.Use(render.Renderer())
	// Inject the database
	m.Map(db)
	m.Get("/", controllers.HandleGetCurrentUser)
	m.RunOnAddr(readEnv("API_URL"))
}
