from fastapi import FastAPI, HTTPException, status
from pydantic import BaseModel
from typing import Optional

app = FastAPI(
    title="Paisa AI Document Processing Service",
    description="SMS fallback parsing, text PDF bank statement extraction, and OCR services",
    version="1.0"
)

class SmsParseRequest(BaseModel):
    raw_sms: str

class SmsParseResponse(BaseModel):
    amount: float
    direction: str  # debit / credit
    merchant_name: str
    vpa: Optional[str] = None
    confidence: float

@app.get("/health")
def health_check():
    return {"status": "healthy", "service": "document-processing"}

@app.post("/sms/parse", response_model=SmsParseResponse)
def parse_sms(req: SmsParseRequest):
    # Blueprint section 8.1: SMS fallback parsing
    # Scaffolding mock logic:
    return SmsParseResponse(
        amount=150.00,
        direction="debit",
        merchant_name="Swiggy",
        vpa="swiggy@upi",
        confidence=0.88
    )

@app.post("/pdf/extract")
def extract_pdf(s3_key: str):
    # Blueprint section 8.4: S3 pre-signed URL + pdfplumber
    return {
        "status": "processing",
        "s3_key": s3_key,
        "message": "PDF extraction scheduled"
    }
