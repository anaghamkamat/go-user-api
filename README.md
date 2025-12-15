# Go User API – DOB & Dynamic Age

A production-ready RESTful API built with **GoFiber**, **PostgreSQL**, and **SQLC** to manage users with `name` and `dob` (date of birth). The API calculates **age dynamically** using Go’s `time` package when fetching users.

---

## ✨ Features

* CRUD APIs for users
* DOB stored in DB; **age computed dynamically**
* Clean architecture (handler → service → repository)
* SQLC-generated DB layer
* Input validation with `go-playground/validator`
* Structured logging with **Uber Zap**
* Middleware for request logging and request IDs

---

## 🗂️ Project Structure

```
/cmd/server/main.go
/config/
/db/migrations/
/db/sqlc/<generated>
/internal/
├── handler/
├── repository/
├── service/
├── routes/
├── middleware/
├── models/
└── logger/
```

---

## 🔧 Tech Stack

* GoFiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator

---

## 📊 Database Schema

```sql
CREATE TABLE users (
  id  SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob  DATE NOT NULL
);
```

---

## 🚀 Getting Started

### Prerequisites

* Go (>= 1.21 recommended)
* PostgreSQL
* sqlc

### Setup

1. **Clone the repo**

```bash
git clone <your-repo-url>
cd go-user-api
```

2. **Configure Database**

* Create database `userdb`
* Update the DSN in `cmd/server/main.go` (URL-encode password if it has special characters)

3. **Run migrations** (if applicable)

```bash
psql -U postgres -d userdb -f db/migrations/001_init.sql
```

4. **Generate SQLC code**

```bash
sqlc generate
```

5. **Run the server**

```bash
go run cmd/server/main.go
```

Server starts at `http://localhost:8080`.

---

## 🧪 API Endpoints

### Create User

**POST** `/users`

```json
{ "name": "Alice", "dob": "1990-05-10" }
```

### Get User by ID

**GET** `/users/:id`

```json
{ "id": 1, "name": "Alice", "dob": "1990-05-10", "age": 35 }
```

### Update User

**PUT** `/users/:id`

```json
{ "name": "Alice Updated", "dob": "1991-03-15" }
```

### Delete User

**DELETE** `/users/:id`

* Returns `204 No Content`

### List Users

**GET** `/users`

```json
[{ "id": 1, "name": "Alice", "dob": "1990-05-10", "age": 35 }]
```
GET /users?page=1&limit=10

---

## 🧠 Design Notes

* **Age is not stored** to avoid stale data; it’s computed at read time.
* Validation is enforced at the handler layer.
* Logging uses structured Zap logs for observability.

---

## 📦 Optional Enhancements

* Docker & docker-compose
* Pagination on `/users`
* Unit tests for age calculation

---

## 👩‍💻 Author

**Anagha Kamat**

---

## 📄 License

MIT
