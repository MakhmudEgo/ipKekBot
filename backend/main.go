package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"net/http"
)

//var dbase *gorm.DB
//
//func Init() *gorm.DB {
//	db, err := gorm.Open("postgres", "users=posgres password-admin")
//}
//
//func getDB() *gorm.DB {
//
//}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Form)
		io.WriteString(w, "hello world\n")
		fmt.Fprintf(w, "%#v", r.URL.Query())
	})
	//Init()
	dsn := "host=192.168.235.110 users=postgres password=lol dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("no connect")
	}
	db.DB()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
