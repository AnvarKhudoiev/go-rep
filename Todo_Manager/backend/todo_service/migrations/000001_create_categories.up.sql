CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(50),
    user_id INTEGER NOT NULL DEFAULT 0,
    is_default BOOLEAN DEFAULT FALSE
);
 
CREATE INDEX IF NOT EXISTS idx_categories_deleted_at ON categories(deleted_at);
 