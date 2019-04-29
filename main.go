package main

import "github.com/josestnggng/Pretest-privy-full-rest-api/app"

func main() {

	app := app.New(app.Config{
		DbUser: "root",
		DbName:    "pretest",
		DbDialect: "mysql",
	})
	app.SetupRouter()
	app.Run(":8000")

}
