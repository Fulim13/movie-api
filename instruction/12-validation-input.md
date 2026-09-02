# To Run the Code

```sh
$ set BODY '{"title":"","year":1000,"runtime":"-123 mins","genres":["sci-fi","sci-fi"]}'
$ curl -i -d "$BODY" localhost:4000/v1/movies
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Wed, 07 Apr 2021 10:51:00 GMT
Content-Length: 180

{
    "error": {
        "genres": "must not contain duplicate values",
        "runtime": "must be a positive integer",
        "title": "must be provided",
        "year": "must be greater than 1888"
    }
}
```
