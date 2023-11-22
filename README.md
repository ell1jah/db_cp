# Курсова работа по базам данных ИУ7 МГТУ им. Н. Э. Баумана

## Тема: Разработка базы данных для хранения и обработки данных магазина одежды

## Суть работы
REST API на Go для магазина одежды с использованием Postgres. Исследование посвящено нагрузочному тестированию сервера с помощью [Locust](https://locust.io/).

## Запуск

### Для быстрого запуска выполните следующие действия:
1. make buildPostgres
2. make runPostgres
3. make fillPostgres
4. make run

### Для запуска тестирования:
1. make test
2. В браузере открыть http://localhost:8089/

## Методы API:

### Авторизация 

- GET /
- POST /register
- POST /login

### Бренд

- GET /brand/{BRAND_ID}
- PUT /brand
- POST /brand/{BRAND_ID}
- DELETE /brand/{BRAND_ID}

### Одежда

- GET /item/{ITEM_ID}
- PUT /item
- POST /item/{ITEM_ID}
- DELETE /item/{ITEM_ID}
- GET /items
    - sex
    - category
    - brand
    - order

### Корзина

- GET /basket
- POST /basket - совершить заказ
- POST /basket/add/{ITEM_ID}
- POST /basket/dec/{ITEM_ID}

### Заказ

- GET /order/{ORDER_ID}
- POST /order/{ORDER_ID}
- GET /orders/my
- GET /orders
