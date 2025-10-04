# Tiny Go API

Run the server:

```powershell
go run .lab2/cmd/api
```

Examples:

Valid GET:

```powershell
curl -i -H "X-API-Key: secret123" "http://localhost:8080/user?id=42"
```

Missing or bad key:

```powershell
curl -i "http://localhost:8080/user?id=42"
```

Create user:

```powershell
curl -i -X POST -H "X-API-Key: secret123" -H "Content-Type: application/json" -d '{"name":"Alice"}' "http://localhost:8080/user"
```
