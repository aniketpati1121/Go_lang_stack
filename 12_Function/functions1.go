package main

import "fmt"

func getLanguages() (string, string, string, bool) {

	return "golang", "python", "java", true
}

func main() {

	lang1, lang2, lang3, _ := getLanguages()
	fmt.Println(lang1, lang2, lang3)

} 


# go run 12_Function/functions1.go
