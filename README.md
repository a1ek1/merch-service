# Сервис по работе с транзакциями

HTTP сервер, реализующий REST API, выполненный в рамках тестового задания на стажировку для компании _"Avito"_.


## Подготовка к запуску сервера

---

### Необходимо, чтобы на вашем устройстве были установлены

- Docker
- Docker compose
- goose (для применения миграций)
- Go 1.23

## Запуск сервера

1. Склонируйте проект

```bash
    git clone git@github.com:a1ek1/merch-service.git
    cd merch-service
```

2. Запустите контейнер с базой данных

```bash
    # Команду выполняем из директории merch-service
    docker-compose up
```

После выполнения данной команды у вас должен запуститься контейнер с PostgreSQL, достуный на порту **5435**

3. Примените миграции

Если у вас не установлен goose, то нужно в командной строке выполнить команду
```bash
    go install github.com/pressly/goose/v3/cmd/goose@latest
```

После установки проверьте корректность работы
```bash
    goose --version
```
Дальше нужно _перейти в директорию merch-service/migrations_ и выполнить следующую команду

```bash
    # Указаны значения по умолчанию. Если поменяете в конфиге и в docker-compose.yml, то здесь будут другие данные
    goose postgres "host=localhost user=postgres port=5434 password=postgres database=merch_service sslmode=disable" up
```

Вам должно вывестить сообщение об успешном применении миграций

4. Запустите файл main.go

```bash
    # Из директории merch-service
    go run cmd/main.go
```
После выполнения этой команды сервер будет доступен на **_localhost:8080_**
## Тестирование работы

1. Регистрация/авторизация в сервисе. Вам возвращается JWT токен.

```bash
    curl -X POST http://localhost:8080/api/auth \
          -H "Content-Type: application/json" \
          -d "{
            \"username\":\"{логин}\",
            \"password\":\"{пароль}\"
          }"
```
2. Покупка предметов

```bash
    # Введите токен, полученный при регистрации/авторизации
    curl -X GET "http://localhost:8080/api/buy/{название_предмета}" \
        -H "Authorization: Bearer {токен}"
    
```

3. Отправка монет пользователю

```bash
    # Введите число нужных транзакций
    curl -X POST http://localhost:8080/api/sendCoin \
        -H "Authorization: Bearer {токен_отправителя}" \
        -H "Content-Type: application/json" \
        -d "{\"toUsername\": \"{логин_получателя}\", \"amount\": {сумма_перевода}}"
```

4. Получение информации о пользователе

```bash
    # Введите токен пользователя, информацию о котором вы хотите получить
    curl -X GET http://localhost:8080/api/info \
        -H "Authorization: Bearer {токен}"
```