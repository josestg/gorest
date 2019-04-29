package app

import "database/sql"

type Seed struct {
	Products []Product
	Categories  []Category
	Images []Image
	CategoryProducts []CategoryProduct
	ProductImages []ProductImage
}

func (A *App) GetSeed() *Seed{
	return &Seed{
		Products:[]Product{
			{Name:"Statistika",Enable:true,Description:"Buku pengantar statistika dasar"},
			{Name:"Pemograman Golang",Enable:true,Description:"Buku pengantar pemograman"},
			{Name:"Aljabar Linear",Enable:true,Description:"Buku Kuliah Teknik Informatika"},
			{Name:"Algoritma dan Struktur Data",Enable:true,Description:"Buku kuliah semester 3"},
			{Name:"Machine Learning",Enable:true,Description:"Buku praktik dan teori machine learning"},
			{Name: "Komputasi Geofisia", Enable: true, Description: "Buku komputasi mahasiswa geofisika"},
			{Name: "Matematika Diskrit", Enable: true, Description: "Buku mahasiswa ilmu komputer"},
		},
		Categories:[]Category{
			{Name:"Matematika",Enable:true},
			{Name:"Geofisika",Enable:true},
			{Name:"Informatika",Enable:true},
			{Name:"Sains",Enable:true},
		},
		Images:[]Image{
			{Name:"Image 1",File:"upload-121371.png"},
			{Name:"Image 2",File:"upload-121362.png"},
			{Name:"Image 3",File:"upload-121353.png"},
			{Name:"Image 4",File:"upload-121344.png"},
			{Name:"Image 5",File:"upload-121335.png"},
			{Name:"Image 6",File:"upload-121326.png"},
			{Name:"Image 8",File:"upload-121317.png"},
		},
		CategoryProducts:[]CategoryProduct{
			{1,1},
			{1,3},
			{1,4},
			{2,3},
			{6,2},
			{5,3},
			{5,1},
		},
		ProductImages:[]ProductImage{
			{1,1},
			{2,2},
			{1,4},
			{2,3},
			{6,2},
			{5,3},
			{5,1},
		},

	}
}



func CreateDB(name string){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_,_ = db.Exec("DROP DATABASE "+name)
	_,_ = db.Exec("CREATE DATABASE "+name)
}

func (A *App) RunSeeder() {
	seed := A.GetSeed()
	for _,x := range seed.Products{A.Db.Save(&x)}
	for _,x := range seed.Categories{A.Db.Save(&x)}
	for _,x := range seed.Images{A.Db.Save(&x)}
	for _,x := range seed.CategoryProducts{A.Db.Save(&x)}
	for _,x := range seed.ProductImages{A.Db.Save(&x)}
}