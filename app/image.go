package app

import (
	"io/ioutil"
	"net/http"
	"os"
)
// Upload Image  to dir Uploads
func UploadImage(filename *string,r *http.Request) error  {
	//max-size 10MB
	if err := r.ParseMultipartForm(10<<20); err!=nil{return err}
	file,_,err := r.FormFile("file")
	if err!= nil {return err}
	defer file.Close()

	temp,err := ioutil.TempFile("uploads","upload-*.png")
	if err!=nil {return err}
	defer temp.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err!= nil {return err}
	var _,_ = temp.Write(fileBytes)

	*filename = temp.Name()
	return nil
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
	r.ParseMultipartForm(10<<20)
	var res Image
	enable , err := parserEnable(r.Form.Get("enable"))
	if err!=nil{
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}

	var filename string
	if err:=UploadImage(&filename,r);err!=nil{
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}

	res.Enable = enable
	res.Name = r.Form.Get("name")
	res.File = filename

	if err:= A.Db.Save(&res).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w,http.StatusOK, &Response{
		Success:true,
		Data:res,
	})


}

// GetImage : GET /api/images/id
func (A *App) GetImage(w http.ResponseWriter, r *http.Request){
	var res Image
	var params = getVars(r)
	var id,err = parserID(params["id"])
	if err!=nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.First(&res, Image{ID: id}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w,http.StatusOK,res)
}

// UpdateImage : PUT /api/images/id
func (A *App) UpdateImage(w http.ResponseWriter, r *http.Request){
	var res Image
	var params = getVars(r)
	var id,err = parserID(params["id"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Find(&res, Image{ID: id}).Error; err!=nil{return}

	r.ParseMultipartForm(10<<20)

	var filename = res.File
	file ,_,_ := r.FormFile("file");
	if file!=nil {
		if err:=UploadImage(&filename,r);err!=nil{
			A.RespondError(w,http.StatusBadRequest,err.Error())
			return
		}
		defer file.Close()
	}


	enable , err := parserEnable(r.Form.Get("enable"))
	if err!=nil{
		A.RespondError(w,http.StatusBadRequest,err.Error())
		return
	}

	res.Enable = enable
	res.Name = r.Form.Get("name")
	res.File = filename
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

// DeleteImage : Delete /api/images/id
func (A *App) DeleteImage(w http.ResponseWriter, r *http.Request){
	var res Image
	var params = getVars(r)
	var id,err = parserID(params["id"])
	if err != nil{
		A.RespondError(w, http.StatusBadRequest,err.Error())
		return
	}
	if err:= A.Db.Find(&res, Image{ID: id}).Error; err!=nil{return}

	// hapus kode berikut jika tidak ada file (keperluan test)
	var path = res.File
	if err := os.Remove(path); err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}
	// end delete file
	
	
	if err:= A.Db.Delete(&res, Image{ID: id}).Error; err!=nil{
		A.RespondError(w,http.StatusInternalServerError,err.Error())
		return
	}

	A.RespondJSON(w, http.StatusOK,&Response{
		Success:true,
		Data:res,
	})
}
