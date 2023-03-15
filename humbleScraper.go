package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func GetTitles(url string) []string {
	var sliceTitles []string
	// Realizar solicitud GET
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Leer cuerpo de respuesta
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	htmlString := string(html)

	re := regexp.MustCompile(`"image_text": "([^"]+)"`)

	// Encontrar todas las coincidencias
	matches := re.FindAllStringSubmatch(htmlString, -1)

	// Imprimir cada coincidencia
	for _, match := range matches {
		//fmt.Println(match[1])
		sliceTitles = append(sliceTitles, match[1])
	}

	return sliceTitles
}

func main() {
	// URL de la página web que queremos extraer el HTML
	url := "https://www.humblebundle.com/books/cybersecurity-packt-2023-books"
	titles := GetTitles(url)
	//fmt.Println(titles)
	fmt.Print(titles[3])

}
