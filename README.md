# Эмулятор платежного сервиса

## Пакеты:
1) [github.com/caarlos0/env](https://github.com/caarlos0/env) - Минимальный пакет для работы с переменными окружения, использовал для конфигурирования сервиса
2) [github.com/gorilla/mux](https://github.com/gorilla/mux) - Роутер
3) [github.com/lib/pq](https://github.com/lib/pq) - Драйвер для postgres
4) [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) - Логгер
## Установка:
   1) Скачать текущий репозиторий
   2) Скачать докер [docker.com/get-started/](https://www.docker.com/get-started/)
   3) Собрать все сервисы описанные в docker-compose. `docker compose build`
   4) Запустить все сервисы описанные в  docker-compose. `docker compose up -d`
   5) Платежный сервис готов к работе
## API:
1) Создание платежа. Запрос: POST http://127.0.0.1:8000/create-transactions

Request body(data-raw):
```
{
    "user-id": uint,
    "email": string,
    "sum": float,
    "currency": string(ограничение переменной в бд 3 символа. Пример: usd, eur, rub)
}
```
Response body: nil
