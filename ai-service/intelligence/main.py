from fastapi import FastAPI, HTTPException, status
from pydantic import BaseModel
from typing import List, Optional

app = FastAPI(
    title="Paisa AI Intelligence Service",
    description="LLM categorisation, narrative generation, and conversational query grounded in user transactions",
    version="1.0"
)

class CategoriseRequest(BaseModel):
    merchant_name: str
    vpa: Optional[str] = None
    amount: float

class CategoriseResponse(BaseModel):
    category_id: str
    subcategory_id: str
    confidence: float

class QueryRequest(BaseModel):
    user_id: str
    question: str
    context_data: dict  # GROUNDED DATA: last 90 days tx, budgets, goals

class QueryResponse(BaseModel):
    answer: str
    sources: List[str]

@app.get("/health")
def health_check():
    return {"status": "healthy", "service": "intelligence"}

@app.post("/categorise", response_model=CategoriseResponse)
def categorise_transaction(req: CategoriseRequest):
    # Blueprint section 8.3: Categorisation confidence scoring
    return CategoriseResponse(
        category_id="cat_groceries",
        subcategory_id="sub_vegetables",
        confidence=0.96
    )

@app.post("/query", response_model=QueryResponse)
def conversational_query(req: QueryRequest):
    # Blueprint section 8.2: Grounding rules (answer only using context_data)
    # Mock response
    return QueryResponse(
        answer="Based on your transaction data, you spent ₹1,200 on Groceries and ₹950 on Dining Out in the last 30 days. You are currently on track to stay within your ₹5,000 Groceries budget.",
        sources=["transactions_last_30_days", "budget_limits"]
    )
