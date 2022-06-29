# Golang multi-tenant API example

Example project featuring a multi-tenant API in go language.

## Software and libraries used

### [gin](https://github.com/gin-gonic/gin) - Web Framework

```
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-gonic/gin
```

### [gorm](https://gorm.io/index.html) - GORM
And postgresql driver
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

### [Docker](https://www.docker.com/) - Need no presentation ;)

### [PostgreSQL](https://www.postgresql.org/) - Database

## Prepare master database

### Create main database

```sql
CREATE DATABASE "go-tenancy";
```

On the "go-tenancy" database:

```sql
CREATE TABLE "tenant"
(
    "id"   serial       NOT NULL,
    "code" varchar(20)  NOT NULL,
    "name" varchar(255) NOT NULL,
    "data" jsonb        NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "tenant_code_key" UNIQUE ("code")
);

INSERT INTO tenant ("id", "code", "name", "data")
VALUES (DEFAULT, 'tenant-1', 'First Tenant',  '{"database" : "go-tenant1", "host" : "localhost", "port" : "5432", "user" : "postgres", "password": "postgres"}'),
       (DEFAULT, 'tenant-2', 'Second Tenant', '{"database" : "go-tenant2", "host" : "localhost", "port" : "5432", "user" : "postgres", "password": "postgres"}');
```

### Create tenants databases

```sql
CREATE DATABASE "go-tenant1"; 
CREATE DATABASE "go-tenant2";

-- run this on each tenant:
CREATE TABLE "product"
(
    "id"          serial         NOT NULL,
    "name"        varchar(100)   NOT NULL,
    "description" varchar(255),
    "price"       decimal(10, 2) NOT NULL,
    PRIMARY KEY ("id")
);
```

Data on first tenant ("go-tenant1"):
```sql
INSERT INTO product ("id", "name", description, price)
VALUES (DEFAULT, 'Keyboard', 'Intelligent', 488.98),
       (DEFAULT, 'Computer', 'Ergonomic', 963.99),
       (DEFAULT, 'Mouse', 'Sleek', 134.50);
```
Data on second tenant ("go-tenant2"):
```sql
INSERT INTO product ("id", "name", description, price)
VALUES (DEFAULT, 'Pizza', 'Gorgeous', 24.99),
       (DEFAULT, 'Salad', 'Rustic', 13.99),
       (DEFAULT, 'Sauce', 'Unbranded', 4.50);
```