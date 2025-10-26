build: 
	@cd cmd/api && go build -o ../../bin/movepilot-api

run:
	@./bin/movepilot-api

dev: 	
	@cd cmd/api && go build -o ../../bin/movepilot-api
	@./bin/movepilot-api

create-migration:
	@if [ -z "$(name)" ]; then \
		echo "❌ Please provide a migration name with: make create-migration name=your_name"; \
	else \
		goose create $(name) go; \
	fi

test:
	@go test -v ./... | grep -v '\[no test files\]' 
	

start-db: 
	@docker compose up -d

stop-db:
	@docker compose down 

exec-db:
	@docker exec -it movepilot psql -U postgres
