package model

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"../helper"
)

type Product struct {
	ID int 		`gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Enable bool `gorm:"not null" json:"enable"`
}


func (M *Model) GetProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := helper.ParseID(vars["id"])
	product := Product{}

	if err := M.DB.First(&product,Product{ID:id}).Error; err != nil {
		helper.RespondError(w, http.StatusNotFound,err.Error())
		return
	}
	helper.RespondJSON(w,http.StatusOK,product)
}

func (M *Model) CreateProduct(w http.ResponseWriter, r *http.Request){
	product := Product{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err!=nil {
		helper.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()

	if err := M.DB.Save(&product).Error ; err!=nil{
		helper.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	helper.RespondJSON(w,http.StatusOK,product)

}

func (M *Model) DeleteProduct(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := helper.ParseID(vars["id"])
	product := Product{}

	if err := M.DB.First(&product,Product{ID:id}).Error; err != nil {return}
	if &product == nil{return}
	if err := M.DB.Delete(&product,Product{ID:id}).Error ; err !=nil {
		helper.RespondError(w,http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondJSON(w,http.StatusOK,map[string]bool{"success": true})
}

func (M *Model) UpdateProduct(w http.ResponseWriter, r *http.Request){

	vars :=mux.Vars(r)
	id:= helper.ParseID(vars["id"])
	product := Product{}
	decoder := json.NewDecoder(r.Body)


	if err := M.DB.First(&product,Product{ID:id}).Error; err != nil {return}
	if err := decoder.Decode(&product); err!=nil {
		helper.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()

	// Menjaga agar ID Tidak berubah
	product.ID = id
	if err := M.DB.Save(&product).Error; err!=nil {
		helper.RespondError(w, http.StatusInternalServerError,err.Error())
		return
	}
	helper.RespondJSON(w, http.StatusOK, product)
}



func (M *Model) GetProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	products := []Product{}

	if err := M.DB.Find(&products).Error ; err!=nil{
		helper.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}
	helper.RespondJSON(w,http.StatusOK,products)
}
