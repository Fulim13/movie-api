# To Load Test

```sh
# To Turn off the rate limter for load testing
go run ./cmd/api -limiter-enabled=false

set BODY '{"email": "alice@example.com", "password": "pa55word"}'

brew install hey
hey -d "$BODY" -m "POST" http://localhost:4000/v1/tokens/authentication
```

View the Output of /debug/var when hey is running
