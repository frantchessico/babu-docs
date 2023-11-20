package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/gorilla/mux"
)

// Estrutura de dados recebida na requisição
type Data struct {
	CollectionKey string            `json:"collectionKey"`
	Fields        map[string]string `json:"fields"`
}

// Valor padrão para CollectionKey
const defaultCollectionKey = "SUA_CHAVE_DE_COLECAO_AQUI"

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	url := "https://babuapi.savanapoint.com/api/core"
	secretKey := "SUA_CHAVE_SECRETA_AQUI"

	// Decodificando o JSON do corpo da requisição
	var requestData Data
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Definindo o valor padrão para CollectionKey se não for fornecido na requisição
	if requestData.CollectionKey == "" {
		requestData.CollectionKey = defaultCollectionKey
	}

	// Preparando a requisição para a API externa
	payload, err := json.Marshal(requestData)
	if err != nil {
		http.Error(w, "Erro ao criar payload JSON", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "Erro ao criar requisição", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("secret_key", secretKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Erro ao enviar requisição", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Lendo a resposta da API externa
	
	if err != nil {
		http.Error(w, "Erro ao ler resposta", http.StatusInternalServerError)
		return
	}

	// Enviando a resposta de volta ao cliente
	w.WriteHeader(resp.StatusCode)
	
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/send-data", handlePostRequest).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", nil)
}
