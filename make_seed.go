package main

import "myprestest/app"

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
