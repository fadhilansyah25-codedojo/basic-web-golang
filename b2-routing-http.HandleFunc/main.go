package main

import (
	"fmt"
	"net/http"
)

/*
# Routing http.HandleFunc

Dalam Go, routing bisa dilakukan dengan beberapa cara, di antaranya:
  - Dengan memanfaatkan fungsi http.HandleFunc()
  - Mengimplementasikan interface http.Handler pada suatu struct, untuk kemudian
    digunakan pada fungsi http.Handle()
  - Membuat multiplexer sendiri dengan memanfaatkan struct http.ServeMux
  - Dan lainnya

learn more: https://dasarpemrogramangolang.novalagung.com/B-routing-http-handlefunc.html
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

	address := ":9000"
	fmt.Printf("server started at localhost:%s", address)
	http.ListenAndServe(address, nil)
}
