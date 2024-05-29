package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello snipetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	log.Println("запуск веб-сервера на http://127.0.0.1:4000")
	err :=http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}
