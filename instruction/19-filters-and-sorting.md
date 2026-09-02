# Create Index

```sh
migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
migrate -path ./migrations -database postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable up
```

# To Test the Validation of Filters and Sort

```sh
curl "localhost:4000/v1/movies?page=-1&page_size=-1&sort=foo"

{
        "error": {
                "page": "must be greater than zero",
                "page_size": "must be greater than zero",
                "sort": "invalid sort value"
        }
}
```

# To Filter and Sort

```sh
curl "localhost:4000/v1/movies?title=black+panther"
curl "localhost:4000/v1/movies?title=the+club"
curl "localhost:4000/v1/movies?genres=adventure"
curl "localhost:4000/v1/movies?title=moana&genres=animation,adventure"
curl "localhost:4000/v1/movies?sort=-title"
curl "localhost:4000/v1/movies?sort=-runtime"
```
