# To install JWT Token Package

```sh
go get github.com/pascaldekloe/jwt@v1
```

# To test

```sh
curl -X POST -d '{"email": "faith@example.com", "password": "pa55word"}' localhost:4000/v1/tokens/authentication

curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJncmVlbmxpZ2h0LmZ1bGltLm5ldCIsInN1YiI6IjI0IiwiYXVkIjpbImdyZWVubGlnaHQuZnVsaW0ubmV0Il0sImV4cCI6MTc4ODY5NzAxMy4wMzM0MzMyLCJuYmYiOjE3ODg2MTA2MTMuMDMzNDMzMiwiaWF0IjoxNzg4NjEwNjEzLjAzMzQzMzJ9.Y86ZlvfaXVopPTE_dLMnNj9cxUy2P6CbhVOmyx3839w" localhost:4000/v1/movies/6

curl -H "Authorization: Bearer INVALID" localhost:4000/v1/movies/2
```
