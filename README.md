## GODO

A simple todo app with Golang, PostgreSQL & Docker

### Getting Started

#### Executing the program

- Copy variables from `.env.example` to `.env`.

- How to run the program

With Docker compose

```
docker-compose up
```

Or with running Postgresql

```
go run ./cmd
```

Or run Postgresql with docker compose

`Postgre will init the todos table automatically`

```
docker-compose up db
go run ./cmd
```

### Tests

For running test run this command while db is up and running

```
go test -v ./test
```

### License

This project is [MIT licensed](http://opensource.org/licenses/MIT).
