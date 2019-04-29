package app

import (
	"encoding/json"
	"net/http"
)

func (A *App) ResponseCategoryProductWrapper(res []CategoryProduct) *[]CategoryProductResponse {
	// mengumpulkan kategory suatu product
	index := map[uint][]uint{}
	for _,data := range res{
		s := index[data.ProductID]
		index[data.ProductID] = append(s, data.CategoryID)
	}
	//membungkus data
	final := []CategoryProductResponse{}
	for k,values := range index{
		data := CategoryProductResponse{}

		p := Product{}
		if err :=A.Db.Find(&p, Product{ID: k}).Error; err !=nil {return nil}

		data.Product = p
		for _,v := range values{
			c := Category{}
			if err :=A.Db.Find(&c, Category{ID: uint(v)}).Error; err !=nil {return nil}
			data.Category = append(data.Category, c)
		}

		final = append(final,data)
	}

	return &final
}

// GetProducts : GET /api/category-products
func (A *App) GetCategoryProducts(w http.ResponseWriter, r *http.Request){
	var res []CategoryProduct
	if err :=A.Db.Find(&res).Error; err !=nil {
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	final := A.ResponseCategoryProductWrapper(res)
	A.RespondJSON(w,http.StatusOK,final)
}

// CreateProduct : POST /api/category-products
func (A *App) CreateCategoryProduct(w http.ResponseWriter, r *http.Request){
	var res CategoryProduct
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

// GetProduct : GET /api/category-products/pid/cid
func (A *App) GetCategoryProduct(w http.ResponseWriter, r *http.Request){
	var res []CategoryProduct
	var params = getVars(r)
	pid,err := parserID(params["pid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	cid,err := parserID(params["cid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Find(&res, CategoryProduct{ProductID: pid,CategoryID:cid}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	final := A.ResponseCategoryProductWrapper(res)

	A.RespondJSON(w,http.StatusOK,final)
}

// UpdateProduct : PUT /api/category-products/pid/cid
func (A *App) UpdateCategoryProduct(w http.ResponseWriter, r *http.Request){
	var res CategoryProduct
	var decoder = json.NewDecoder(r.Body)
	var params = getVars(r)
	pid,err := parserID(params["pid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	cid,err := parserID(params["cid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Delete(&res, CategoryProduct{ProductID: pid,CategoryID:cid}).Error; err!=nil{return}
	if err:= decoder.Decode(&res); err!=nil{
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	defer r.Body.Close()

	if err:=A.Db.Save(&res).Error; err!=nil{
		A.RespondError(w, http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w, http.StatusOK,&Response{
		Success:true,
		Data:res,
	})
}

// DeleteProduct : Delete /api/category-products/pid/cid
func (A *App) DeleteCategoryProduct(w http.ResponseWriter, r *http.Request){
	var res CategoryProduct
	var params = getVars(r)
	pid,err := parserID(params["pid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	cid,err := parserID(params["cid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Find(&res, CategoryProduct{ProductID: pid,CategoryID:cid}).Error; err!=nil{return}
	if err:= A.Db.Delete(&res, CategoryProduct{ProductID: pid,CategoryID:cid}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w, http.StatusOK,&Response{
		Success:true,
		Data:res,
	})
}

