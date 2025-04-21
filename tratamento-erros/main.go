package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	res, err := http.Get("http://google.com.br")
	if err != nil {
		log.Fatal("Erro no request:", err.Error())
	}

	fmt.Println("Status Code:", res.Header)

	fmt.Println("------------------")
	resultado, err := soma(10, 5)
	if err != nil {
		log.Fatal("Erro na soma:", err.Error())
	}
	fmt.Println("Resultado da soma:", resultado)
}

func soma(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("não pode ser negativo")
	}
	return a + b, nil
}