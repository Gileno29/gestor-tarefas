package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gileno29/gestor-tarefas/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error:Register not found %v", err)
	}

	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(todos)
}
