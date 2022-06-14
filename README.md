# Эмулятор платежного сервиса

## Пакеты:
1) [github.com/caarlos0/env](https://github.com/caarlos0/env) - Минимальный пакет для работы с переменными окружения, использовал для конфигурирования сервиса
2) [github.com/gorilla/mux](https://github.com/gorilla/mux) - Роутер
3) [github.com/lib/pq](https://github.com/lib/pq) - Драйвер для postgres
4) [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Логгер
5) [golang.org/x/crypto/bcrypt](https://golang.org/x/crypto/bcrypt) - Хеширование, использовал для авторизации
6) [github.com/asaskevich/govalidator](https://github.com/asaskevich/govalidator) - Валидация
## Установка:
   1) Скачать текущий репозиторий
   2) Скачать docker [docker.com/get-started/](https://www.docker.com/get-started/)
   3) Собрать все сервисы описанные в docker-compose. `docker compose build`
   4) Запустить все сервисы описанные в  docker-compose. `docker compose up -d`
   5) Выполнить миграцию таблицы в БД. Путь к файлу с таблицей: `migratinos/up.sql`
   6) Платежный сервис готов к работе
## API:
___
1) Создание платежа. POST https://127.0.0.1:443/create-payment

Request body(data-raw):
````json
{
    "user-id": 1,
    "email": "1@email.com",
    "sum": 100,
    "currency": "usd"
}
````
Response body(JSON):
````json
"Payment created"
````
___
2) Изменение статуса платежа.  PATCH https://127.0.0.1:443/change-payment-status

Authorization(Basic Auth): 
````
Username: service
Password: service
````
Request body(data-raw):
````json
{
    "transact-id": 1,
    "status": "УСПЕХ"
}
````
Response body(JSON):
````json
"Payment status changed successfully"
````
___
3) Проверка статуса платежа по ID. GET https://127.0.0.1:443/get-payment-status-by-id/{id}

Request body: none

Response body(JSON):
````json
"УСПЕХ"
````
___
4) Получение списка всех платежей пользователя по его ID. GET https://127.0.0.1:443/get-payments-by-userid/{id}

Request body: none

Response body(JSON):
````json
[
    {
        "user-id": 1,
        "transact-id": 1,
        "email": "1@email.com",
        "sum": 100,
        "currency": "usd",
        "date-time-create": "2022-06-13T19:35:56.974408Z",
        "date-time-last-change": "2022-06-13T19:40:27.529686Z",
        "status": "УСПЕХ"
    }
]
````
___
5) Получение списка всех платежей пользователя по его e-mail. GET https://127.0.0.1:443/get-payments-by-email/{email}

Request body: none

Response body(JSON):
````json
[
  {
    "user-id": 1,
    "transact-id": 1,
    "email": "1@email.com",
    "sum": 100,
    "currency": "usd",
    "date-time-create": "2022-06-13T19:35:56.974408Z",
    "date-time-last-change": "2022-06-13T19:40:27.529686Z",
    "status": "УСПЕХ"
  }
]
````
___
6) Отмена платежа по его ID. PATCH https://127.0.0.1:443/cancel-payment-by-id/{id}

Request body: none

Response body(JSON):
````json
"Payment canceled successfully"
````
___