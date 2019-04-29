package app

import (
	"net/http"
)
// GetProducts : GET /api/products
func (A *App) GetProducts(w http.ResponseWriter, r *http.Request){
	var res []Product
	if err :=A.Db.Find(&res).Error; err !=nil {
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	A.RespondJSON(w,http.StatusOK,res)
}

// CreateProduct : POST /api/products
func (A *App) CreateProduct(w http.ResponseWriter, r *http.Request){

}

// GetProduct : GET /api/products/id
func (A *App) GetProduct(w http.ResponseWriter, r *http.Request){

}

// UpdateProduct : PUT /api/products/id
func (A *App) UpdateProduct(w http.ResponseWriter, r *http.Request){

}

// DeleteProduct : Delete /api/products/id
func (A *App) DeleteProduct(w http.ResponseWriter, r *http.Request){

}
