package handler

import (
	"bytes"
	"net/http"
	"testing"
)

const (
	validCredentials   = `{"Username": "testuser", "Password": "password123"}`
	invalidCredentials = `{"Username": "testuser", "Password": "wrongpassword"}`
)

// Тест успешной аутентификации (200 OK)
func TestAuthenticate_Success(t *testing.T) {
	url := baseURL + "/auth"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := sendAuthRequest(t, "POST", url, headers, []byte(validCredentials))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 401 (неправильные учетные данные)
func TestAuthenticate_Unauthorized(t *testing.T) {
	url := baseURL + "/auth"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := sendAuthRequest(t, "POST", url, headers, []byte(invalidCredentials))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 400 (некорректный JSON)
func TestAuthenticate_BadRequest(t *testing.T) {
	url := baseURL + "/auth"

	invalidJSON := `{"Username": 123, "Password": true}` // Некорректные типы данных

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := sendAuthRequest(t, "POST", url, headers, []byte(invalidJSON))
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Ожидался статус 400, но получен %d", resp.StatusCode)
	}
}

// Функция для отправки HTTP-запроса
func sendAuthRequest(t *testing.T, method, url string, headers map[string]string, body []byte) *http.Response {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
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
