package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"ascii-art-web/tools"
)

type data struct {
	TitlePage string
	Banner    string
	InputUser string
	AsciiArt  string
	Error     string
	Style     bool
}

type ErrorData struct {
	CodeStatus 	int
	Message    	string
	Style 		bool
}

// for pointer for use file CSS or no by command (os.Args)
func modeRun() bool {
	withStyle := false

	if len(os.Args) == 2 {
		if os.Args[1] == "--with-style" || os.Args[1] == "-ws" {
			withStyle = true
		}
	}

	return withStyle
}

// Initialitation for home page
func FormAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		if r.URL.Path != "/" {
			ErrorPage(w, http.StatusNotFound, "Page not found !")
		}
		if r.Method != "GET" {
			ErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed !")
		}
		return
	}

	template, err := template.ParseFiles("../front-end/template/formAsciiArt.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "server intrernal error !")
		return
	}

	dataPage := data{
		TitlePage: "Ascii Art Web",
		Style:     modeRun(),
	}

	err = template.ExecuteTemplate(w, "formAsciiArt.html", dataPage)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "server intrernal error !")
		return
	}
}

// use data of user and return result in page of result
func GenerateAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		r.ParseForm()
		banner := r.FormValue("banner")
		inputUser := r.FormValue("textarea")

		dataPage := data{
			TitlePage: "Show Ascii art",
			Banner:    banner,
			InputUser: inputUser,
			Style:     modeRun(),
		}

		// check input user and name banner
		if inputUser == "" || !tools.CheckNameBanner(banner) {
			ErrorPage(w, http.StatusBadRequest, "Bad request !")
			return
		}

		dataPage.AsciiArt = tools.AsciArt(inputUser, banner)

		template, err := template.ParseFiles("../front-end/template/resultAsciiArt.html")
		if err != nil {
			ErrorPage(w, http.StatusInternalServerError, "server intrernal error !")
			return
		}

		if dataPage.AsciiArt == "" {
			ErrorPage(w, http.StatusInternalServerError, "server intrernal error !")
			return
		}
		err = template.Execute(w, dataPage)
		if err != nil {

			ErrorPage(w, http.StatusInternalServerError, "server intrernal error !")
			return
		}
	} else {
		ErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed !")
		return
	} 
}

func ErrorPage(w http.ResponseWriter, statutcode int, message string) {

	template, err := template.ParseFiles("../front-end/template/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statutcode)
	data := ErrorData{
		CodeStatus: statutcode,
		Message:    message,

		Style: modeRun(),
	}

	err = template.ExecuteTemplate(w, "error.html", data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
