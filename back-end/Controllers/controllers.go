package controllers

import (
	"os"
	"fmt"
	"net/http"
	"html/template"
	"ascii-art-web/tols"
)

type data struct{
	TitlePage string
	Banner string
	InputUser string
	AsciiArt string
	Error string
	Style bool
}

func modeRun ()bool{
	withStyle := false

	if len(os.Args) == 2 {
		if os.Args[1] == "--with-style" || os.Args[1] == "-ws"{
			withStyle = true
		}
	}

	return withStyle
}

func FormAsciiArt(w http.ResponseWriter, r *http.Request){
	withStyle := modeRun()
	
	template,err := template.ParseFiles("../front-end/template/formAsciiArt.html")

	if err != nil {
		http.Error(w , err.Error(), http.StatusInternalServerError)
		return
	}

	dataPage := data{
		TitlePage: "Ascii Art Web",
		Style: withStyle,
	}

	err = template.Execute(w, dataPage)
	if err != nil {
		fmt.Println("ERR")
		http.Error(w , err.Error(), http.StatusInternalServerError)
	} else {

	}
}



func GenerateAsciiArt(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {

		withStyle := modeRun()

		r.ParseForm()
		banner := r.FormValue("banner")
		inputUser := r.FormValue("textarea")

		dataPage := data{
			TitlePage: "Show Ascii art",
			Banner: banner,
			InputUser: inputUser,
			Style: withStyle,
		}

		//check input user and name banner
		if inputUser == "" || !tols.CheckNameBanner(banner) {

			if inputUser == ""{
				dataPage.Error= "The field text is required"
				
			} else if !tols.CheckNameBanner(banner) {
				dataPage.Error = "Invalid style selected"
			}

			template,err := template.ParseFiles("../front-end/template/formAsciiArt.html")

			if err != nil {
				http.Error(w,err.Error(), http.StatusInternalServerError)
				return
			}
			template.Execute(w,dataPage)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		dataPage.AsciiArt = tols.AsciArt(inputUser,banner)
		//fmt.Println(dataPage.AsciiArt)

		template,err := template.ParseFiles("../front-end/template/resultAsciiArt.html")

		if err != nil {
			http.Error(w,err.Error(), http.StatusInternalServerError)
			return
		}

		err = template.Execute(w,dataPage)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}	
}