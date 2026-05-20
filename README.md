# 📝 Todo Manager

A full-stack task management application built with **Go microservices** and **React** frontend.

---

## 🏗 Architecture

```
Todo Manager
├── backend/
│   ├── auth_service/       # Authentication microservice (port 8081)
│   ├── todo_service/       # Todo & Category microservice (port 8082)
│   └── docker-compose.yml
└── frontend/
    └── client/             # React + Vite + shadcn/ui (port 3000)
```

### Microservices Communication

```
Browser → Frontend (3000)
            ↓
    Auth Service (8081) ←──── JWT Token ────→ Todo Service (8082)
            ↓                                         ↓
        auth_db (5441)                          todo_db (5442)
       [PostgreSQL]                             [PostgreSQL]
```

- **Auth Service** issues JWT tokens on login/register
- **Todo Service** validates JWT via middleware and calls Auth Service via **Resty v2** to fetch user info
- Services communicate internally via Docker network

---

## 🛠 Tech Stack

### Backend
| Technology | Usage |
|-----------|-------|
| Go 1.25 | Primary language |
| Gin | HTTP framework |
| GORM | ORM for PostgreSQL |
| PostgreSQL 15 | Database |
| golang-migrate | Database migrations |
| golang-jwt/jwt v5 | JWT authentication |
| gin-contrib/cors | CORS middleware |
| go-resty/resty v2 | Inter-service HTTP client |
| godotenv | Environment variables |
| bcrypt | Password hashing |

### Frontend
| Technology | Usage |
|-----------|-------|
| React 19 | UI framework |
| TypeScript | Type safety |
| Vite | Build tool |
| Tailwind CSS v4 | Styling |
| shadcn/ui | Component library |
| React Router v7 | Routing |
| Recharts | Statistics charts |
| @dnd-kit | Drag & drop |
| next-themes | Dark/light mode |
| sonner | Toast notifications |

### DevOps
| Technology | Usage |
|-----------|-------|
| Docker | Containerization |
| Docker Compose | Multi-container orchestration |
| nginx | Frontend static file serving |

---

## 🚀 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started) & Docker Compose
- [Go 1.25+](https://go.dev/dl/) (for local development)
- [Node.js 18+](https://nodejs.org/) (for local frontend development)

### Run with Docker (recommended)

```bash
# Clone the repository
git clone https://github.com/yourusername/Todo_Manager.git
cd Todo_Manager/backend

# Start all services
docker-compose up --build
```

| Service | URL |
|---------|-----|
| Frontend | http://localhost:3000 |
| Auth API | http://localhost:8081 |
| Todo API | http://localhost:8082 |

### Run locally (without Docker)

**1. Start PostgreSQL databases**
```bash
# auth_db on port 5441
# todo_db on port 5442
# Or use your own PostgreSQL and update .env files
```

**2. Configure environment variables**

`backend/auth_service/.env`:
```env
DB_HOST=localhost
DB_PORT=5441
DB_USER=user
DB_PASSWORD=password
DB_NAME=auth_db
JWT_SECRET=your_super_secret_key
```

`backend/todo_service/.env`:
```env
DB_HOST=localhost
DB_PORT=5442
DB_USER=user
DB_PASSWORD=password
DB_NAME=todo_db
JWT_SECRET=your_super_secret_key
```

**3. Run services**
```bash
cd backend

# Auth Service
go run auth_service/main.go &

# Todo Service
go run todo_service/main.go &
```

**4. Run frontend**
```bash
cd frontend/client
npm install
npm run dev
```

---

## 📁 Project Structure

```
backend/
├── auth_service/
│   ├── config/
│   │   └── db.go               # DB connection + migrations
│   ├── handlers/
│   │   ├── auth.go             # Register, Login, Logout, GetUserByID
│   │   └── auth_test.go        # 10 unit tests
│   ├── migrations/
│   │   ├── 000001_create_users.up.sql
│   │   └── 000001_create_users.down.sql
│   ├── models/
│   │   └── user.go
│   ├── Dockerfile
│   └── main.go
│
├── todo_service/
│   ├── clients/
│   │   └── user_client.go      # Resty v2 — calls auth_service
│   ├── config/
│   │   └── db.go               # DB connection + migrations
│   ├── handlers/
│   │   ├── todo_handler.go     # Todo CRUD + Overdue
│   │   ├── category_handler.go # Category CRUD
│   │   ├── stats_handler.go    # Statistics
│   │   └── todo_handler_test.go # 10 unit tests
│   ├── middleware/
│   │   └── auth.go             # JWT validation middleware
│   ├── migrations/
│   │   ├── 000001_create_categories.up.sql
│   │   ├── 000001_create_categories.down.sql
│   │   ├── 000002_create_tags.up.sql
│   │   ├── 000002_create_tags.down.sql
│   │   ├── 000003_create_todos.up.sql
│   │   └── 000003_create_todos.down.sql
│   ├── models/
│   │   ├── todo.go
│   │   ├── category.go
│   │   └── tag.go
│   ├── Dockerfile
│   └── main.go
│
└── docker-compose.yml

frontend/client/
├── src/
│   ├── api/
│   │   ├── Auth.ts             # Auth API functions
│   │   └── Todo.ts             # Todo/Category API functions
│   ├── app/
│   │   ├── auth/page.tsx       # Login/Register page
│   │   └── dashboard/page.tsx  # Dashboard with routing
│   ├── components/
│   │   ├── todos/
│   │   │   ├── TodoList.tsx
│   │   │   ├── AddTodoPage.tsx
│   │   │   ├── EditTodoModal.tsx
│   │   │   └── OverduePage.tsx
│   │   ├── profile/
│   │   │   └── ProfilePage.tsx # Statistics & charts
│   │   └── ThemeToggle.tsx
│   ├── shadcn/components/      # shadcn/ui components
│   ├── store/
│   │   ├── authStore.tsx       # Auth context
│   │   └── todoStore.tsx       # Todo context
│   └── App.tsx
├── Dockerfile
└── nginx.conf
```

---

## 📡 API Reference

### Auth Service — `http://localhost:8081`

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | `/register` | ❌ | Register new user |
| POST | `/login` | ❌ | Login, returns JWT token |
| POST | `/logout` | ✅ | Logout, blacklists token |
| GET | `/users/:id` | ✅ | Get user by ID |

#### Register
```http
POST /register
Content-Type: application/json

{
  "username": "john",
  "password": "secret123"
}
```
```json
{
  "message": "User registered successfully",
  "user_id": 1
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
  "username": "john",
  "password": "secret123"
}
```
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### Todo Service — `http://localhost:8082`

> All `/api/*` routes require header: `Authorization: Bearer <token>`

#### Todos

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/todos` | Get all todos for current user |
| GET | `/api/todos/overdue` | Get overdue incomplete todos |
| GET | `/api/todos/:id` | Get todo by ID (includes author) |
| POST | `/api/todos` | Create todo |
| PUT | `/api/todos/:id` | Update todo |
| DELETE | `/api/todos/:id` | Delete todo (soft delete) |

#### Create Todo
```http
POST /api/todos
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Submit report",
  "description": "Q3 quarterly report",
  "priority": "high",
  "due_date": "2026-05-30T15:00:00Z",
  "category_id": 1,
  "tags": [
    { "name": "work" },
    { "name": "urgent" }
  ]
}
```

**Priority values:** `low` | `medium` (default) | `high`

#### Categories

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/categories` | Get default + user categories |
| GET | `/api/categories/:id` | Get category by ID |
| POST | `/api/categories` | Create custom category |
| PUT | `/api/categories/:id` | Update category (own only) |
| DELETE | `/api/categories/:id` | Delete category (own only) |

> Default categories (Личное, Работа, Учёба, Здоровье, Покупки) cannot be edited or deleted.

#### Statistics

```http
GET /api/stats
Authorization: Bearer <token>
```
```json
{
  "total": 15,
  "completed": 6,
  "active": 9,
  "overdue": 3,
  "completion_rate": 40,
  "week": [
    { "date": "2026-05-14", "label": "Чт", "created": 2, "completed": 1 }
  ],
  "by_category": [
    { "name": "Работа", "color": "#F44336", "count": 5 }
  ]
}
```

---

## 🗄 Database Schema

### auth_db

```sql
users
├── id          SERIAL PRIMARY KEY
├── created_at  TIMESTAMP WITH TIME ZONE
├── updated_at  TIMESTAMP WITH TIME ZONE
├── deleted_at  TIMESTAMP WITH TIME ZONE
├── username    VARCHAR(100) UNIQUE NOT NULL
└── password    VARCHAR(255) NOT NULL
```

### todo_db

```sql
categories
├── id          SERIAL PRIMARY KEY
├── name        VARCHAR(100) NOT NULL
├── color       VARCHAR(20)
├── user_id     INTEGER DEFAULT 0
└── is_default  BOOLEAN DEFAULT FALSE

tags
├── id    SERIAL PRIMARY KEY
└── name  VARCHAR(100) UNIQUE NOT NULL

todos
├── id           SERIAL PRIMARY KEY
├── title        VARCHAR(255) NOT NULL
├── description  TEXT
├── completed    BOOLEAN DEFAULT FALSE
├── user_id      INTEGER NOT NULL
├── category_id  INTEGER REFERENCES categories(id) ON DELETE SET NULL
├── priority     VARCHAR(10) DEFAULT 'medium'
└── due_date     TIMESTAMP WITH TIME ZONE

todo_tags (junction table)
├── todo_id  INTEGER REFERENCES todos(id) ON DELETE CASCADE
└── tag_id   INTEGER REFERENCES tags(id) ON DELETE CASCADE
```

---

## 🧪 Running Tests

```bash
cd backend

# Auth service tests (10 tests)
go test ./auth_service/handlers/... -v

# Todo service tests (10 tests)
go test ./todo_service/handlers/... -v

# All tests
go test ./... -v
```

### Test Coverage

| Service | Tests | Coverage |
|---------|-------|----------|
| auth_service | 10 | Register, Login, Logout, Blacklist, GetUser |
| todo_service | 10 | CreateTodo, GetTodos, GetById, Update, Delete, Categories |

---

## 🔐 Authentication Flow

```
1. POST /register → user created in auth_db
                  → default categories seeded in todo_db

2. POST /login    → bcrypt password check
                  → JWT signed with HS256 (expires in 24h)
                  → token returned to client

3. Request        → Authorization: Bearer <token>
                  → AuthMiddleware validates signature + expiry
                  → user_id extracted from claims
                  → passed to handler via context

4. POST /logout   → token added to in-memory blacklist
                  → subsequent requests with this token → 401
```

---

## 🐳 Docker

### Services

```yaml
auth_service   → port 8081  (depends on auth_db)
todo_service   → port 8082  (depends on todo_db + auth_service)
frontend       → port 3000  (nginx serving React build)
auth_db        → port 5441  (PostgreSQL, volume: auth_db_data)
todo_db        → port 5442  (PostgreSQL, volume: todo_db_data)
```

### Useful commands

```bash
# Start everything
docker-compose up --build

# Start in background
docker-compose up -d --build

# Stop (data preserved)
docker-compose down

# Stop and delete all data
docker-compose down -v

# View logs
docker-compose logs -f todo_service
docker-compose logs -f auth_service

# Rebuild single service
docker-compose up --build todo_service
```

---

## 🗺 Database Migrations

Migrations are managed by **golang-migrate** and run automatically on service startup.

```bash
# Manual migration (from backend/ directory)
go run auth_service/main.go   # applies auth_service/migrations/
go run todo_service/main.go   # applies todo_service/migrations/
```

Migration files follow the naming convention:
```
000001_create_users.up.sql    ← apply
000001_create_users.down.sql  ← rollback
```

---

## 🖥 Frontend Features

- **Authentication** — Login / Register with JWT
- **Task Management** — Create, edit, delete, complete tasks
- **Priorities** — Low / Medium / High with color indicators
- **Due Dates** — Deadline picker, overdue tasks highlighted
- **Categories** — Default + custom categories with colors
- **Tags** — Add multiple tags per task
- **Overdue Section** — Dedicated view for overdue tasks
- **Statistics** — Bar chart (7-day activity) + Donut chart (by category)
- **Dark / Light Mode** — Theme toggle in sidebar
- **Responsive** — Works on desktop and mobile

---

## 📮 Postman Collection

Import `Todo_Manager.postman_collection.json` into Postman.

**Run order:**
1. Register
2. **Login** ← token auto-saved to `{{token}}`
3. Get User By ID
4. Get All Todos
5. Create Todo (simple)
6. Create Todo (with priority + due_date)
7. Create Todo (with tags)
8. Get Todo By ID
9. Update Todo
10. Toggle completed
11. Get Overdue Todos
12. Get All Categories
13. Get Category By ID
14. Create Category
15. Update Category
16. Delete Category
17. Get Statistics
18. Internal - Default Categories
19. Delete Todo
20. **Logout** ← run last

> ⚠️ Run **Login** before any authenticated requests. Run **Logout** last — it invalidates the token.

---

## 👤 Author

**Anvar Shokhkhudoiev**

---

## 📄 License

MIT
