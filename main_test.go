package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверка статус кода
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Статус код не 200")
	// проверка, что тело ответа не пустое
	responseBody := responseRecorder.Body
	assert.NotEmpty(t, responseBody, "Тело ответа пустое")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=12&city=moscow", nil)

	// здесь нужно добавить необходимые проверки

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Статус код не 200")

	responseBody := responseRecorder.Body.String()

	result := strings.Split(responseBody, ",")
	assert.Len(t, result, totalCount, "Некорректная длина")

}

func TestMainHandlerWrongCityValue(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=perm", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	expectedErrorMessage := "wrong city value"

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Сервис не возвращает 400 код ответа")

	responseBody := responseRecorder.Body.String()
	assert.Equal(t, expectedErrorMessage, responseBody, "Неправильное сообщение")

}
