# Golang REST API con PostgreSQL

Configurar archivo docker compose 

```bash
version: '2'
services:
  postgres:
    image: 'postgres:latest'
    restart: always
    volumes:
      - './postgres_data:/var/lib/postgresql/data'
    environment:
      - POSTGRES_PASSWORD=secure_pass_here
    ports:
      - '5432:5432'
```


## Dockerize Postgres

Ejecutar ...

```bash
docker-compose up -d
docker ps
```

Conectar PostgreSQL con DBeaver

```mysql

Seleccionar nueva conexión para PostgreSQL
Server host: localhost
Port: 5432
User name: postgres
Password: secure_pass_here
        

```

Error: No se puede crear ...

```mysql
validar la conexion con la base de datos docker        
```