<p align="center">
  <h1 align="center">Clean architecture and Level 3 of REST</h1>
  <p align="center">An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations.</p>
  <p align="center">
    <a href="https://github.com/booscaaa/golang-clean-arch-hateoas-example/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/golang-clean-arch-hateoas-example.svg?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/golang-clean-arch-hateoas-example/actions?query=workflow%3ADeploy"><img alt="Build status" src="https://img.shields.io/github/workflow/status/booscaaa/golang-clean-arch-hateoas-example/Deploy?style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/golang-clean-arch-hateoas-example"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/golang-clean-arch-hateoas-example/master.svg?style=for-the-badge"></a>
    <a href="https://todo-list-hateoas.herokuapp.com/doc/index.html"><img alt="Swagger docs" src="https://img.shields.io/badge/api docs-swagger-purple?style=for-the-badge"></a>
  </p>
</p>

<br>

## Why?

This project is part of my personal portfolio, so, I'll be happy if you could provide me any feedback about the project, code, structure or anything that you can report that could make me a better developer!

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/).

<br>

## Functionalities

- Include tasks
- Delete tasks
- Change tasks
- Search tasks
- Search tasks by acronym

<br>

## Getting Started

### Prerequisites

To run this project in the development mode, you'll need to have a basic environment to run:

- A Golang SDK, that can be found [here](https://golang.org/).

<br>

### Installing

**Cloning the Repository**

```bash
$ git clone https://github.com/booscaaa/golang-clean-arch-hateoas-example

$ cd golang-clean-arch-hateoas-example
```

<br>

### Rename config.example.json to config.json

config.json - this is my heroku app

```json
{
  "database": {
    "url": "postgres://izrykubjdwzynh:d6d5c87ab5b0d734323acc8dc729c3f389f4368c8dc73cbd9be844bce3173fb2@ec2-50-16-198-4.compute-1.amazonaws.com:5432/d6r28h5h9fqrn3"
  },
  "hateoas": {
    "base": "https://todo-list-hateoas.herokuapp.com"
  }
}

```

<br>

# Running with docker-compose !!!!!!!
```bash
$ docker-compose up --build -d
```
<br>
<br>

# Running local

```bash
go run main.go
```

# Testing

```bash
$ go test -v ./... 
```

## To get test coverage
```bash
$ go test -v -coverprofile cover.out ./...
```

<br>
<br>
<br>



### URLs to show the aplications

- API = http://YOUR_MACHINE_IP:<CONFIG_JSON_PORT>

### To access the endpoints documentation

- https://todo-list-hateoas.herokuapp.com/doc/index.html

<br>
<br>


## Api application built With

- [Golang](https://golang.org/)
- [Govalidator](https://github.com/asaskevich/govalidator)
- [Gorilla Handlers](https://github.com/gorilla/handlers)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Go dotenv](https://github.com/joho/godotenv)
- [Lib PQ](https://github.com/lib/pq)
- [GO sqlmock](https://github.com/DATA-DOG/go-sqlmock)
- [Mockery](https://github.com/vektra/mockery/v2)
- [Swaggo](https://github.com/swaggo/http-swagger)
- [Testfy](https://github.com/stretchr/testify)

<br>
<br>
<br>

## Contributing

You can send how many PR's do you want, I'll be glad to analyze and accept them! And if you have any question about the project...

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/)

Thank you!

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/booscaaa/golang-clean-arch-hateoas-example/blob/master/LICENSE) file for details
