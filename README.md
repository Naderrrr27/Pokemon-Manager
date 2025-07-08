# Pokemon-Manager

This is a tiny web server built with Go and Gin that lets you view a list of caught Pokémon — all stored in memory, no database needed. 

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
`localhost8080/pokemons`

**Example (using curl):**
```bash
curl http://localhost:8080/pokemons
```

**What you’ll get:**

```json
[
  {
    "id": "1",
    "name": "Pikachu",
    "type": "Electric",
    "level": "10"
  },
  {
    "id": "2",
    "name": "Charmander",
    "type": "Fire",
    "level": "8"
  }
]
```

---

## Notes

- Right now, the API only supports viewing Pokémon (`/pokemons`).
- You can't add, update, or delete yet, or deal with external api!
- Data is stored in memory, so it resets every time you restart the server.

---

## Built With

- Go (Golang)
- Gin (Web Framework)

---
