package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Gileno29/gestor-tarefas/handlers"
	"github.com/Gileno29/gestor-tarefas/models"
)

func TestGetByID(t *testing.T) {
	todo := models.Todo{
		ID:          '1',
		Title:       "Test",
		Description: "todo for test",
		Done:        true,
	}

	newRecord := httptest.NewRecorder()
	//method, URL, Body Content
	newRequest := httptest.NewRequest(http.MethodGet, "http://localhost/", nil)
	newRequest.URL.RawQuery = "id=1"
	fmt.Print(newRequest)
	handlers.Get(newRecord, newRequest)

	if http.StatusOK != newRecord.Code {
		t.Errorf("Status for Get By ID Code Expected: %d, got %d", http.StatusOK, newRecord.Code)
	}

	var eb models.Todo
	err := json.NewDecoder(newRequest.Body).Decode(&eb)
	if err != nil {
		t.Errorf("Erro ao decodificar resposta JSON: %v", err)
	}

	if !reflect.DeepEqual(eb, todo) {
		t.Errorf("Resposta incorreta: %v", todo)

	}

}
