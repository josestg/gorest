package app

func (A *App) SetupRouter(){

	// Product Endpoints
	A.Get("/api/products",A.GetProducts)
	A.Post("/api/products",A.CreateProduct)
	A.Get("/api/products/{id}",A.GetProduct)
	A.Put("/api/products/{id}",A.UpdateProduct)
	A.Del("/api/products/{id}",A.DeleteProduct)

	// Category Endpoints
	A.Get("/api/categories",A.GetCategories)
	A.Post("/api/categories",A.CreateCategory)
	A.Get("/api/categories/{id}",A.GetCategory)
	A.Put("/api/categories/{id}",A.UpdateCategory)
	A.Del("/api/categories/{id}",A.DeleteCategory)

	// Image Endpoints
	A.Get("/api/images",A.GetImages)
	A.Post("/api/images",A.CreateImage)
	A.Get("/api/images/{id}",A.GetImage)
	A.Put("/api/images/{id}",A.UpdateImage)
	A.Del("/api/images/{id}",A.DeleteImage)

	// Category Product Endpoints
	A.Get("/api/category-products",A.GetCategoryProducts)
	A.Post("/api/category-products",A.CreateCategoryProduct)
	A.Get("/api/category-products/{pid}/{cid}",A.GetCategoryProduct)
	A.Put("/api/category-products/{pid}/{cid}",A.UpdateCategoryProduct)
	A.Del("/api/category-products/{pid}/{cid}",A.DeleteCategoryProduct)

}