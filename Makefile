install :
	go get github.com/go-sql-driver/mysql 
	go get github.com/jinzhu/gorm
	go get github.com/gorilla/mux   
start :
	go run main.go

seed :
	go run make_seed.go