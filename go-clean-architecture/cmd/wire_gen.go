// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"my-clean-rchitecture/api"
	"my-clean-rchitecture/api/handlers"
	"my-clean-rchitecture/app"
	"my-clean-rchitecture/repo"
	"my-clean-rchitecture/service"
)

// Injectors from wire.go:

func InitServer() *app.Server {
	engine := app.NewGinEngine()
	db := app.InitDB()
	iArticleRepo := repo.NewMysqlArticleRepository(db)
	iArticleService := service.NewArticleService(iArticleRepo)
	articleHandler := handlers.NewArticleHandler(iArticleService)
	router := api.NewRouter(articleHandler)
	server := app.NewServer(engine, router)
	return server
}
