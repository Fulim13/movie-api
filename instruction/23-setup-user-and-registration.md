# Create Migration files for users

```sh
migrate create -seq -ext=.sql -dir=./migrations create_users_table
```

# Migrate the files

```sh
migrate -path=./migrations -database=postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable up
```

# Install Bcrypt

```sh
go get golang.org/x/crypto/bcrypt@latest
```

# Test the Endpoint

```sh
set BODY '{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'
curl -i -d "$BODY" localhost:4000/v1/users

set BODY '{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'
curl -i -d "$BODY" localhost:4000/v1/users

set BODY '{"name": "", "email": "bob@invalid.", "password": "pass"}'
curl -i -d "$BODY" localhost:4000/v1/users
```
