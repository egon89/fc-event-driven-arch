# FC Event-Driven Architecture
This is a sample project to demonstrate an event-driven architecture using Go, PostgreSQL, and Kafka. The project includes a simple API to manage user balances and uses Kafka to handle balance updates asynchronously.

## Database
### golang-migrate
- [Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#linux-deb-package)
- [How to use migrations with Golang](https://medium.com/@albertcolom/how-to-use-migrations-with-golang-f46f4737beda)

```bash
# up
make migrate-up

# down
make migrate-down
```

### Accessing the database
You can access the PostgreSQL database using the following command:
```bash
docker exec -it postgres psql -U postgres -d event_driven_arch

# step by step version
docker-compose exec postgres /bin/bash

psql -U postgres # connect to the database

\l # list databases

\c event_driven_arch # connect to the database

\dt # list tables
```
