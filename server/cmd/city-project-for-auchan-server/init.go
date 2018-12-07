package main

import (
	"github.com/itimofeev/auchan/server/restapi"
	"github.com/itimofeev/auchan/service"
	"github.com/itimofeev/auchan/store"
)

func init() {
	s := store.NewStore("postgresql://postgres@db:5432/postgres?sslmode=disable")
	restapi.Service = service.NewService(s)
}
