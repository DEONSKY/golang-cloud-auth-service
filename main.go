package main

import (
	_ "github.com/forfam/authentication-service/organization"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/server"
)

func main() {
	postgres.InitAuthenticationDb()

	// app.Post("/files", files.UploadFileEndpoint)

	server.Listen()
}
