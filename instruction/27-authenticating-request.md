# To Test

```sh
curl -d '{"email": "alice@example.com", "password": "pa55word"}' localhost:4000/v1/tokens/authentication

curl -H "Authorization: Bearer FXCZM44TVLC6ML2NXTOW5OHFUE" localhost:4000/v1/healthcheck

curl -i -H "Authorization: Bearer XXXXXXXXXXXXXXXXXXXXXXXXXX" localhost:4000/v1/healthcheck

curl -i -H "Authorization: INVALID" localhost:4000/v1/healthcheck
```
