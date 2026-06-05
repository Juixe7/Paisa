-- Define Enum Types (Blueprint section 3.3)
CREATE TYPE transaction_direction AS ENUM ('debit', 'credit');
CREATE TYPE transaction_source AS ENUM ('sms', 'manual', 'pdf', 'intent');
CREATE TYPE transaction_status AS ENUM ('uncategorised', 'confirmed', 'flagged', 'excluded');
CREATE TYPE transaction_inflow_type AS ENUM ('salary', 'refund', 'friend_payback', 'cashback');
CREATE TYPE merchant_cache_source AS ENUM ('user_correction', 'ai_learned', 'rule');

-- Create Categories & Subcategories tables (Categorisation package)
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('expense', 'income')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS subcategories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Transactions Table (Blueprint section 3.4)
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount NUMERIC(12, 2) NOT NULL,
    direction transaction_direction NOT NULL,
    merchant_name TEXT NOT NULL,
    vpa TEXT,
    raw_sms TEXT,
    source transaction_source NOT NULL,
    category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    subcategory_id UUID REFERENCES subcategories(id) ON DELETE SET NULL,
    confidence_score DOUBLE PRECISION,
    status transaction_status NOT NULL DEFAULT 'uncategorised',
    inflow_type transaction_inflow_type,
    transacted_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Pending Intents Table (Transaction package)
CREATE TABLE IF NOT EXISTS pending_intents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount NUMERIC(12, 2) NOT NULL,
    merchant_name TEXT NOT NULL,
    category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Merchant Cache Table (Blueprint section 3.4)
CREATE TABLE IF NOT EXISTS merchant_cache (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    merchant_key TEXT NOT NULL,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    subcategory_id UUID REFERENCES subcategories(id) ON DELETE SET NULL,
    source merchant_cache_source NOT NULL DEFAULT 'ai_learned',
    confidence DOUBLE PRECISION NOT NULL DEFAULT 0.0,
    use_count INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Indexes for performance (Blueprint section 3.3)
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_category_id ON transactions(category_id);
CREATE INDEX idx_transactions_transacted_at ON transactions(transacted_at);
CREATE INDEX idx_pending_intents_user_id ON pending_intents(user_id);
CREATE INDEX idx_merchant_cache_user_merchant ON merchant_cache(user_id, merchant_key);
