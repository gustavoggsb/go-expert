package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Struct para a resposta da BrasilAPI
type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

// Struct para a resposta da ViaCEP
type ViaCEPResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	cep := "01153000"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	resultChan := make(chan string)
	errorChan := make(chan error)

	// Fazendo requisições simultâneas
	go fetchFromBrasilAPI(ctx, cep, resultChan, errorChan)
	go fetchFromViaCEP(ctx, cep, resultChan, errorChan)

	select {
	case result := <-resultChan:
		fmt.Println(result)
	case err := <-errorChan:
		fmt.Printf("Erro: %v\n", err)
	case <-ctx.Done():
		fmt.Println("Erro: Timeout de 1 segundo atingido.")
	}
}

func fetchFromBrasilAPI(ctx context.Context, cep string, resultChan chan<- string, errorChan chan<- error) {
	url := "https://brasilapi.com.br/api/cep/v1/" + cep
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		errorChan <- err
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errorChan <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorChan <- fmt.Errorf("BrasilAPI retornou status code %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorChan <- err
		return
	}

	var endereco BrasilAPIResponse
	if err := json.Unmarshal(body, &endereco); err != nil {
		errorChan <- err
		return
	}

	result := fmt.Sprintf("API: BrasilAPI\nCep: %s\nEstado: %s\nCidade: %s\nBairro: %s\nLogradouro: %s\nService: %s", endereco.Cep, endereco.State, endereco.City, endereco.Neighborhood, endereco.Street, endereco.Service)
	resultChan <- result
}

func fetchFromViaCEP(ctx context.Context, cep string, resultChan chan<- string, errorChan chan<- error) {
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		errorChan <- err
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errorChan <- err
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorChan <- fmt.Errorf("ViaCEP retornou status code %d", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorChan <- err
		return
	}

	var endereco ViaCEPResponse
	if err := json.Unmarshal(body, &endereco); err != nil {
		errorChan <- err
		return
	}

	result := fmt.Sprintf("API: ViaCEP\nCep: %s\nLogradouro: %s\nComplemento: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nUnidade: %s\nIBGE: %s\nGIA: %s\nDDD: %s\nSIAFI: %s", endereco.Cep, endereco.Logradouro, endereco.Complemento, endereco.Bairro, endereco.Localidade, endereco.Uf, endereco.Unidade, endereco.Ibge, endereco.Gia, endereco.Ddd, endereco.Siafi)
	resultChan <- result
}
