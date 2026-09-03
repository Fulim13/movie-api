# To install the package

```sh
go get github.com/wneessen/go-mail@latest
```

# To test the email

```sh
set BODY '{"name": "Dave Smith", "email": "dave@example.com", "password": "pa55word"}'
curl -w '\nTime: %{time_total}\n' -d "$BODY" localhost:4000/v1/users
```

# To test graceful shutdown of background sending email

```sh
set BODY '{"name": "Edith Smith", "email": "edith@example.com", "password": "pa55word"}'
curl -d "$BODY" localhost:4000/v1/users & pkill -SIGTERM api &
```
