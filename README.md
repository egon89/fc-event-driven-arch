# FC Event-Driven Architecture
This is a sample project to demonstrate an event-driven architecture using Go, PostgreSQL, and Kafka. The project includes a simple API to manage user balances and uses Kafka to handle balance updates asynchronously.

The core service emits events to Kafka when a transaction occurs, and the balance service listens to these events to update the user balances in the database. The original core service code is located in [https://github.com/devfullcycle/fc-eda](https://github.com/devfullcycle/fc-eda).

## Start
Create a `.env` file according to the `.env.example` file.
```bash
cp .env.example .env
```

To start the project, you can use Docker Compose. Run the following command:

```bash
docker-compose up
```

All initial database scripts and migrations will run automatically when you start the project with Docker Compose.

The core service will have two clients with two accounts each:
- Client John Doe:
  - Account ID: 546fbcb8-180a-4dd9-b36b-16304cf3e60a
  - Initial Balance: 1000.00

- Client Jane Doe:
  - Account ID: ca88c60a-6092-49a7-9d58-6bd16fbc30d2
  - Initial Balance: 500.00

Use the `request.http` file to do some transactions between these accounts.

---

## Dev information
### golang-migrate
- [Installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#linux-deb-package)
- [How to use migrations with Golang](https://medium.com/@albertcolom/how-to-use-migrations-with-golang-f46f4737beda)

```bash
# up
make migrate-up

# down
make migrate-down
```

Using golang-migrate docker image:
```bash
docker run --rm -v $(pwd)/database/migrations:/migrations --network host migrate/migrate \
    -path=/migrations -database="postgres://postgres:postgres@localhost:5432/event_driven_arch?sslmode=disable" up

```

### Accessing the database - PostgreSQL
You can access the PostgreSQL database using the following command:
```bash
docker exec -it postgresdb psql -U postgres -d event_driven_arch

# step by step version
docker-compose exec postgresdb /bin/bash

psql -U postgres # connect to the database

\l # list databases

\c event_driven_arch # connect to the database

\dt # list tables
```

### Accessing the database - MySQL
You can access the MySQL database using the following command:
```bash
docker-compose exec mysql bash

# Inside the MySQL container
mysql -u root -p

# Show databases
SHOW DATABASES;

# Use the wallet database
USE wallet;

# Show tables
SHOW TABLES;
```

### Producing messages
Use the Kafka console producer to send messages to the `balance` topic:
```bash
docker exec -it kafka bash

# Inside the Kafka container
kafka-console-producer --broker-list localhost:9092 --topic balance

# Produce a message
{
	"account_id_from": "48777f20-e467-4ed5-b379-7475378067fb",
	"account_id_to": "18fa87a6-4293-4acf-a003-1122efa0acc3",
	"balance_account_id_from": 100.21,
	"balance_account_id_to": 50.90
}
```
Or you can use the Control Center to produce messages to the `balance` topic using the UI.
