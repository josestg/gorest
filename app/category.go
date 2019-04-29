package app

import (
	"encoding/json"
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
	var res Category
	var decoder = json.NewDecoder(r.Body)
	if err := decoder.Decode(&res); err!=nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()
	if err:= A.Db.Save(&res).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w,http.StatusOK, &Response{
		Success:true,
		Data:res,
	})
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
