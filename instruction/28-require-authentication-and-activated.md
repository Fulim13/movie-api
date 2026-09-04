# To Test

```sh
# Unauthorized
curl -i localhost:4000/v1/movies/1

# Not Activated
set BODY '{"email": "alice@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/tokens/authentication
curl -i -H "Authorization: Bearer ZVX3WDV7V4CBOPPHNK7HETK3JY" localhost:4000/v1/movies/1

# Activated
set BODY '{"email": "faith@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/tokens/authentication
curl -H "Authorization: Bearer XS3YBZKAPXMEBK7BICC2CR4JZF" localhost:4000/v1/movies/5
curl -H "Authorization: Bearer XS3YBZKAPXMEBK7BICC2CR4JZF" localhost:4000/v1/movies/5

```
