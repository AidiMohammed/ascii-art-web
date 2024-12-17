package main

import (
	"fmt"
	"net/http"

	"ascii-art-web/controllers"
)

func main() {
	
	http.HandleFunc("/", controllers.FormAsciiArt)//GET
	http.HandleFunc("/ascii-art", controllers.GenerateAsciiArt)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../front-end/static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../front-end/static/js"))))

	fmt.Println("starting server on port :8081 ...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
