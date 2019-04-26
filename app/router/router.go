package router

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Router struct {
	Mux *mux.Router
	DB *gorm.DB
}

func (R *Router) New()  {
	R.Mux = mux.NewRouter()
}

func (R *Router ) Get(route string, fn func(w http.ResponseWriter, r *http.Request)){
	R.Mux.HandleFunc(route,fn).Methods("GET")
}

func (R *Router ) Post(route string, fn func(w http.ResponseWriter, r *http.Request)){
	R.Mux.HandleFunc(route,fn).Methods("POST")
}

func (R *Router ) Put(route string, fn func(w http.ResponseWriter, r *http.Request)){
	R.Mux.HandleFunc(route,fn).Methods("PUT")
}

func (R *Router ) Del(route string, fn func(w http.ResponseWriter, r *http.Request)){
	R.Mux.HandleFunc(route,fn).Methods("DELETE")
}
