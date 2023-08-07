package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gileno29/gestor-tarefas/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Eror when its try decode json: %v", err)
		/*this method receives 3 parameters ResponseWriter, the error and the status code*/
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Eror to delete: %v", err)
		/*this method receives 3 parameters ResponseWriter, the error and the status code*/
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if rows > 1 {
		log.Printf("Error: do much rows affect %v", rows)
	}
	resp := map[string]any{
		"Error":    false,
		"Message:": "delete sucessifull",
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
