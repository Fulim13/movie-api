# To Create Token Table

```sh
migrate create -seq -ext .sql -dir ./migrations create_tokens_table
migrate -path ./migrations -database postgres://greenlight:pa55word@localhost/grseenlight?sslmode=disable up
```

# To Test

```sh
set BODY '{"name": "Faith Smith", "email": "faith@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/users
curl -X PUT -d '{"token": "invalid"}' localhost:4000/v1/users/activated
curl -X PUT -d '{"token": "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}' localhost:4000/v1/users/activated
curl -X PUT -d '{"token": "7WTYWERW6DRV3FI7ZF3XTOPA3W"}' localhost:4000/v1/users/activated
```
