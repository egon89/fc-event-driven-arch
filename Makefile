include .env

migrate-up: ## Migration up
	@migrate -path=./balance-service/database/migrations -database "${BALANCE_DATABASE_URL}" -verbose up

migrate-down: ## Migration down
	@migrate -path=./balance-service/database/migrations -database "${BALANCE_DATABASE_URL}" -verbose down
