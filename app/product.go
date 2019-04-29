package app

import (
	"encoding/json"
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
	var res Product
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

// GetProduct : GET /api/products/id
func (A *App) GetProduct(w http.ResponseWriter, r *http.Request){
	var res Product
	var params = getVars(r)
	var id,err = parserID(params["id"])
	if err!=nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.First(&res, Product{ID: id}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w,http.StatusOK,res)
}

// UpdateProduct : PUT /api/products/id
func (A *App) UpdateProduct(w http.ResponseWriter, r *http.Request){
	var res Product
	var decoder = json.NewDecoder(r.Body)
	var params = getVars(r)
	var id,err = parserID(params["id"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Find(&res, Product{ID: id}).Error; err!=nil{return}
	if err:= decoder.Decode(&res); err!=nil{
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()

	// keep ID
	res.ID = id
	if err:=A.Db.Save(&res).Error; err!=nil{
		A.RespondError(w, http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w, http.StatusOK,&Response{
		Success:true,
		Data:res,
	})
}

// DeleteProduct : Delete /api/products/id
func (A *App) DeleteProduct(w http.ResponseWriter, r *http.Request){
	var res Product
	var params = getVars(r)
	var id,err = parserID(params["id"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Find(&res, Product{ID: id}).Error; err!=nil{return}
	if err:= A.Db.Delete(&res, Product{ID: id}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w, http.StatusOK,&Response{
		Success:true,
		Data:res,
	})
}
