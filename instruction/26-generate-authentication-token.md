# To Test

```sh
set BODY '{"email": "faith@example.com", "password": "pa55word"}'
curl -i -d "$BODY" localhost:4000/v1/tokens/authentication

set BODY '{"email": "alice@example.com", "password": "wrong pa55word"}'
curl -i -d "$BODY" localhost:4000/v1/tokens/authentication
```
