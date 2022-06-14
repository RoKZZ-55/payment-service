# Эмулятор платежного сервиса

## Пакеты:
1) [github.com/caarlos0/env](https://github.com/caarlos0/env) - Минимальный пакет для работы с переменными окружения, использовал для конфигурирования сервиса
2) [github.com/gorilla/mux](https://github.com/gorilla/mux) - Роутер
3) [github.com/lib/pq](https://github.com/lib/pq) - Драйвер для postgres
4) [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Логгер
5) [golang.org/x/crypto/bcrypt](https://golang.org/x/crypto/bcrypt) - Хеширование, использовал для авторизации
## Установка:
   1) Скачать текущий репозиторий
   2) Скачать docker [docker.com/get-started/](https://www.docker.com/get-started/)
   3) Собрать все сервисы описанные в docker-compose. `docker compose build`
   4) Запустить все сервисы описанные в  docker-compose. `docker compose up -d`
   5) Выполнить миграцию таблицы в БД, файл `migratinos/up.sql`
   6) Платежный сервис готов к работе
## API:
___
1) Создание платежа. POST https://127.0.0.1:443/create-payment

Request body(data-raw):
````
{
    "user-id": uint,
    "email": string,
    "sum": float,
    "currency": string(максимальное кол-во символов 3. Пример: usd, eur, rub)
}
````
Response body(JSON):
````
"Payment created"
````
___
2) Изменение статуса платежа. POST https://127.0.0.1:443/change-payment-status

Authorization(Basic Auth): 
````
Username: service
Password: service
````
Request body(data-raw):
````
{
    "transact-id": uint,
    "status": string
}
````
Response body(JSON):
````
"Payment status changed successfully"
````
___
3) Проверка статуса платежа по ID. POST https://127.0.0.1:443/get-payment-status-by-id/{id}

Request body: none

Response body(JSON):
````
"УСПЕХ"
````
___
4) Получение списка всех платежей пользователя по его ID. POST https://127.0.0.1:443/get-payments-by-userid/{id}

Request body: none

Response body(JSON):
````
[
    {
        "user-id": uint,
        "transact-id": uint,
        "email": string,
        "sum": float,
        "currency": string(максимальное кол-во символов 3. Пример: usd, eur, rub),
        "date-time-create": string,
        "date-time-last-change": string,
        "status": string
    }
]
````
___
5) Получение списка всех платежей пользователя по его e-mail. POST https://127.0.0.1:443/get-payments-by-email/{email}

Request body: none

Response body(JSON):
````
[
    {
        "user-id": uint,
        "transact-id": uint,
        "email": string,
        "sum": float,
        "currency": string(максимальное кол-во символов 3. Пример: usd, eur, rub),
        "date-time-create": string,
        "date-time-last-change": string,
        "status": string
    }
]
````
___
6) Отмена платежа по его ID. POST https://127.0.0.1:443/cancel-payment-by-id/{id}

Request body: none

Response body(JSON):
````
"Payment canceled successfully"
````
___