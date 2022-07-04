# MyApp is a trivial REST backend to use so I can have something to deploy (in order learn infra stuff)

### To update dependencies based on imports in your project (from root folder):
```bash
go get .
```

### To run tests (from root folder):
```bash
go test .
```

### To run the server (from root folder):
```bash
go run .
```

### Or:
```bash
go run main.go
```

### To tidy the mod file:
```bash
go mod tidy
```

```bash
curl http://localhost:8080/records

curl http://localhost:8080/records \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "5f26f150-3df3-443c-b7f6-5a9b2b611258","title": "Reckless","artist": "Todd Helder","price": 22.29}'

curl http://localhost:8080/records/2

curl http://localhost:8080/records/858c98db-10c3-4959-ad68-0c8c526dca07
```
