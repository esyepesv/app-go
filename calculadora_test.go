package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalcularCuadrado(t *testing.T) {
	if calcularCuadrado(5) != 25 {
		t.Errorf("CalcularCuadrado(5) should be 25")
	}
}

func TestCalcularCubo(t *testing.T) {
	if calcularCubo(3) != 27 {
		t.Errorf("CalcularCubo(3) should be 27")
	}
}

func TestHolaMundo(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(holaMundo)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hola Mundo\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestManejadorCuadrado(t *testing.T) {
	req, err := http.NewRequest("GET", "/cuadrado/4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(manejadorCuadrado)

	handler.ServeHTTP(rr, req)

	expected := "Cuadrado de 4 es 16"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestManejadorCubo(t *testing.T) {
	req, err := http.NewRequest("GET", "/cubo/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(manejadorCubo)

	handler.ServeHTTP(rr, req)

	expected := "Cubo de 3 es 27"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestExtractNumber(t *testing.T) {
	if num := extractNumber("/cuadrado/15"); num != 15 {
		t.Errorf("Expected 15, got %d", num)
	}
}
