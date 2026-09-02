# To Run the Code

```sh
$ curl -d '{"title": "Moana", "rating":"PG"}' localhost:4000/v1/movies
{
    "error": "body contains unknown key \"rating\""
}

$ curl -d '{"title": "Moana"}{"title": "Top Gun"}' localhost:4000/v1/movies
{
    "error": "body must only contain a single JSON value"
}

$ curl -d '{"title": "Moana"} :~()' localhost:4000/v1/movies
{
    "error": "body must only contain a single JSON value"
}

$ wget -O /tmp/largefile.json https://www.alexedwards.net/static/largefile.json
$ curl -d @/tmp/largefile.json localhost:4000/v1/movies
{
    "error": "body must not be larger than 1048576 bytes"
}
```
