# To install IP Rate Limiter Tools

```sh
go get golang.org/x/time/rate@latest
go get github.com/tomasen/realip@latest
```

# Rate Limiting concept

- We will have a bucket that starts with b tokens in it.
- Each time we receive an HTTP request, we will remove one token from the bucket.
- Every 1/r seconds, a token is added back to the bucket — up to a maximum of b total tokens.
- If we receive an HTTP request and the bucket is empty, then we should return a 429 Too Many Requests response.
- In practice this means that our application would allow a maximum ‘burst’ of b HTTP requests in quick succession, but over time it would allow an average of r requests per second.

# To Test the IP Rate Limiter

```sh
go run ./cmd/api/ -limiter-burst=2
for i in (seq 6); curl http://localhost:4000/v1/healthcheck; end
```
