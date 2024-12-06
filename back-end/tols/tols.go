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

	myResultAscii ,_:= LoadBannerv2(banner)

	for i := 0 ; i < 8; i++ {
		for index = 0; index < len(strToReuns); index++{
			valueRuturn += myResultAscii[int(strToReuns[index]+1)][i] 
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

func LoadBanner(banner string) (map[int]string, error){
	result := make(map[int]string)
	var tempCharAscii string
	indexAscii := 33
	
	//ouvrire le ficher en lecture
	file , err := os.Open("banners/"+banner+".txt")

	if err != nil{
		fmt.Println("Erreur lors de l'ouverteur de ficher : ",err)
		return result, err
	}
	defer file.Close() //Fermer le ficher à la fine 

	//créer un scanner pour lire le ficher ligne par ligne
	scanner := bufio.NewScanner(file)

	//Lire les lignes
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			result[indexAscii] = tempCharAscii
			indexAscii++
			tempCharAscii = ""
		}
		
		tempCharAscii += "\n"+line 
	}

	errScanner := scanner.Err()
	if errScanner != nil {
		fmt.Println("Erreur lors de la lecteur du fichier ", errScanner)
	}

	return result,err
}

func LoadBannerv2(banner string)(map[int][]string, error){
	result := make(map[int][]string)
	indexAscii := 33
	
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

	/*fmt.Println(result[33])
	fmt.Println(result[34])*/
	fmt.Println(len(result[33]))
	fmt.Println(len(result[34]))

	fmt.Println(" index map (33) ")
	for index,char := range result[33]{
		fmt.Printf("index : %v | char %v\n",index,char)
	}

	fmt.Println(" index map (34) ")
	for index,char := range result[34]{
		fmt.Printf("index : %v | char %v\n",index,char)
	}
	return result,err
}