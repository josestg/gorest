package app

import (
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