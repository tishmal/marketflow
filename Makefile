up:
	docker compose -f deployments/docker-compose.yml up --build

down:
	docker compose -f deployments/docker-compose.yml down -v
