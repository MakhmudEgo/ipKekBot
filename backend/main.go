package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Form)
		io.WriteString(w, "hello world\n")
		fmt.Fprintf(w, "%#v", r.URL.Query())
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
