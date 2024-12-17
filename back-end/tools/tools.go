package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
)

// elements in strings no accept code ASCII
func findNonASCII(input string) []rune {
	var nonASCII []rune
	for _, char := range input {
		if char < 0 || char > 127 {
			nonASCII = append(nonASCII, char)
		}
	}
	return nonASCII
}

// transfer text to ASCI ART by banner of input (chose)
func AsciArt(inputUser string, banner string) string {
	fmt.Println("input usr : " + inputUser)
	fmt.Println("banner : " + banner)
	var valueRuturn string
	lignes := strings.Split(inputUser, "\n")

	/*if len(lignes) == 0 {
		return AsciArtOneWord(inputUser, banner)
	}*/
	
	for i := 0; i < len(lignes); i++ {

		valueRuturn += AsciArtOneWord(lignes[i], banner)
		if i != (len(lignes) - 1) {
			valueRuturn += "\n"
		}
	}
	return valueRuturn
}

// transfer part of text Input (split is by new ligne)
func AsciArtOneWord(inputUser string, banner string) string {
	var valueRuturn string
	var index int

	strToReuns := []rune(inputUser)

	myResultAscii, err := LoadBanner(banner)

	if err == nil {
		for i := 0; i < 8; i++ {
			for index = 0; index < len(strToReuns); index++ {
				_, ok := myResultAscii[int(strToReuns[index])]
				if ok {
					valueRuturn += myResultAscii[int(strToReuns[index])][i]
				} else {
					valueRuturn += myResultAscii[32][i]
				}
			}
			valueRuturn += "\n"
		}

		if len(findNonASCII(inputUser)) != 0 {
			fmt.Printf("\nnon-ASCII characters: %q\n", findNonASCII(inputUser))
		}
		return valueRuturn
	}else {
		fmt.Println("Une erreur est survenue lors du chargement de la banner. Veuillez réessayer ultérieurement.")
		return ""
	}
}

// check banner is correct name (found...)
func CheckNameBanner(inputBanner string) bool {
	banners := []string{"standard", "shadow", "thinkertoy"}
	for _, banner := range banners {
		if inputBanner == banner {
			fmt.Println("banner is valid")
			return true
		}
	}
	fmt.Println("banner name is invalid")
	return false
}

// down info file of banner, return map (code ASCII,string of code)
func LoadBanner(banner string) (map[int][]string,error) {
	result := make(map[int][]string)
	indexAscii := 32

	file, err := os.Open("banners/" + banner + ".txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverteur de ficher : ", err)
		return result,err
	}else{
		defer file.Close() // close file in end

		// create scanner for read file ligne by ligne:
		scanner := bufio.NewScanner(file)
		NewChar := true

		for scanner.Scan() {

			line := scanner.Text() // read ligne

			if line == "" && !NewChar{
				NewChar = true
				indexAscii++
				continue
			}

			if NewChar {//create new key end value in map
				result[indexAscii] = []string{line}
				NewChar = false
			} else {//update key

					result[indexAscii] = append(result[indexAscii], line)
				
			}

		}
		fmt.Println(len(result[33]))
		return result,err		
	}


}
