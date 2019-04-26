package main

import (
	"fmt"
	"./app"
)

func main() {
	fmt.Println("App started on Port http://127.0.0.1:3000/api/")
	app.Run(":3000")
}
