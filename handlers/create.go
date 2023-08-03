package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gileno29/gestor-tarefas/models"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {

		log.Printf("Eror when its try decode json: %v", err)
		/*this method receives 3 parameters ResponseWriter, the error and the status code*/
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
		}
	}

}
