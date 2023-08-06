# ⚙️ An example REST API built on a Fiber.
REST API built on a Fiber framework with Swagger documentation, a PostgreSQL database, JWT authentication and contenrezation in Docker compose.
![](https://github.com/c0caina/other/blob/main/Frame%201.jpg?raw=true)

## How to deploy
1) Install [Docker](https://docs.docker.com/engine/install).
2) Launch `docker-compose.yaml` at the root of the project:
```
docker compose up
```
3) Unfortunately, I have not yet implemented automatic deployment of tables in the database, so I will need to manually initialize the [migration](https://github.com/golang-migrate/migrate) or send an SQL query to the [psql](https://www.postgresql.org/docs/current/app-psql.html) utility.
4) Visit the swagger documentation page [http://127.0.0.1:3000/swagger/api/v1/index.html](http://127.0.0.1:3000/swagger/api/v1/index.html).
![](https://github.com/c0caina/other/blob/main/image.png?raw=true)
