package main

import "myprestest/app"

func main() {
	app := app.New(app.Config{
		DbUser: "root",
		DbName:    "pretest",
		DbDialect: "mysql",
	})

	app.Run(":8000")

}
