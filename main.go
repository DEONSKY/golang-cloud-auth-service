package main

import (
	_ "github.com/forfam/authentication-service/group"
	_ "github.com/forfam/authentication-service/grouppolicy"
	_ "github.com/forfam/authentication-service/organization"
	_ "github.com/forfam/authentication-service/policy"
	"github.com/forfam/authentication-service/postgres"
	"github.com/forfam/authentication-service/server"
)

func main() {
	postgres.InitAuthenticationDb()

	// app.Post("/files", files.UploadFileEndpoint)

	server.Listen()
}
