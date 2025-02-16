package handler

import (
	"bytes"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"               // Адрес сервиса
const itemName = "cup"                                // Название товара
const itemNameNotFound = "nonexistent_item"           // Несуществующий товар
const userID = "000d41d4-dd00-4ed4-8d14-5483e3147b31" // Пример UUID пользователя

// Тест покупки товара зарегистрированным пользователем
func TestBuyMerch(t *testing.T) {
	url := baseURL + "/api/buy/" + itemName

	headers := map[string]string{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3OTA0MDEsInVzZXJfaWQiOiIwMDBkNDFkNC1kZDAwLTRlZDQtOGQxNC01NDgzZTMxNDdiMzEiLCJ1c2VybmFtZSI6Im9sZWcifQ.lqRB9Qz6e0FrQtNmdEMLK_FvmDoVLIW1XAeLgfujGOg",
		"Content-Type":  "application/json",
		"userID":        userID,
	}

	resp := sendRequest(t, "GET", url, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", resp.StatusCode)
	}
}

// Тест попытки покупки товара незарегистрированным пользователем (без токена)
func TestBuyMerchUnauthorized(t *testing.T) {
	url := baseURL + "/api/buy/" + itemName

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := sendRequest(t, "GET", url, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401 (Unauthorized), но получен %d", resp.StatusCode)
	}
}

// Тест покупки товара, которого нет в наличии
func TestBuyMerchNotFound(t *testing.T) {
	url := baseURL + "/api/buy/" + itemNameNotFound

	headers := map[string]string{
		"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3OTA0MDEsInVzZXJfaWQiOiIwMDBkNDFkNC1kZDAwLTRlZDQtOGQxNC01NDgzZTMxNDdiMzEiLCJ1c2VybmFtZSI6Im9sZWcifQ.lqRB9Qz6e0FrQtNmdEMLK_FvmDoVLIW1XAeLgfujGOg",
		"Content-Type":  "application/json",
		"userID":        userID,
	}

	resp := sendRequest(t, "GET", url, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Ожидался статус 500, но получен %d", resp.StatusCode)
	}
}

func sendRequest(t *testing.T, method, url string, headers map[string]string) *http.Response {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(nil))
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Ошибка при выполнении запроса: %v", err)
	}

	return resp
}
