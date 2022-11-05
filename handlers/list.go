package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo-go/models"
)

func List(w http.ResponseWriter, r *http.Request){
	fmt.Println("Teste2")
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error ao obter registros %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}