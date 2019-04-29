package app

import (
	"net/http"
)

func (A *App) ProductImageWrapper(res []ProductImage) *[]ProductImagesResponse {

	return nil
}

// GetProducts : GET /api/product-images
func (A *App) GetProductImages(w http.ResponseWriter, r *http.Request){


}

// CreateProduct : POST /api/product-images
func (A *App) CreateProductImage(w http.ResponseWriter, r *http.Request){


}

// GetProduct : GET /api/product-images/pid/iid
func (A *App) GetProductImage(w http.ResponseWriter, r *http.Request){

}

// UpdateProduct : PUT /api/product-images/pid/iid
func (A *App) UpdateProductImage(w http.ResponseWriter, r *http.Request){

}

// DeleteProduct : Delete /api/product-images/pid/iid
func (A *App)DeleteProductImage(w http.ResponseWriter, r *http.Request)  {

}