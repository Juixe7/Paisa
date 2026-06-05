FROM python:3.11-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Run as non-root user (Blueprint section 11.3)
RUN useradd -u 1000 appuser && chown -R appuser /app
USER appuser

# Defaults to Document Processing service. Override MODULE and PORT for Intelligence service.
ENV PORT=8001
ENV MODULE=document_processing.main:app

EXPOSE 8001
EXPOSE 8002

CMD uvicorn $MODULE --host 0.0.0.0 --port $PORT
