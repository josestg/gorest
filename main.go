package main

import "myprestest/app"

func main() {

	// Optional
	app.CreateDB("pretest")
	//
	app := app.New(app.Config{
		DbUser: "root",
		DbName:    "pretest",
		DbDialect: "mysql",
	})
	//optional
	app.RunSeeder()
	//
	app.SetupRouter()
	app.Run(":8000")

}
