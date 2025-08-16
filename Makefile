include .env

migrate-up: ## Migration up
	@migrate -path=./database/migrations -database "${DATABASE_URL}" -verbose up

migrate-down: ## Migration down
	@migrate -path=./database/migrations -database "${DATABASE_URL}" -verbose down
