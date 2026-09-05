# To Test

```sh
curl -X POST -d '{"email": "alice@example.com"}' localhost:4000/v1/tokens/password-reset
{
    "message": "an email will be sent to you containing password reset instructions"
}
```

```sh
set BODY '{"password": "your new password", "token": "743TFC5MKZ4R7L3Q4Y5BBYY4NP"}'
curl -X PUT -d "$BODY" localhost:4000/v1/users/password
{
    "message": "your password was successfully reset"
}
```
