package app

import (
	"net/http"
	"./config"
	"./model"
	"./router"
)

var R *router.Router
var M *model.Model


//
func Listen(port string){
	
	R.Get("/api/categories",M.GetCategories)
	R.Post("/api/categories",M.CreateCategory)
	R.Get("/api/categories/{id}",M.GetCategory)
	R.Put("/api/categories/{id}",M.UpdateCategory)
	R.Del("/api/categories/{id}",M.DeleteCategory)

	
	R.Get("/api/products",M.GetProducts)
	R.Post("/api/products",M.CreateProduct)
	R.Get("/api/products/{id}",M.GetProduct)
	R.Put("/api/products/{id}",M.UpdateProduct)
	R.Del("/api/products/{id}",M.DeleteProduct)

	http.ListenAndServe(port,R.Mux)

}


func Run(port string){
	DB := config.Connect()
	R = &router.Router{}
	M = &model.Model{}
	M.New(DB)
	R.New()

	Listen(port)
}