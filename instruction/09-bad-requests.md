# Before Running the Code

```sh
# Send some XML as the request body
$ curl -d '<?xml version="1.0" encoding="UTF-8"?><note><to>Alice</to></note>' localhost:4000/v1/movies
{
    "error": "invalid character '\u003c' looking for beginning of value"
}

# Send some malformed JSON (notice the trailing comma)
$ curl -d '{"title": "Moana", }' localhost:4000/v1/movies
{
    "error": "invalid character '}' looking for beginning of object key string"
}

# Send a JSON array instead of an object
$ curl -d '["foo", "bar"]' localhost:4000/v1/movies
{
    "error": "json: cannot unmarshal array into Go value of type struct { Title string
    \"json:\\\"title\\\"\"; Year int32 \"json:\\\"year\\\"\"; Runtime int32 \"json:\\
    \"runtime\\\"\"; Genres []string \"json:\\\"genres\\\"\" }"
}

# Send a numeric 'title' value (instead of string)
$ curl -d '{"title": 123}' localhost:4000/v1/movies
{
    "error": "json: cannot unmarshal number into Go struct field .title of type string"
}

# Send an empty request body
$ curl -X POST localhost:4000/v1/movies
{
    "error": "EOF"
}
```

# After Running the Code

```sh
# Send some XML as the request body
$ curl -d '<?xml version="1.0" encoding="UTF-8"?><note><to>Alex</to></note>' localhost:4000/v1/movies
{
    "error": "body contains badly-formed JSON (at character 1)"
}

# Send some malformed JSON (notice the trailing comma)
$ curl -d '{"title": "Moana", }' localhost:4000/v1/movies
{
    "error": "body contains badly-formed JSON (at character 20)"
}

# Send a JSON array instead of an object
$ curl -d '["foo", "bar"]' localhost:4000/v1/movies
{
    "error": "body contains incorrect JSON type (at character 1)"
}

# Send a numeric 'title' value (instead of string)
$ curl -d '{"title": 123}' localhost:4000/v1/movies
{
    "error": "body contains incorrect JSON type for \"title\""
}

# Send an empty request body
$ curl -X POST localhost:4000/v1/movies
{
    "error": "body must not be empty"
}
```
