package main

import (
	"fmt"
	"net/http"
	"ascii-art-web/controllers"
	"os"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/ascii-art", controllers.GenerateAsciiArt)

    /*
        mux.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
            panic("test for r 500")
        })
    */

    mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../front-end/static/css"))))
    mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../front-end/static/js"))))

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
            w.WriteHeader(http.StatusNotFound)
            http.ServeFile(w, r, "../front-end/template/404.html")
            return
        }

        if r.URL.Path == "/" {
            controllers.FormAsciiArt(w, r)
        }
    })

    fmt.Println("starting server on port :8081 ...")
    if err := http.ListenAndServe(":8081", errorHandler(mux)); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func errorHandler(h http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if rec := recover(); rec != nil {

                fmt.Fprintln(os.Stderr, "Recovered from panic:", rec)
                w.WriteHeader(http.StatusInternalServerError)
                http.ServeFile(w, r, "../front-end/template/500.html")
            }
        }()
        h.ServeHTTP(w, r)
    }
}