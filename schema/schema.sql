CREATE TABLE urls(
    id           BIGSERIAL PRIMARY KEY,
    original_url VARCHAR(10) NOT NULL,
    short_code   TEXT NOT NULL UNIQUE,
    expires_at   TIMESTAMPTZ,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_urls_short_code ON urls(short_code);
