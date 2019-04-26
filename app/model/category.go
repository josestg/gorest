package model

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"../helper"
)

type Category struct {
	ID int 		`gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Enable bool `gorm:"not null" json:"enable"`
}


func (M *Model) GetCategory(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := helper.ParseID(vars["id"])
	category := Category{}

	if err := M.DB.First(&category,Category{ID:id}).Error; err != nil {
		helper.RespondError(w, http.StatusNotFound,err.Error())
		return
	}
	helper.RespondJSON(w,http.StatusOK,category)
}

func (M *Model) CreateCategory(w http.ResponseWriter, r *http.Request){
	category := Category{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err!=nil {
		helper.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()

	if err := M.DB.Save(&category).Error ; err!=nil{
		helper.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	helper.RespondJSON(w,http.StatusOK,category)

}

func (M *Model) DeleteCategory(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := helper.ParseID(vars["id"])
	category := Category{}

	if err := M.DB.First(&category,Category{ID:id}).Error; err != nil {return}
	if &category == nil{return}
	if err := M.DB.Delete(&category,Category{ID:id}).Error ; err !=nil {
		helper.RespondError(w,http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondJSON(w,http.StatusOK,map[string]bool{"success": true})
}

func (M *Model) UpdateCategory(w http.ResponseWriter, r *http.Request){

	vars :=mux.Vars(r)
	id:= helper.ParseID(vars["id"])
	category := Category{}
	decoder := json.NewDecoder(r.Body)


	if err := M.DB.First(&category,Category{ID:id}).Error; err != nil {return}
	if err := decoder.Decode(&category); err!=nil {
		helper.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()

	// Menjaga agar ID Tidak berubah
	category.ID = id
	if err := M.DB.Save(&category).Error; err!=nil {
		helper.RespondError(w, http.StatusInternalServerError,err.Error())
		return
	}
	helper.RespondJSON(w, http.StatusOK, category)
}



func (M *Model) GetCategories(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	categories := []Category{}

	if err := M.DB.Find(&categories).Error ; err!=nil{
		helper.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}
	helper.RespondJSON(w,http.StatusOK,categories)
}
