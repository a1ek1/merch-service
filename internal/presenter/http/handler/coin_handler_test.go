package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8080"

const (
	validToken   = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk3OTA2MTAsInVzZXJfaWQiOiJlMTUwZWEwMC1hZDkxLTQ0NTktODUyZS0wNmYwNWQ2MGM3N2IiLCJ1c2VybmFtZSI6InZsYWQifQ.78a7fRsG-AizavwTQmG-IBK587qxOqPNu5ErHAPbNMY"
	invalidToken = "Bearer invalid_token"
)

const (
	validReceiver   = "oleg"
	invalidReceiver = "" // Некорректное поле для теста 400
)

// Тест успешной отправки монет (200 OK)
func TestSendCoins_Success(t *testing.T) {
	url := baseURL + "/api/sendCoin"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"toUsername": validReceiver,
		"amount":     100,
	})

	headers := map[string]string{
		"Authorization": validToken,
		"Content-Type":  "application/json",
	}

	resp := sendCoinRequest(t, "POST", url, headers, requestBody)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 400 (невалидный JSON)
func TestSendCoins_BadRequest(t *testing.T) {
	url := baseURL + "/api/sendCoin"

	invalidRequestBody, _ := json.Marshal(map[string]interface{}{
		"toUsername": invalidReceiver, // Пустое имя получателя
		"amount":     "wrong_type",    // Некорректный тип данных
	})

	headers := map[string]string{
		"Authorization": validToken,
		"Content-Type":  "application/json",
	}

	resp := sendCoinRequest(t, "POST", url, headers, invalidRequestBody)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Ожидался статус 400, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 401 (отсутствует токен)
func TestSendCoins_Unauthorized(t *testing.T) {
	url := baseURL + "/api/sendCoin"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"toUsername": validReceiver,
		"amount":     100,
	})

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp := sendCoinRequest(t, "POST", url, headers, requestBody)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 500 (внутренняя ошибка сервера)
func TestSendCoins_InternalServerError(t *testing.T) {
	url := baseURL + "/api/sendCoin"

	requestBody, _ := json.Marshal(map[string]interface{}{
		"toUsername": "oleg",
		"amount":     -999999, // Некорректное значение, вызывающее ошибку сервера
	})

	headers := map[string]string{
		"Authorization": validToken,
		"Content-Type":  "application/json",
	}

	resp := sendCoinRequest(t, "POST", url, headers, requestBody)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Ожидался статус 500, но получен %d", resp.StatusCode)
	}
}

// Функция для отправки HTTP-запроса
func sendCoinRequest(t *testing.T, method, url string, headers map[string]string, body []byte) *http.Response {
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
