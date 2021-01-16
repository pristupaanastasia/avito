# API Avito 

## Запуск
### Создание таблицы
sudo docker-compose run db bash
psql --host=db --username=postgres
pass: 1805
\connect avito_db

### Выполнение
sudo docker-compose up 

## Реализованные методы

### IndexHandler

В данном методе реализован вывод базы данных "Hotel", содержащий информацию о имеющихся номерах отеля. Вывод реализован в виде html страницы. Вызов метода : http://localhost:9000/

### IndexBookHandler

Метод реализует вывод строк базы данных "Booking", содержащий информацию о имеющихся бронирований номеров отеля. Вывод реализован в виде html страницы. Вызов метода : http://localhost:9000/booking

### CreateHandler

Метод реализует добавление нового бронирования в базу данных Booking. Для добавления нового бронирования нужно в теле url указать room_id(id номера отеля), date_start(начало брони),date_end(конец брони). Пример вызова: curl -X POST -d "room_id=2&date_start=2019-01-30&date_end=2022-01-02" http://localhost:9000/bookings/create
Ответ от API: {"booking_id":22}

### CreateRoomHandler
Метод реализует добавление нового номера отеля в базу данных Hotel. Для добавления нового номера нужно в теле url указать price(цена за сутки), description(описание). Пример вызова: curl -X POST -d "price=3000&description=the good place" http://localhost:9000/room/create
Ответ от API: {"room_id":13}

### FindHandler
Метод реализует поиск по room_id(id номера отеля) и вывод в формате json строк базы данных Booking. Вывод отсортирован по дате начала брони. Пример вызова:  curl -X GET "http://localhost:9000/bookings/list?room_id=24"
Ответ от API: [{"room_id":2,"booking_id":22,"date_start":"2019-01-30","date_end":"2022-01-02"},{"room_id":2,"booking_id":20,"date_start":"2019-12-30","date_end":"2022-01-02"}]

### RoomHandler
Метод реализует вывод в формате json строк базы данных Hotel. Вывод отсортирован по дате добавления номера отеля.Также можно отсортировать по цене и по дате добавления по убыванию.
Примеры вызова:
curl -X GET "http://localhost:9000/room/list"

Ответ API(отсортирован по дате добавления номера по возрастанию): 
[{"id":1,"price":10000,"description":"the good place","update":"2020-01-30"},{"id":2,"price":1200,"description":"\"the better place\"","update":"2021-01-12"},{"id":3,"price":1300,"description":"the room","update":"2021-01-12"},{"id":4,"price":1900,"description":"5 stars","update":"2021-01-12"},{"id":5,"price":1600,"description":"4 stars","update":"2021-01-12"},{"id":6,"price":1300,"description":"3 stars","update":"2021-01-12"},{"id":7,"price":1000,"description":"2 stars","update":"2021-01-12"},{"id":8,"price":1000,"description":"2 stars","update":"2021-01-12"},{"id":9,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":10,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":11,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":12,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":13,"price":3000,"description":"the good place","update":"2021-01-15"}]

curl -X GET "http://localhost:9000/room/list?sort=price"

Ответ API(отсортирован по цене):
[{"id":7,"price":1000,"description":"2 stars","update":"2021-01-12"},{"id":8,"price":1000,"description":"2 stars","update":"2021-01-12"},{"id":12,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":9,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":10,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":11,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":2,"price":1200,"description":"\"the better place\"","update":"2021-01-12"},{"id":3,"price":1300,"description":"the room","update":"2021-01-12"},{"id":6,"price":1300,"description":"3 stars","update":"2021-01-12"},{"id":5,"price":1600,"description":"4 stars","update":"2021-01-12"},{"id":4,"price":1900,"description":"5 stars","update":"2021-01-12"},{"id":13,"price":3000,"description":"the good place","update":"2021-01-15"},{"id":1,"price":10000,"description":"the good place","update":"2020-01-30"}]

curl -X GET "http://localhost:9000/room/list?sort=date_desc"

Ответ API(отсортирован по дате добавления номера по убыванию): 
[{"id":13,"price":3000,"description":"the good place","update":"2021-01-15"},{"id":2,"price":1200,"description":"\"the better place\"","update":"2021-01-12"},{"id":3,"price":1300,"description":"the room","update":"2021-01-12"},{"id":4,"price":1900,"description":"5 stars","update":"2021-01-12"},{"id":5,"price":1600,"description":"4 stars","update":"2021-01-12"},{"id":6,"price":1300,"description":"3 stars","update":"2021-01-12"},{"id":8,"price":1000,"description":"2 stars","update":"2021-01-12"},{"id":9,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":10,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":11,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":12,"price":1040,"description":"2 stars","update":"2021-01-12"},{"id":7,"price":1000,"description":"2 stars","update":"2021-01-12"},{"id":1,"price":10000,"description":"the good place","update":"2020-01-30"}]

### DeleteHandler
Метод реализует удаление брони по номеру id брони.
Пример запроса: curl -X GET "http://localhost:9000/bookings/delete?booking_id=8"

### DeleteRoomHandler
Метод реализует удаление номера отеля и всех его бронирований.
Пример запроса: curl -X GET "http://localhost:9000/room/delete?room_id=8"

### ErrorHandler
Метод реализует вывод ошибок в формате json.
Примеры: curl -X POST -d "price=t&description=the good place" http://localhost:9000/room/create
Ответ: {"Severity":"ОШИБКА","Code":"22P02","Message":"неверный синтаксис для типа integer: \"t\"","Detail":"","Hint":"","Position":"","InternalPosition":"","InternalQuery":"","Where":"","Schema":"","Table":"","Column":"","DataTypeName":"","Constraint":"","File":"numutils.c","Line":"256","Routine":"pg_strtoint32"}

 curl -X POST -d "room_id=2&date_start=е2019-01-30&date_end=2022-01-02" http://localhost:9000/bookings/create

{"Severity":"ОШИБКА","Code":"22007","Message":"неверный синтаксис для типа date: \"е2019-01-30\"","Detail":"","Hint":"","Position":"","InternalPosition":"","InternalQuery":"","Where":"","Schema":"","Table":"","Column":"","DataTypeName":"","Constraint":"","File":"datetime.c","Line":"3758","Routine":"DateTimeParseError"}
