package app

import (
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

