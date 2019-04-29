package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct{
	Router *mux.Router	
	Logger *log.Logger
	Db *gorm.DB
}

func GetDB(uri, dialect string) *gorm.DB{
	var db, err = gorm.Open(dialect,uri)
	if err !=nil{
		panic(fmt.Errorf("Error while connectin to DB %s\n",
			err.Error()))
		return nil
	}
	db.AutoMigrate(&Product{},&Image{},&Category{},
		&CategoryProduct{},&ProductImage{})
	return db
}

func New(c Config) *App{

	logger := log.New(os.Stdout,"Log : ",log.LstdFlags| log.Lshortfile)
	uri := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		c.DbUser,c.DbPass,c.DbName)

	var app App
	app.Router = mux.NewRouter()
	app.Logger = logger
	app.Db = GetDB(uri,c.DbDialect)
	return &app
}

func (A *App) Run(port string){
	A.Logger.Printf("App start on port \t\t%s\n",port)
	if err  := http.ListenAndServe(port, A.Router);
		err !=nil {
		panic("Port not available!")
	}
}

// Utility
func getVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// MIDDLEWARE
// Log
func (A *App) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime :=time.Now()
		defer A.Logger.Printf("%s %s\t : Request processed in %s\n",
			r.Method, r.URL.Path,time.Now().Sub(startTime),
		)
		next(w,r)
	}
}

// WRAPPER METHOD
func (A *App) Get(route string, fn func(w http.ResponseWriter, r *http.Request)){
	A.Router.HandleFunc(route,
		A.Log(fn)).Methods("GET")
}

func (A *App) Post(route string, fn func(w http.ResponseWriter, r *http.Request)){
	A.Router.HandleFunc(route,
		A.Log(fn)).Methods("POST")
}

func (A *App) Put(route string, fn func(w http.ResponseWriter, r *http.Request)){
	A.Router.HandleFunc(route,
		A.Log( fn)).Methods("PUT")
}

func (A *App) Del(route string, fn func(w http.ResponseWriter, r *http.Request)){
	A.Router.HandleFunc(route,
		A.Log(fn)).Methods("DELETE")
}


