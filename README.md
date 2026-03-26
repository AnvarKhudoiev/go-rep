# Rest API practice

Build a REST API to manage books, authors, and categories, using in-memory storage (slices and maps). The project will include CRUD operations, pagination, filters, and data validation.

### 📊 **API Features:**

1. **Book Management**
    - `GET /books` – List all books (with pagination & filters)
    - `POST /books` – Add a new book
    - `GET /books/:id` – Get a book by ID
    - `PUT /books/:id` – Update a book
    - `DELETE /books/:id` – Delete a book
2. **Author Management**
    - `GET /authors` – List all authors
    - `POST /authors` – Add a new author
3. **Category Management**
    - `GET /categories` – List all categories
    - `POST /categories` – Add a new category
4. **Advanced Features**
    - Pagination and filters (e.g., `/books?category=Fiction&page=1`)
    - Input validation (e.g., required fields, minimum price)

### 📂 **Project Structure:**

```
bookstore/
├── main.go
├── models/
│    ├── book.go
│    ├── author.go
│    └── category.go
└── handlers/
     ├── book_handler.go
     ├── author_handler.go
     └── category_handler.go

```

### 📘 **Step 1: Define Models**

**models/book.go**

```go
package models

type Book struct {
    ID         int    `json:"id"`
    Title      string `json:"title"`
    AuthorID   int    `json:"author_id"`
    CategoryID int    `json:"category_id"`
    Price      float64 `json:"price"`
}

```

**models/author.go**

```go
package models

type Author struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

```

**models/category.go**

```go
package models

type Category struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

```
