package app

import (
	"net/http"
)
// GetCategories : GET /api/categories
func (A *App) GetCategories(w http.ResponseWriter, r *http.Request){
	var res []Category
	if err :=A.Db.Find(&res).Error; err !=nil {
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	A.RespondJSON(w,http.StatusOK,res)

}

// CreateCategory : POST /api/categories
func (A *App) CreateCategory(w http.ResponseWriter, r *http.Request){

}

// GetCategory : GET /api/categories/id
func (A *App) GetCategory(w http.ResponseWriter, r *http.Request){

}

// UpdateCategory : PUT /api/categories/id
func (A *App) UpdateCategory(w http.ResponseWriter, r *http.Request){

}

// DeleteCategory : Delete /api/categories/id
func (A *App) DeleteCategory(w http.ResponseWriter, r *http.Request){

}
