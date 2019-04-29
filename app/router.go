package app

func (A *App) SetupRouter(){

	// Product Endpoints
	A.Get("/api/products",A.GetProducts)
	A.Post("/api/products",A.CreateProduct)
	A.Get("/api/products/{id}",A.GetProduct)
	A.Put("/api/products/{id}",A.UpdateProduct)
	A.Del("/api/products/{id}",A.DeleteProduct)

}