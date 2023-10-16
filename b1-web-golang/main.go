package main

import (
	"fmt"
	"net/http"
)

/*
# Golang Web App: Hello World

Pada serial chapter B ini, kita masih tetap akan belajar tentang topik fundamental atau dasar tapi
lebih spesifik ke arah web development. Kita awali dengan pembahasan bagaimana cara membuat aplikasi
web "Hello World" sederhana menggunakan Go.

learn more: https://dasarpemrogramangolang.novalagung.com/B-golang-web-hello-world.html
*/

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World!"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// var address = "localhost:9000"
	// fmt.Printf("server started at %s\n", address)
	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
