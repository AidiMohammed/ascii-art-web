package tols

import (
	"fmt"
	"bufio"
	"os"
)

func AsciArt(inputUser string, banner string) string{
	fmt.Println("input usr : "+inputUser)
	fmt.Println("banner : "+banner)
	var valueRuturn string
	var index int

	strToReuns := []rune(inputUser)

	myResultAscii ,_:= LoadBanner(banner)

	for i := 0 ; i < 8; i++ {
		for index = 0; index < len(strToReuns); index++{
			_, ok := myResultAscii[int(strToReuns[index])]
			if ok {
				valueRuturn += myResultAscii[int(strToReuns[index])][i] 
			} else {
				if string(strToReuns[index]) == "\n"{
					fmt.Println("le retour à la ligne pas encore traiter")
					valueRuturn += "\n"
				} else {
					fmt.Println("la clé "+string(strToReuns[index])+" n'existe pas")
					valueRuturn += myResultAscii[32][index]
				}
			}
		}
		valueRuturn += "\n"
	}

	return valueRuturn
}

func CheckNameBanner(inputBanner string) bool {
	banners := []string{"standard","shadow","thinkertoy"}
	for _, banner := range banners {
		if inputBanner == banner{
			fmt.Println("banner is valid")
			return true
		}
	}
	fmt.Println("banner name is invalid")
	return false
}

func LoadBanner(banner string)(map[int][]string, error){
	result := make(map[int][]string)
	indexAscii := 32
	
	file, err := os.Open("banners/"+banner+".txt")

	if err != nil {
		fmt.Println("Erreur lors de l'ouverteur de ficher : ",err)
		return result, err
	}

	defer file.Close() //Fermer le ficher à la fine 

	//créer un scanner pour lire le ficher ligne par ligne
	scanner := bufio.NewScanner(file)
	NewChar := true

	for scanner.Scan(){
	
		line := scanner.Text() //lire une ligne

		if line == "" && !NewChar {
			NewChar = true
			indexAscii++
			continue
		}

		if NewChar {
			result[indexAscii] = []string{scanner.Text()}
			NewChar = false
		} else {
			if line != "" {
				result[indexAscii] = append(result[indexAscii], line)
			}
		}

	}
	return result,err
}