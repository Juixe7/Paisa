# Paisa

Paisa is an AI-powered personal finance companion built specifically for Indian users, designed to run as a modular monolith.

## Repository Structure

The project is structured as a monorepo containing all services, mobile apps, and infrastructure:

- **`backend/`**: Go monolith implementing the primary application logic, database migrations, and HTTP APIs.
- **`ai-service/`**: Python FastAPI services containing AI algorithms (SMS fallback parsing, text extraction, LLM grounding).
- **`mobile/`**: Expo React Native mobile application.
- **`infrastructure/`**: IaC definitions for Terraform, Helm charts, Ansible, and raw K8s manifests.
- **`docs/`**: API specifications (OpenAPI 3.0) and architecture documents (ADRs).

## Getting Started

### Prerequisites

You will need the following tools installed:
- Go 1.22+
- Node.js 20 LTS + NPM
- Python 3.11+
- Docker Desktop
- Expo CLI

### Setup Local Environment

1. Clone the repository and navigate to the root directory.
2. Initialize the local environment file:
   ```bash
   cp .env.example .env
   ```
3. Start the infrastructure dependencies (PostgreSQL, Redis, RabbitMQ, MinIO, MailHog) via Docker Compose:
   ```bash
   docker compose up -d
   ```
4. Verify all infrastructure services are running:
   ```bash
   docker compose ps
   ```

### Running Backend (Go)

Navigate to the `backend/` directory:
1. Load dependencies:
   ```bash
   go mod download
   ```
2. Run database migrations (requires `golang-migrate` CLI or run via Makefile if configured):
   ```bash
   make migrate-up
   ```
3. Start the backend server:
   ```bash
   go run cmd/server/main.go
   ```

### Running AI Service (Python)

Navigate to the `ai-service/` directory:
1. Create a virtual environment:
   ```bash
   python -m venv .venv
   source .venv/bin/activate  # On Windows: .venv\Scripts\activate
   ```
2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```
3. Start Document Processing service (port 8001):
   ```bash
   uvicorn document_processing.main:app --port 8001 --reload
   ```
4. Start Intelligence service (port 8002):
   ```bash
   uvicorn intelligence.main:app --port 8002 --reload
   ```

### Running Mobile (Expo)

Navigate to the `mobile/` directory:
1. Install dependencies:
   ```bash
   npm install
   ```
2. Start the Expo development server:
   ```bash
   npm run start
   ```

## Development Commands

We use a root-level `Makefile` to simplify common development tasks. Run `make help` to see all available commands.
