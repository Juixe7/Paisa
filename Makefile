.PHONY: help dev-deps dev-deps-down run-backend run-ai-doc run-ai-intel install-mobile run-mobile

help:
	@echo "Paisa Monorepo Make Commands:"
	@echo "  dev-deps          - Start local Docker dependencies (Postgres, Redis, RabbitMQ, MinIO, MailHog)"
	@echo "  dev-deps-down     - Stop local Docker dependencies"
	@echo "  run-backend       - Start the Go backend monolith"
	@echo "  run-ai-doc        - Start Python AI Document Processing service (port 8001)"
	@echo "  run-ai-intel      - Start Python AI Intelligence service (port 8002)"
	@echo "  install-mobile    - Install dependencies for Expo mobile app"
	@echo "  run-mobile        - Start Expo mobile app development server"

dev-deps:
	docker compose up -d

dev-deps-down:
	docker compose down

run-backend:
	cd backend && go run cmd/server/main.go

run-ai-doc:
	cd ai-service && uvicorn document_processing.main:app --port 8001 --reload

run-ai-intel:
	cd ai-service && uvicorn intelligence.main:app --port 8002 --reload

install-mobile:
	cd mobile && npm install

run-mobile:
	cd mobile && npm run start
