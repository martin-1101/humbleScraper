# HumbleScrapper

Este es un programa en Go que permite obtener los títulos de los paquetes de libros de Humble Bundle.

## Uso

El programa se ejecuta mediante la línea de comandos y requiere una única bandera `-u` que debe contener la URL del paquete de libros de Humble Bundle del que se desean obtener los títulos.

Ejemplo:
```sh
go run scraper.go -u https://www.humblebundle.com/books/pack-que-quieres
```

El programa devolverá una lista numerada con los títulos de los libros incluidos en el paquete.

## Motivación

El objetivo principal de este programa es aprender sobre el lenguaje de programación Go y mejorar habilidades de programación en general.
