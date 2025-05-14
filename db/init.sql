DROP TABLE IF EXISTS price_stats;

CREATE TABLE IF NOT EXISTS price_stats (
    id SERIAL PRIMARY KEY,
    pair_name VARCHAR(50) NOT NULL,
    exchange VARCHAR(50) NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    average_price DECIMAL(24, 8) NOT NULL,
    min_price DECIMAL(24, 8) NOT NULL,
    max_price DECIMAL(24, 8) NOT NULL,
    UNIQUE(pair_name, exchange, timestamp)
);

CREATE INDEX idx_pair_timestamp ON price_stats(pair_name, timestamp);

CREATE INDEX idx_exchange_pair_timestamp ON price_stats(exchange, pair_name, timestamp);