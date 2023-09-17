package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gileno29/gestor-tarefas/handlers"
)

func Testlist(t *testing.T) {

	newRecord := httptest.NewRecorder()
	//method, URL, Body Content
	newRequest := httptest.NewRequest(http.MethodGet, "/", nil)

	handlers.List(newRecord, newRequest)

	if http.StatusOK != newRecord.Code {
		t.Errorf("Status Code Expected: %d, got %d", http.StatusOK, newRecord.Code)
	}

}
