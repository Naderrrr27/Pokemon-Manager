# Pokemon-Manager

This is a tiny web server built with Go and Gin that lets you view, add, update, delete, and get info about caught Pokemon — all stored in memory with no database needed. 

It also supports fetching basic data from the external PokeAPI.


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

### Get All Pokemons

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

### Add a New Pokemon

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
  "name": "Pikachu II",
  "type": "IDC",
  "level": 4
}
```

---

### Update an Existing Pokemon

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

### Delete a Pokemon

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

### Get Pokemon Info from External API

**Endpoint:**  
`GET /pokemon-info/{name}`

**Example:**

```bash
curl http://localhost:8080/pokemon-info/pikachu
```

**Sample Response:**

```json
{
  "name": "pikachu",
  "height": 4,
  "weight": 60,
  "types": ["electric"]
}
```

---

## Notes

- You can now view, add, update, delete, and fetch external info for Pokemon.
- Data is stored in memory, so it resets every time you restart the server.
- Basic validation is included.

---

## Built With

- Go (Golang)
- Gin (Web Framework)

---
