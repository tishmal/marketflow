DROP TABLE IF EXISTS market_aggregates;

CREATE TABLE IF NOT EXISTS market_aggregates (
    id SERIAL PRIMARY KEY,
    pair_name VARCHAR(50) NOT NULL,
    exchange VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP NOT NULL,
    average_price DECIMAL(24, 8) NOT NULL,
    min_price DECIMAL(24, 8) NOT NULL,
    max_price DECIMAL(24, 8) NOT NULL,
    UNIQUE(pair_name, exchange, timestamp)
);

CREATE INDEX idx_pair_timestamp ON market_aggregates(pair_name, timestamp);

CREATE INDEX idx_exchange_pair_timestamp ON market_aggregates(exchange, pair_name, timestamp);