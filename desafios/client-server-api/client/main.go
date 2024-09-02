package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatalf("Erro criando request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro fazendo request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro lendo response: %v", err)
	}

	fmt.Printf("Cotação recebida: %s\n", body)

	err = os.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", body)), 0644)
	if err != nil {
		log.Fatalf("Erro escrevendo no arquivo: %v", err)
	}
}
