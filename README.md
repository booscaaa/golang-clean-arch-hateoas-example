<p align="center">
  <h1 align="center">Clean architecture and Level 3 of REST</h1>
  <p align="center">An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations.</p>
  <p align="center">
    <a href="https://github.com/booscaaa/golang-clean-arch-hateoas-example/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/golang-clean-arch-hateoas-example.svg?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/golang-clean-arch-hateoas-example/actions?query=workflow%3ADeploy"><img alt="Build status" src="https://img.shields.io/github/workflow/status/booscaaa/golang-clean-arch-hateoas-example/Test?style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/golang-clean-arch-hateoas-example"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/golang-clean-arch-hateoas-example/master.svg?style=for-the-badge"></a>
    <a href="https://todo-list-hateoas.herokuapp.com/swagger/index.html"><img alt="Swagger docs" src="https://img.shields.io/badge/api docs-swagger-purple?style=for-the-badge"></a>
  </p>
</p>

<br>


## Getting Started

### Prerequisites

To run this project in the development mode, you'll need to have a basic environment to run:

- A Golang SDK, that can be found [here](https://golang.org/).

<br>

### Installing

**Cloning the Repository**

```bash
git clone https://github.com/booscaaa/golang-clean-arch-hateoas-example

cd github.com/booscaaa/golang-clean-arch-hateoas-example
```

<br>

### Rename config.example.json to config.json

config.json - this is my heroku app

```json
{
  "database": {
    "url": ""
  },
  "hateoas": {
    "base": "https://todo-list-hateoas.herokuapp.com"
  }
}

```

<br>

# Running with docker-compose !!!!!!!
```bash
docker-compose up --build -d
```
<br>
<br>

# Running local

```bash
go mod tidy
go run main.go
```

# Testing

```bash
go mod tidy
go test -v ./... 
```

## To get test coverage
```bash
go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
```

## To get new mocks for testing
```bash
mockgen -source=core/domain/item.go -destination=core/domain/mocks/fake_item_repository.go -package=mocks
```

<br>
<br>
<br>



### URLs to show the aplications

- API = http://YOUR_MACHINE_IP:<CONFIG_JSON_PORT>

### To access the endpoints documentation

- https://todo-list-hateoas.herokuapp.com/swagger/index.html

<br>
<br>


## Api application built With

- [Golang](https://golang.org/)
- [Gorilla Handlers](https://github.com/gorilla/handlers)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Viper](https://github.com/spf13/viper)
- [PGX](https://github.com/jackc/pgx)
- [Swaggo](https://github.com/swaggo/http-swagger)
- [Testfy](https://github.com/stretchr/testify)


## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/booscaaa/golang-clean-arch-hateoas-example/blob/master/LICENSE) file for details
