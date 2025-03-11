package handler

import (
	"net/http"
	"testing"
)

// Тест успешного получения информации о пользователе (200 OK)
func TestGetUserInfo_Success(t *testing.T) {
	url := baseURL + "/user/info"

	headers := map[string]string{
		"Authorization": validToken,
	}

	resp := sendInfoRequest(t, "GET", url, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 401 (неавторизованный доступ)
func TestGetUserInfo_Unauthorized(t *testing.T) {
	url := baseURL + "/user/info"

	headers := map[string]string{}

	resp := sendInfoRequest(t, "GET", url, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401, но получен %d", resp.StatusCode)
	}
}

// Тест ошибки 403 (невалидный токен)
func TestGetUserInfo_Forbidden(t *testing.T) {
	url := baseURL + "/user/info"

	headers := map[string]string{
		"Authorization": invalidToken,
	}

	resp := sendInfoRequest(t, "GET", url, headers)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Ожидался статус 403, но получен %d", resp.StatusCode)
	}
}

// Функция для отправки HTTP-запроса
func sendInfoRequest(t *testing.T, method, url string, headers map[string]string) *http.Response {
	req, err := http.NewRequest(method, url, nil)
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
