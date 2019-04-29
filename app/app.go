package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
)

type App struct{
	Router *mux.Router	
	Logger *log.Logger
	Db *gorm.DB
}

func GetDB(){

}

func New(){

}

func (A *App) Run(port string){

}

// MIDDLEWARE


// WRAPPER METHOD

