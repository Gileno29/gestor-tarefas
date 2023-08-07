package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gileno29/gestor-tarefas/models"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Eror when its try decode json: %v", err)
		/*this method receives 3 parameters ResponseWriter, the error and the status code*/
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Eror update: %v", err)
		/*this method receives 3 parameters ResponseWriter, the error and the status code*/
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
