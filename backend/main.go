package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type user struct {
	ID       int
	Username string
}

type UserHistory struct {
	Query  string
	Userid int `gorm:"column:user_id"`
}

func main() {
	var usrs []user
	dsn := "host=bd user=postgres password=lol dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("backend no connect: " + err.Error())
	}

	http.HandleFunc("/get_users", func(w http.ResponseWriter, r *http.Request) {
		db.Select("id", "username").Find(&usrs)
		js, err := json.Marshal(usrs)
		if err != nil {
			log.Fatal("get_users: " + err.Error())
		}
		fmt.Fprint(w, string(js))
	})

	http.HandleFunc("/get_user", func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("id")

		res := db.Where("id = ?", userId).Find(&usrs)
		if res.Error != nil {
			_, err := fmt.Fprintln(w, userId, "not found")
			if err != nil {
				log.Fatal("get_user: " + err.Error())
			}
			return
		}
		js, err := json.Marshal(usrs)
		if err != nil {
			log.Fatal("get_users: " + err.Error())
		}
		fmt.Fprint(w, string(js))
	})

	http.HandleFunc("/del_hist", func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Println(err.Error())
		}
		userQuery := r.URL.Query().Get("query")
		usrHist := UserHistory{Userid: userId, Query: userQuery}
		res := db.Where("query = ?", userQuery).Delete(&usrHist)
		if res.Error != nil {
			_, err := fmt.Fprintln(w, userId, "not found")
			if err != nil {
				log.Fatal("del_hist: " + err.Error())
			}
			return
		}
		fmt.Fprint(w, "success")
	})

	log.Fatal(http.ListenAndServe(":8999", nil))
}
