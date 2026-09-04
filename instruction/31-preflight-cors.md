Simple CORS (Must be meet three of them)

1. The request HTTP method is one of the three CORS-safe methods: HEAD, GET or POST.
2. The request headers are all either forbidden headers or one of the four CORS-safe headers:
   Accept
   Accept-Language
   Content-Language
   Content-Type
3. The value for the Content-Type header (if set) is one of:
   application/x-www-form-urlencoded
   multipart/form-data
   text/plain

Preflight CORS
request that don't follow above rules

# To Test Preflight

```sh
# FE
go run ./cmd/examples/cors/preflight

# BE
go run ./cmd/api -cors-trusted-origins="http://localhost:9000"
```

Output: CORS Blocked
Two Requests blocked by browser

1. An OPTIONS /v1/tokens/authentication request (this is the preflight request).
2. A POST /v1/tokens/authentication request (this is the ‘real’ request).

# To Identify Preflight Request

Preflight requests always have three components:

1. the HTTP method OPTIONS,
2. an Origin header, and
3. an Access-Control-Request-Method header.
   If any one of these pieces is missing, we know that it is not a preflight request.

Once we identify that it is a preflight request, we need to send a 200 OK response with some special headers to let the browser know whether or not it’s OK for the real request to proceed. These are:

1. An Access-Control-Allow-Origin response header, which reflects the value of the preflight request’s Origin header (just like in the previous chapter).
2. An Access-Control-Allow-Methods header listing the HTTP methods that can be used in real cross-origin requests to the URL.
3. An Access-Control-Allow-Headers header listing the request headers that can be included in real cross-origin requests to the URL.

```
Access-Control-Allow-Origin: <reflected trusted origin>
Access-Control-Allow-Methods: OPTIONS, PUT, PATCH, DELETE
Access-Control-Allow-Headers: Authorization, Content-Type
```

# To Test Preflight

```sh
# FE
go run ./cmd/examples/cors/preflight

# BE
go run ./cmd/api -cors-trusted-origins="http://localhost:9000"
```
