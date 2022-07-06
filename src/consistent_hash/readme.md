# start proxy

```bash
go run main.go
```

# start cache server

```bash
cd server
go run main.go
go run main.go -p 8081
go run main.go -p 8082
go run main.go -p 8083
```

# Query

```bash
curl localhost:18888/key?key=123
```

# Consistent Hash with Load Bound Test

```bash
curl localhost:18888/key_least?key=123
```
