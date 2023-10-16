package main

import (
	"fmt"
	"net/http"
)

/*
# Routing Static Assets

Pada bagian ini kita akan belajar bagaimana cara routing static assets atau static contents.
Seperti file css, js, gambar, umumnya dikategorikan sebagai static assets.

learn more: https://dasarpemrogramangolang.novalagung.com/B-routing-static-assets.html
*/

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello again"))
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./b3-routing-static-assets/assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
