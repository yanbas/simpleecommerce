package main

import (
	"ecommerce/DB"
	"ecommerce/app"
	"log"
)

func main() {
	// Init DB
	db, err := DB.DbInit()
	if err != nil {
		log.Fatal(err)
	}

	r := app.App{}
	err = r.Run(db)
	if err != nil {
		log.Fatal(err)
	}

}
