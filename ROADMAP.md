# ScaleURL — My Project Roadmap

> **Goal:** Build a production-grade URL shortener that can handle 100M+ requests.

---

## My Progress Overview

| Phase   | Focus               | Status         |
| ------- | ------------------- | -------------- |
| Phase 1 | Core API Foundation | ⬜ Not started |
| Phase 2 | Auth & Middleware   | ⬜ Not started |
| Phase 3 | Analytics Engine    | ⬜ Not started |
| Phase 4 | Scale & Performance | ⬜ Not started |
| Phase 5 | Production Infra    | ⬜ Not started |
| Phase 6 | Advanced Features   | ⬜ Not started |
| Phase 7 | Scale Architecture  | ⬜ Not started |

---

## Phase 1 — Core API Foundation

**Goal:** Get a working REST API with URL shortening, redirect, and delete.

### Setup

- [ ] Initialize Go module (`go mod init`)
- [ ] Install Gin or Chi router
- [ ] Set up project folders (`cmd/`, `internal/`, `config/`)
- [ ] Load environment variables with `.env` (using cleanenv)

### Database

- [ ] Spin up PostgreSQL (local or Docker)
- [ ] Create `urls` table
- [ ] Connect to DB using `GORM`
- [ ] Write a DB connection helpers

### Core Handlers

- [ ] `POST /shorten` — accept a long URL, return a short code
- [ ] `GET /:code` — redirect to the original URL
- [ ] `DELETE /:code` — delete a short URL
- [ ] Add input validation (empty URL, invalid format)
- [ ] Return proper error responses (400, 404, 500 with JSON)

### Code Quality

- [ ] Use structured logging (`zerolog` or `zap`)
- [ ] Keep JSON responses consistent: `{ data, error, status }`
- [ ] Write basic unit tests for handler logic

---

## Phase 2 — Auth & User System

**Goal:** Let users register, log in, and own their short URLs.

### User System

- [ ] Create a `users` table
- [ ] `POST /register` — hash password with bcrypt, store user
- [ ] `POST /login` — verify password, return JWT (access + refresh tokens)
- [ ] `POST /refresh` — rotate refresh token, return new access token
- [ ] `POST /logout` — invalidate refresh token

### JWT Auth

- [ ] Sign JWT with a secret key (RS256 or HS256)
- [ ] Write JWT - sign, verify, parse claims
- [ ] Add auth middleware
- [ ] Attach `user_id` to context after token validation

### Ownership

- [ ] Link `urls` table to `users`
- [ ] Protect delete endpoint — users can only delete their own URLs
- [ ] `GET /my-urls` — list all URLs for the authenticated user

### Rate Limiting

- [ ] Rate limit by IP for guest users
- [ ] Rate limit by `user_id` for authenticated users
- [ ] Return `429 Too Many Requests` with `Retry-After` header

### API Key Auth

- [ ] Generate an API key on user creation
- [ ] Accept `X-API-Key` header as an alternative to JWT
- [ ] Useful for programmatic access without login

---

## Phase 3 — Analytics Engine

**Goal:** Track every redirect and expose analytics endpoints.

### Click Tracking

- [ ] Create a `clicks` table
- [ ] Record clicks on every redirect (async — don’t slow down redirects)
- [ ] Parse `User-Agent` for device/browser info
- [ ] Use IP geolocation (free API or MaxMind GeoLite2)

### Analytics Endpoints

- [ ] `GET /urls/:code/stats` — total clicks, unique clicks
- [ ] `GET /urls/:code/stats/timeline` — clicks grouped by day/week/month
- [ ] `GET /urls/:code/stats/referrers` — top referrer domains
- [ ] `GET /urls/:code/stats/geo` — clicks by country
- [ ] `GET /urls/:code/stats/devices` — breakdown by device type

### Queries

- [ ] Use PostgreSQL `tsvector` or `GROUP BY` for aggregations
- [ ] Add indexes on `url_id` and `clicked_at` for fast queries
- [ ] Write query helpers

---

## Phase 4 — Scale & Performance

**Goal:** Achieve sub-millisecond redirects and prove the system can handle load.

### Redis Cache

- [ ] Add Redis to Docker Compose
- [ ] On shorten: write `short_code → original_url` to Redis with TTL
- [ ] On redirect: check Redis first, fallback to DB (read-through cache)
- [ ] On delete: invalidate Redis key
- [ ] On update: invalidate and re-set Redis key
- [ ] Write cache helpers

### Short Code Generation

- [ ] Implement Base62 encoding
- [ ] Use auto-increment DB ID → encode to Base62 (no collisions at scale)
- [ ] Or use Snowflake ID for distributed-safe unique IDs
- [ ] Document the choice in `DESIGN.md`

### DB Optimization

- [ ] Use `pgxpool` for DB connection pooling
- [ ] Add indexes: `short_code` (unique), `user_id`, `created_at`
- [ ] Use `EXPLAIN ANALYZE` to verify query plans

### Load Testing

- [ ] Install `k6` or `wrk`
- [ ] Write a load test script targeting `GET /:code` (redirect endpoint)
- [ ] Run test: simulate 10k concurrent users
- [ ] Record results (RPS, p50/p95/p99 latency) in `DESIGN.md`

---

## Phase 5 — Production Infra

**Goal:** Make the app deployable, observable, and CI/CD-ready.

### Docker

- [ ] Write a multi-stage `Dockerfile` (builder + minimal runtime)
- [ ] Write `docker-compose.yml` (app + postgres + redis)
- [ ] Verify the app starts cleanly with `docker compose up`
- [ ] Use `.env` for secrets — never hardcode

### Health & Observability

- [ ] `GET /health` — returns 200 if the app is running
- [ ] `GET /ready` — returns 200 only if DB + Redis are reachable
- [ ] Add request ID to every log line (trace requests end-to-end)
- [ ] Graceful shutdown — drain in-flight requests on SIGTERM

---

## Phase 6 — Advanced Features

**Goal:** Add features to make the project feel real and polished.

### Custom Slugs

- [ ] Let users provide a custom short code on creation
- [ ] Validate: alphanumeric, max 32 chars, not already taken
- [ ] Return `409 Conflict` if the slug is taken

### Link Expiry

- [ ] Add `expires_at` field to `urls` table
- [ ] Accept optional `ttl` (seconds) on shorten request
- [ ] Set Redis key TTL to match expiry
- [ ] Return `410 Gone` if the link is expired on redirect

### QR Code Generation

- [ ] `GET /urls/:code/qr` — returns a QR code PNG for the short URL
- [ ] Use a Go QR library (`skip2/go-qrcode`)
- [ ] Write the generator in `internal/shortener/qr.go`

### Password-Protected Links

- [ ] Add optional `password` field on URL creation (hashed with bcrypt)
- [ ] On redirect: if password is set, require `?pwd=` query param
- [ ] Return `401` if the password is wrong/missing

### Webhooks

- [ ] Accept optional `webhook_url` on URL creation
- [ ] On every redirect: fire a POST to `webhook_url` with click data (async)
- [ ] Retry on failure (up to 3 attempts)
- [ ] Write the logic

---

## Phase 7 — Scale Architecture

**Goal:** Design and document a system that can handle 100M+ requests.

### Async Analytics Queue

- [ ] Add RabbitMQ or Redis Streams to Docker Compose
- [ ] On redirect: publish click event to queue (non-blocking)
- [ ] Write a consumer
- [ ] Consumer reads from queue, writes to DB in batches
- [ ] Redirect speed is no longer affected by analytics writes

### Load Tests at Scale

- [ ] Run a load test: 100k concurrent users on the redirect endpoint
- [ ] Measure RPS with cache hot vs. cache cold
- [ ] Record p50/p95/p99 latency in `docs/DESIGN.md`
- [ ] Compare results before and after Redis caching

### Integration Tests

- [ ] Write end-to-end tests in `tests/integration/`
- [ ] Cover: shorten → redirect → analytics → delete flow
- [ ] Run against a test DB (separate from dev DB)

### DESIGN.md (Critical for Interviews)

- [ ] Why Base62 over UUID for short codes
- [ ] How cache + DB consistency is maintained on delete/update
- [ ] How the system scales horizontally (stateless app + shared Redis/DB)
- [ ] Async analytics write design and tradeoffs
- [ ] Load test results with actual numbers

---

## My Tech Stack

| Layer            | Choice                    |
| ---------------- | ------------------------- |
| Language         | Go                        |
| Router           | Gin or Chi                |
| Database         | PostgreSQL (GORM)         |
| Cache            | Redis                     |
| Auth             | JWT (HS256) + bcrypt      |
| Queue            | RabbitMQ or Redis Streams |
| Containerization | Docker + Docker Compose   |
| CI/CD            | GitHub Actions            |
| Load Testing     | k6 or wrk                 |
| Logging          | zerolog or zap            |

---

## Interview Talking Points

Questions I should be ready to answer after each phase:

- **Phase 1–2:** How does my middleware chain work? How do I handle JWT expiry?
- **Phase 3:** How do I track analytics without slowing down redirects?
- **Phase 4:** Walk me through a redirect — cache hit vs. miss.
- **Phase 4:** Why Base62? Why not UUID? How do I avoid collisions?
- **Phase 5:** How does my app handle graceful shutdown?
- **Phase 6:** How do I handle a webhook that keeps failing?
- **Phase 7:** How would I scale this to handle 1 billion requests a day?
