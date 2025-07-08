# Pokemon-Manager

This is a tiny web server built with Go and Gin that lets you view and add caught Pokemon — all stored in memory with no database needed.

---

## How to Run It

1. Make sure you have Go installed: https://go.dev/dl/

2. In your terminal:

```bash
go mod init Pokemon-Manager
go get github.com/gin-gonic/gin
go run main.go
```

The server will start at:  
**http://localhost:8080**

---

## What You Can Do

### Get All Pokémons

**Endpoint:**  
`GET /pokemons`

**Example:**

```bash
curl http://localhost:8080/pokemons
```

**Sample Response:**

```json
[
  {
  "id": "1",
  "name": "Omar",
  "type": "Mentor",
  "level": "100000"
  }
]
```

---

### Add a New Pokémon

**Endpoint:**  
`POST /pokemons`

**Example:**

```bash
curl -X POST http://localhost:8080/pokemons   -H "Content-Type: application/json"   -d @data.json
```

**Sample Response:**

```json
{
  "id": "1",
  "name": "Omar",
  "type": "Mentor",
  "level": "100000"
}
```

---

## 📝 Notes

- You can now **view** and **add** Pokemon.
- No update, delete, or external API yet.
- Data is stored in memory, so it resets every time you restart the server.

---

## 🔧 Built With

- Go (Golang)
- Gin (Web Framework)s

---