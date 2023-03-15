package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func GetTitles(url string) ([]string, error) {
	// Realizar solicitud GET
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer cuerpo de respuesta
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	htmlString := string(html)

	// Expresion regular para encontrar los titulos
	re := regexp.MustCompile(`"image_text"\s*:\s*"([^"]+)"`)

	// Encontrar todas las coincidencias
	matches := re.FindAllStringSubmatch(htmlString, -1)

	// Si no se encontraron coincidencias
	if len(matches) == 0 {
		return nil, errors.New("no se encontraron títulos en la página")
	}

	// Construir slice de títulos
	var sliceTitles []string
	for _, match := range matches {
		sliceTitles = append(sliceTitles, match[1])
	}

	return sliceTitles, nil
}

func checkInputErrors(url string) error {
	urlRegex := regexp.MustCompile(`^https://www\.humblebundle\.com/books/.*$`)
	if url == "" {
		return errors.New("falta la URL a escrapear")
	}
	if !urlRegex.MatchString(url) {
		return errors.New(`se necesita una URL válida que empiece por: "https://www.humblebundle.com/books/" `)
	}
	return nil
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "Uso: %s -u <URL>\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "Opciones:")
	flag.PrintDefaults()
}

func main() {
	// La url de la pagina a scrapear se introducira mediante la flag -u y debera comenzar por https://www.humblebundle.com/books/
	var url string

	flag.StringVar(&url, "u", "", "URL a buscar")
	flag.Usage = showUsage
	flag.Parse()

	// Controlamos si no tiene ninguna flag
	if flag.NFlag() == 0 {
		showUsage()
		os.Exit(0)
	}

	// Comprobamos si hay errores en la entrada de usuario
	if err := checkInputErrors(url); err != nil {
		fmt.Fprintln(os.Stderr, err)
		showUsage()
		os.Exit(1)
	}

	titles, err := GetTitles(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for index, value := range titles {
		fmt.Println(index+1, value)
	}
}
