# Matching APP API

## Requirement

* Docker
* docker compose

## Dependency
* Golang 1.17
* Gin
* Gorm
* MySQL 8.0

## Usage

### setup
```bash
cd YOUR_ROOT_PATH
make setup
```

### run
```bash
cd YOUR_ROOT_PATH
make up
or
make up-d // No log display
```

### Into the container
```bash
cd YOUR_ROOT_PATH
make exec-api
```

```bash
cd YOUR_ROOT_PATH
make exec-db
```

### Generate API Doc
```bash
cd YOUR_ROOT_PATH
make generate-api-doc
```

go to <http://localhost:8080/swagger/index.html>

### phpMyAdmin
go to <http://localhost:8888/>