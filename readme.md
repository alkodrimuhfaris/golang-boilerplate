## Migration Master

```BASH
migrate -path src/db/migration -database 'sqlserver://{username}:{password}@127.0.0.1:1433/sql-db?sslmode=disable&database={DB_NAME}' -verbose up
```

## Migration Rollback Master

```BASH
migrate -path src/db/migration -database 'sqlserver://{username}:{password}@127.0.0.1:1433/sql-db?sslmode=disable&database={DB_NAME}' -verbose down
```


## How to Build

### Pre-requisities

1. Docker

### Commands

```
$ cd <source-directory>

$ docker build --network=host -t gcr.io/papitupi/papitupi-web:1.0.0-UAT .

$ docker push gcr.io/papitupi/papitupi-web:1.0.0-UAT
```

## How to Run

### Docker Swarm

#### Pre-requisities

1. Docker Swarm
1. Built docker image

#### Commands

```
$ docker network create --attachable --driver overlay --scope swarm cluster

$ cd deploy

$ vi secrets.env

$ docker secret create papitupi-web-secrets ./secrets.env

$ STAGE=uat VERSION=1.0.0-UAT docker stack deploy --with-registry-auth --compose-file ./docker-stack.yml papitupi-web

```
