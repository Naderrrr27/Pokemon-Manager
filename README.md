# Pokemon-Manager

This is a tiny web server built with Go and Gin that lets you view, add, update, and delete caught Pokemon — all stored in memory with no database needed.

---

## 🚀 How to Run It

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
    "id": 1,
    "name": "Pikachu",
    "type": "IDK",
    "level": 5
  }
]
```

---

### Add a New Pokémon

**Endpoint:**  
`POST /pokemons`

**Example:**

```bash
curl -X POST http://localhost:8080/pokemons \
  -H "Content-Type: application/json" \
  -d @data.json
```

**Sample `data.json`:**

```json
{
  "name": "Pikachu II",
  "type": "IDC",
  "level": 4
}
```

**Sample Response:**

```json
{
  "id": 2,
  "name": "Charmander",
  "type": "Fire",
  "level": 7
}
```

---

### Update an Existing Pokémon

**Endpoint:**  
`PUT /pokemons/{id}`

**Example:**

```bash
curl -X PUT http://localhost:8080/pokemons/1 \
  -H "Content-Type: application/json" \
  -d @data.json
```

**Sample `data.json`:**

```json
{
  "name": "Charmeleon",
  "type": "Fire",
  "level": 16
}
```

**Sample Response:**

```json
{
  "id": 1,
  "name": "Charmeleon",
  "type": "Fire",
  "level": 16
}
```

---

### Delete a Pokémon

**Endpoint:**  
`DELETE /pokemons/{id}`

**Example:**

```bash
curl -X DELETE http://localhost:8080/pokemons/1
```

**Sample Response:**

```json
{
  "Message": "Pokemon released successfully"
}
```

---

## Notes

- You can now **view**, **add**, **update**, and **delete** Pokemon.
- No external API support yet.
- Data is stored in memory, so it resets every time you restart the server.
- Validation is included: name and type must not be empty, and level must be greater than 0.

---

## Built With

- Go (Golang)
- Gin (Web Framework)

---