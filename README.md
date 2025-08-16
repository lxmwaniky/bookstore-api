# 📚 Bookstore API - Public Playground

A **public learning API** built with Go and Gin for testing CRUD operations. Anyone can play with it!

🌍 **Live Demo**: http://vps.lxmwaniky.me:8080/api/books

## 🎯 Purpose
- **Learning Go** and API development
- **Public testing** of CRUD operations
- **No authentication** - open for everyone to use
- **Educational project** - not production code

## ⚠️ Important Notes
- This API is **intentionally public**
- Anyone can create, read, update, or delete books
- Data may be modified by other users
- Perfect for testing API clients, Postman, etc.

## 🚀 Features
- **CRUD Operations** for books
- **PostgreSQL** database
- **Docker** support
- **RESTful** API design
- **Input validation**

## 🛠️ Tech Stack
- **Go 1.24.6**
- **Gin** web framework
- **PostgreSQL** database
- **Docker & Docker Compose**

## 📋 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/books` | Get all books |
| `GET` | `/api/books/:id` | Get book by ID |
| `POST` | `/api/books` | Create new book |
| `PUT` | `/api/books/:id` | Update book |
| `DELETE` | `/api/books/:id` | Delete book |

## 🎮 Try It Out

### Get all books
```bash
curl http://vps.lxmwaniky.me:8080/api/books
```

### Create a book
```bash
curl -X POST http://vps.lxmwaniky.me:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Book",
    "author": "Test Author",
    "publication_year": 2024,
    "genre": "Fiction",
    "description": "A test book for learning"
  }'
```

### Update a book
```bash
curl -X PUT http://vps.lxmwaniky.me:8080/api/books/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Book Title",
    "author": "Updated Author"
  }'
```

### Delete a book
```bash
curl -X DELETE http://vps.lxmwaniky.me:8080/api/books/1
```

## 🏗️ Local Development Setup

### Prerequisites
- Go 1.24+
- Docker & Docker Compose
- Git

### 🐳 Option 1: Docker (Recommended)

#### 1. Clone the repository
```bash
git clone <your-repo-url>
cd bookstore-api
```

#### 2. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your database credentials
```

#### 3. Build and run everything with Docker
```bash
docker-compose up --build -d
```

#### 4. Initialize the database
```bash
docker exec -i bookstore-db psql -U $DB_USER -d $DB_NAME < init.sql
```

#### 5. Check status
```bash
docker-compose ps
docker-compose logs api
```

The API will be available at `http://localhost:8080`

### 🔧 Option 2: Local Go Development

#### 1. Start only the database
```bash
docker-compose up postgres -d
```

#### 2. Run the Go application locally
```bash
go mod tidy
go run main.go
```

### 🛠️ Docker Commands

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down

# Rebuild and restart
docker-compose up --build -d

# Check service status
docker-compose ps
```

## 📁 Project Structure
```
bookstore-api/
├── database/          # Database connection
├── handlers/          # HTTP request handlers
├── models/            # Data models
├── main.go            # Application entry point
├── init.sql           # Database initialization
├── docker-compose.yml # Docker services
├── Dockerfile         # Container configuration
├── .dockerignore      # Docker build exclusions
├── go.mod             # Go dependencies
└── .env.example       # Environment template
```

## 🔒 Security Notice
This is a **learning project** with intentionally open access:
- No authentication required
- No rate limiting
- No input sanitization
- **NOT suitable for production use**

## 🤝 Contributing
Feel free to:
- Test the API endpoints
- Report bugs or issues
- Suggest improvements
- Use as a learning resource

## 🎓 Learning Resources
- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/docs/)
- [Docker](https://docs.docker.com/)

---

**Happy Learning! 🚀**
