# Clean archtecture and Level 3 of REST

## About this Project

The idea of the App is:

_"An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations."._

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

### Config our environment variables like this (this config set to my heroku application database)

.env.production and .env.development

```bash
DB_HOST=ec2-50-16-198-4.compute-1.amazonaws.com
DB_USER=izrykubjdwzynh
DB_PASSWORD=d6d5c87ab5b0d734323acc8dc729c3f389f4368c8dc73cbd9be844bce3173fb2
DB_NAME=d6r28h5h9fqrn3
SSL_MODE=
BASE_URL=http://localhost:3333
```

# Running !!!!!!!

```bash
$ GO_ENV=development go run main.go
```

or

```bash
$ export GO_ENV=development
$ go run main.go
```

# Testing !!!!!!!
```bash
go test ./...
```

<br>

### URLs to show the aplications

- API = http://YOUR_MACHINE_IP:3333

### To access the endpoints documentation

- https://todo-list-hateoas.herokuapp.com/doc/index.html

<br>
<br>
<br>

## Api application built With

- [Golang](https://golang.org/)
- [Govalidator](https://github.com/asaskevich/govalidator)
- [Gorilla Handlers](https://github.com/gorilla/handlers)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [Go dotenv](https://github.com/joho/godotenv)
- [Lib PQ](https://github.com/lib/pq)

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
