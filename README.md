# go-clean-architecture

## Description

### En

This is an example of implementation of Clean Architecture in Golang

Characteristics of Clean Architecture by Uncle Bob:

- _Independent of Frameworks_. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- _Testable_. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- _Independent of UI_. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- _Independent of Database_. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- _Independent of any external agency_. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

---

### Ru

Это пример реализации чистой архитектуры на Golang

Характеристики Чистой Архитектуры от дяди Боба:

- _Независимость от фреймворков_. Архитектура не зависит от наличия какой-либо библиотеки. Это позволяет рассматривать фреймворки как инструменты, вместо того чтобы стрататься втиснуть систему в их рамки.
- _Простота тестирования_. Бизнес-правила можно тестировать без пользовательского интерефейста, базы данных, веб-сервера и любых других внешних элементов.
- _Назависимость от пользовательского интерфейса_. Пользовательский интерфейс можно легко изменять, не затрагивая остальную систему. Напрмер, веб-интерфейс можно заменить консольным интерфейсом, не изменяя бизнес-правил.
- _Независимость от базы данных_. Вы можете поменять Oracle или SQLServer на Mongo, BigTable, CouchDB или что-то еще. Бизнес-правила не привязаны к базе данных.
- _Независимость от любых внешних агентов_. Ваши бизнес-правила ничего не знают об интерфейсах, ведущих во внешний мир.

Подробнее на https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

---

### Schemas

![CleanArchitecture](https://github.com/Koichi-hub/go-clean-architecture/raw/master/CleanArchitecture.jpg)

![AppScheme](https://github.com/Koichi-hub/go-clean-architecture/raw/master/AppScheme.jpg)

---

### Start

```bash
$ git clone https://github.com/Koichi-hub/go-clean-architecture.git

$ cd go-clean-architecture

$ cp .example.env .env

$ make docker-build

$ make docker-compose-up
```

### Dev

```bash
$ git clone https://github.com/Koichi-hub/go-clean-architecture.git

$ cd go-clean-architecture

$ cp .example.env .dev.env

# run tests
$ make tests

### run app locally ###

# run database, for example mysql in docker
$ docker run --rm --name mysql -p 3306:3306 --network my_net -d -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=todo_db mysql:latest

# run app
$ make dev
```
