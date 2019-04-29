package app

import (
	"encoding/json"
	"net/http"
)

func (A *App) ProductImageWrapper(res []ProductImage) *[]ProductImagesResponse {
	// mengumpulkan kategory suatu product
	index := map[uint][]uint{}
	for _,data := range res{
		s := index[data.ProductID]
		index[data.ProductID] = append(s, data.ImageID)
	}
	//membungkus data
	var final []ProductImagesResponse
	for k,values := range index{
		data := ProductImagesResponse{}

		p := Product{}
		if err :=A.Db.Find(&p, Product{ID: k}).Error; err !=nil {return nil}

		data.Product = p
		for _,v := range values{
			c := Image{}
			if err :=A.Db.Find(&c, Image{ID: uint(v)}).Error; err !=nil {return nil}
			data.Image = append(data.Image, c)
		}

		final = append(final,data)
	}

	return &final
}

// GetProducts : GET /api/product-images
func (A *App) GetProductImages(w http.ResponseWriter, r *http.Request){

	var res []ProductImage
	if err :=A.Db.Find(&res).Error; err !=nil {
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}

	final := A.ProductImageWrapper(res)
	A.RespondJSON(w,http.StatusOK,final)
}

// CreateProduct : POST /api/product-images
func (A *App) CreateProductImage(w http.ResponseWriter, r *http.Request){

	var res ProductImage
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

// GetProduct : GET /api/product-images/pid/iid
func (A *App) GetProductImage(w http.ResponseWriter, r *http.Request){
	var res []ProductImage
	var params = getVars(r)

	pid,err := parserID(params["pid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	iid,err := parserID(params["iid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	if err:= A.Db.Find(&res, ProductImage{ProductID: pid,ImageID:iid}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	final := A.ProductImageWrapper(res)
	A.RespondJSON(w,http.StatusOK,final)
}

// UpdateProduct : PUT /api/product-images/pid/iid
func (A *App) UpdateProductImage(w http.ResponseWriter, r *http.Request){
	var res ProductImage
	var decoder = json.NewDecoder(r.Body)
	var params = getVars(r)
	pid,err := parserID(params["pid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	iid,err := parserID(params["iid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	if err:= A.Db.Delete(&res, ProductImage{ProductID: pid,ImageID:iid}).Error; err!=nil{return}
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

// DeleteProduct : Delete /api/product-images/pid/iid
func (A *App)DeleteProductImage(w http.ResponseWriter, r *http.Request)  {
	var res ProductImage
	var params = getVars(r)
	pid,err := parserID(params["pid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	iid,err := parserID(params["iid"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}

	if err:= A.Db.Find(&res, ProductImage{ProductID: pid,ImageID:iid}).Error; err!=nil{return}
	if err:= A.Db.Delete(&res, ProductImage{ProductID: pid,ImageID:iid}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w, http.StatusOK,&Response{
		Success:true,
		Data:res,
	})
}