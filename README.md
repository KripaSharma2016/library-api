# 📚 **library-api**

>This microservice provides RESTful APIs for managing books.


## 🚀 Build & Run

1. **Navigate to the `library-app` directory:**
   ```sh
   cd library-app
   ```
2. **Build the Docker containers:**
   ```sh
   docker-compose build
   ```
3. **Start the services:**
   ```sh
   docker-compose up
   ```

---

## 📖 API Endpoints

### 1. Create a Book
**Request:**
```sh
curl --location 'http://localhost:8080/books' \
  --header 'Content-Type: application/json' \
  --data '{
    "title": "Second Animal Book",
    "author": "Kripa Sharma",
    "isbn": "false"
  }'
```
**Response:**
```json
{
  "id": 2,
  "title": "Second Animal Book",
  "author": "Kripa Sharma",
  "isbn": "false"
}
```

---

### 2. Update a Book
**Request:**
```sh
curl --location --request PUT 'http://localhost:8080/books/2' \
  --header 'Content-Type: application/json' \
  --data '{
    "title": "Second Fiction Book",
    "author": "Kripa Sharma",
    "isbn": "false"
  }'
```
**Response:**
```json
{
  "id": 2,
  "title": "Second Fiction Book",
  "author": "Kripa Sharma",
  "isbn": "false"
}
```

---

### 3. List All Books
**Request:**
```sh
curl --location 'http://localhost:8080/books'
```
**Response:**
```json
[
  {
    "id": 1,
    "title": "First Animal Book",
    "author": "Kripa Shankar Sharma",
    "isbn": "true"
  },
  {
    "id": 2,
    "title": "Second Fiction Book",
    "author": "Kripa Sharma",
    "isbn": "false"
  }
]
```

---

### 4. Delete a Book
**Request:**
```sh
curl --location --request DELETE 'http://localhost:8080/books/1'
```
**Response:**
```
HTTP Status: 204 No Content
```
