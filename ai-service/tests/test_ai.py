from fastapi.testclient import TestClient
from document_processing.main import app as doc_app
from intelligence.main import app as intel_app

doc_client = TestClient(doc_app)
intel_client = TestClient(intel_app)

def test_document_processing_health():
    response = doc_client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "healthy", "service": "document-processing"}

def test_intelligence_health():
    response = intel_client.get("/health")
    assert response.status_code == 200
    assert response.json() == {"status": "healthy", "service": "intelligence"}
