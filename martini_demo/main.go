package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/gzip"
	"github.com/codegangsta/martini-contrib/render"
	"io/ioutil"
	"log"
	"martini_demo/controllers"
	"martini_demo/middleware"
	"martini_demo/models"
	"net/http"
	"time"
)

func readEnv(key string) string {
	file, err := ioutil.ReadFile("./env.json")
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	var env map[string]string
	if err = json.Unmarshal(file, &env); err != nil {
		log.Println(err.Error())
		return ""
	}
	return env[key]
}

func HandleIndex(r render.Render) {
	r.JSON(200, map[string]string{"hello": "you've reached the api"})
}

func HandleNotFound(r render.Render) {
	r.JSON(404, map[string]string{"error": "Resource not found."})
}

func main() {
	var err error
	// Open database connection
	db, err := models.OpenDB("mysql", readEnv("DB_URI"))
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err.Error())
		return
	}

	// Start setting up martini
	m := martini.Classic()

	// Set up the middleware
	m.Use(gzip.All())
	m.Use(render.Renderer())
	m.Use(middleware.BackchannelAuth("test"))
	m.Use(func(req *http.Request, c martini.Context) {
		reqC := &controllers.ShareForAllRequest{
			ExString: "global " + time.Now().Format("2006-01-02 15:04:05"),
			Gid:      controllers.GetGoroutineId(),
		}
		// ShareForAllRequest是所有request共享的，而且每接收一个请求都会new一个新的实例
		c.Map(reqC)
	})

	// Inject the database
	m.Map(db)

	// Map the URL routes
	m.Get("/", HandleIndex)
	m.Group("/api", func(r martini.Router) {
		r.Get("/users", controllers.HandleGetCurrentUser)
		r.Get("/test1", controllers.MyHandle1, controllers.MyDo)
		r.Get("/test2", controllers.MyHandle2)
		r.Get("/test3", controllers.MyHandle3)
	})
	m.NotFound(HandleNotFound)

	m.RunOnAddr(readEnv("API_URL"))
}
