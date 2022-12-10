# go-backend-template

## third-party libraries

### chi-router
go get -u github.com/go-chi/chi/v5

### pgx
Potgres driver for Go Lang

go get github.com/jackc/pgx/v4
go get github.com/jackc/pgconn (just to be safe)


### sqlboiler

```bash
https://github.com/volatiletech/sqlboiler
```


## development

### start auxiliary services via docker-compose

- PosgresDb service'


```bash
docker-compose up -d
```


### generate database access

```bash
sqlboiler psql
```