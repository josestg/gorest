package app

import (
	"net/http"
)

func (A *App) ResponseCategoryProductWrapper(res []CategoryProduct) *[]CategoryProductResponse {

	return nil
}

// GetProducts : GET /api/category-products
func (A *App) GetCategoryProducts(w http.ResponseWriter, r *http.Request){

}

// CreateProduct : POST /api/category-products
func (A *App) CreateCategoryProduct(w http.ResponseWriter, r *http.Request){


}

// GetProduct : GET /api/category-products/pid/cid
func (A *App) GetCategoryProduct(w http.ResponseWriter, r *http.Request){

}

// UpdateProduct : PUT /api/category-products/pid/cid
func (A *App) UpdateCategoryProduct(w http.ResponseWriter, r *http.Request){

}

// DeleteProduct : Delete /api/category-products/pid/cid
func (A *App) DeleteCategoryProduct(w http.ResponseWriter, r *http.Request){

}

