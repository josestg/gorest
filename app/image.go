package app

import (
	"net/http"
)

func UploadImage(filename *string,r *http.Request)  {

}

// GetImages : GET /api/images
func (A *App) GetImages(w http.ResponseWriter, r *http.Request){
	var res []Image
	if err :=A.Db.Find(&res).Error; err !=nil {
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}
	A.RespondJSON(w,http.StatusOK,res)
}

// CreateImage : POST /api/images
func (A *App) CreateImage(w http.ResponseWriter, r *http.Request){


}

// GetImage : GET /api/images/id
func (A *App) GetImage(w http.ResponseWriter, r *http.Request){

}

// UpdateImage : PUT /api/images/id
func (A *App) UpdateImage(w http.ResponseWriter, r *http.Request){

}

// DeleteImage : Delete /api/images/id
func (A *App) DeleteImage(w http.ResponseWriter, r *http.Request){

}
