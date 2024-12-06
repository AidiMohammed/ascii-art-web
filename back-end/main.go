package main

import (
	"fmt"
	"net/http"
	"ascii-art-web/controllers"
)

func main(){

	http.HandleFunc("/",controllers.FormAsciiArt)// from méthode GET
	http.HandleFunc("/ascii-art",controllers.GenerateAsciiArt)// from méthode POST

	http.Handle("/css/",http.StripPrefix("/css/", http.FileServer(http.Dir("../front-end/static/css"))))
	http.Handle("/js/",http.StripPrefix("/js/", http.FileServer(http.Dir("../front-end/static/js"))))
   
	fmt.Println("starting server port :8081 ...")
	http.ListenAndServe(":8081",nil)
}