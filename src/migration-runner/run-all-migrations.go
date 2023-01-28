package main

import (
	"github.com/forfam/authentication-service/src/config"
	"github.com/forfam/authentication-service/src/migration"
	"fmt"
)

func main() {

	db := config.SetupDatabaseConnection()

	migrations := map[string] Migration{
		"initial-migration":NewInitialMigration()
	}

	switch param := os.Args[2];param{
		case "up":
			migrations[os.Args[1]].up(db)
			fmt.Println("Migration Successfull")
		case "down":
			migrations[os.Args[1]].down(db)
			fmt.Println("Migration Successfull")
	}
}
