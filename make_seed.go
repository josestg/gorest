package main

import "github.com/josestnggng/Pretest-privy-full-rest-api/app"

func main() {
	app.CreateDB("pretest")
	app := app.New(app.Config{
		DbUser: "root",
		DbName:    "pretest",
		DbDialect: "mysql",
	})


	app.Logger.Printf("Generating tables from seed")
	app.RunSeeder()
	app.Logger.Printf("Ok")

}
