# Migration For Permission Table

```sh
migrate create -seq -ext .sql -dir ./migrations add_permissions

migrate -path=./migrations -database=postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable up
```

# SQL

```sql
-- Set the activated field for alice@example.com to true.
UPDATE users SET activated = true WHERE email = 'alice@example.com';

-- Give all users the 'movies:read' permission
INSERT INTO users_permissions
SELECT id, (SELECT id FROM permissions WHERE code = 'movies:read') FROM users;

-- Give faith@example.com the 'movies:write' permission
INSERT INTO users_permissions
VALUES (
    (SELECT id FROM users WHERE email = 'faith@example.com'),
    (SELECT id FROM permissions WHERE  code = 'movies:write')
);

-- List all activated users and their permissions.
SELECT email, array_agg(permissions.code) as permissions
FROM permissions
INNER JOIN users_permissions ON users_permissions.permission_id = permissions.id
INNER JOIN users ON users_permissions.user_id = users.id
WHERE users.activated = true
GROUP BY email;
```

# To Test

```sh
# To Test the user only with read permission
set BODY '{"email": "alice@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/tokens/authentication
curl -H "Authorization: Bearer A6LEHGSJQG3JUAHZ7J3MSVJZVP " localhost:4000/v1/movies/5
curl -X DELETE -H "Authorization: Bearer A6LEHGSJQG3JUAHZ7J3MSVJZVP" localhost:4000/v1/movies/5

# To Test the user with read and write permissions
set BODY '{"email": "faith@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/tokens/authentication

curl -X DELETE -H "Authorization: Bearer FWHLX6AIFNJMEXV26RNDHCFF7I" localhost:4000/v1/movies/5
```

# Create New User automatically will have movies:read permission

```sh
set BODY '{"name": "Grace Smith", "email": "grace@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/users
```
