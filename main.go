package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", holaMundo)
	http.HandleFunc("/cuadrado/", manejadorCuadrado)
	http.HandleFunc("/cubo/", manejadorCubo)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func holaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}

func manejadorCuadrado(w http.ResponseWriter, r *http.Request) {
	n := extractNumber(r.URL.Path)
	result := calcularCuadrado(n)
	fmt.Fprintf(w, "Cuadrado de %d es %d", n, result)
}

func manejadorCubo(w http.ResponseWriter, r *http.Request) {
	n := extractNumber(r.URL.Path)
	result := calcularCubo(n)
	fmt.Fprintf(w, "Cubo de %d es %d", n, result)
}

func extractNumber(path string) int {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		n, err := strconv.Atoi(parts[2])
		if err != nil {
			return 0 // O maneja el error adecuadamente según tu lógica de aplicación
		}
		return n
	}
	return 0
}
