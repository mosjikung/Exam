# IT 06-1 — Product Code Management System

ระบบจัดการรหัสสินค้า พร้อม Barcode Code 39

## Stack

| Layer | Technology |
|-------|-----------|
| Frontend | Angular 17 (Standalone Components) |
| Backend | Golang + Fiber v2 |
| Database | PostgreSQL 16 |
| Barcode | Code 39 — HTML Canvas |
| Reverse Proxy | Nginx |
| Container | Docker + Docker Compose |
| DB GUI | pgAdmin 4 |

---

## Project Structure

```
├── .gitignore
├── docker-compose.yml
├── init.sql
├── backend/
│   ├── cmd/api/main.go              # Entrypoint
│   ├── internal/
│   │   ├── config/                  # Environment config
│   │   ├── database/                # PostgreSQL connection
│   │   ├── domain/product/          # Domain model
│   │   ├── handler/                 # HTTP handlers
│   │   ├── middleware/              # CORS, logger, error handler
│   │   ├── repository/              # Data access layer
│   │   └── service/                 # Business logic
│   ├── pkg/validator/               # Validation
│   ├── go.mod
│   ├── .env
│   └── Dockerfile
└── frontend/
    ├── src/
    │   └── app/
    │       ├── app.component.ts/html/scss
    │       ├── barcode.component.ts  # Code 39 renderer
    │       ├── product.service.ts    # HTTP Service
    │       └── app.config.ts
    ├── nginx.conf
    ├── angular.json
    ├── package.json
    └── Dockerfile
```

---

## Getting Started

### Prerequisites
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

### Run

```bash
docker compose up --build
```

| Service | URL |
|---------|-----|
| 🌐 Frontend | http://localhost:4200 |
| ⚙️ Backend API | http://localhost:3000 |
| 🗄️ pgAdmin | http://localhost:5050 |
| 🐘 PostgreSQL | localhost:5432 |

### pgAdmin Login
| | |
|---|---|
| Email | admin@it06.com |
| Password | admin1234 |

---

## API Endpoints

| Method | URL | Description |
|--------|-----|-------------|
| `GET` | `/api/products` | ดึงรายการสินค้าทั้งหมด |
| `POST` | `/api/products` | เพิ่มสินค้าใหม่ |
| `DELETE` | `/api/products/:id` | ลบสินค้า |
| `GET` | `/health` | Health check |

### POST /api/products
```json
{ "product_code": "AB12-CD34-EF56-GH78" }
```

---

## Database Schema

```sql
CREATE TABLE products (
    id           SERIAL      PRIMARY KEY,
    product_code VARCHAR(19) NOT NULL,

    CONSTRAINT uni_products_product_code UNIQUE (product_code),
    CONSTRAINT chk_product_code_length   CHECK (LENGTH(product_code) = 19),
    CONSTRAINT chk_product_code_format   CHECK (
        product_code ~ '^[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}-[A-Z0-9]{4}$'
    )
);
```

### Validation Rules
| Rule | Detail |
|------|--------|
| Format | `XXXX-XXXX-XXXX-XXXX` |
| Length | 16 ตัวอักษร + 3 ขีด = 19 ตัว |
| Characters | A-Z และ 0-9 เท่านั้น |
| Unique | ห้ามซ้ำกัน |

---

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `DB_HOST` | `postgres` | PostgreSQL host |
| `DB_PORT` | `5432` | PostgreSQL port |
| `DB_USER` | `it06user` | DB username |
| `DB_PASSWORD` | `it06pass` | DB password |
| `DB_NAME` | `it06db` | Database name |
| `PORT` | `3000` | API server port |

---

## Useful Commands

```bash

docker compose up --build


docker compose up


docker compose down


docker compose down -v


docker compose logs backend
docker compose logs frontend


docker compose ps
```
