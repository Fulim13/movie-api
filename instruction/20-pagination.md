# Page Formula

```
LIMIT = page_size
OFFSET = (page - 1) * page_size
```

# Test

```sh
curl "localhost:4000/v1/movies?page_size=2&page=2"
curl "localhost:4000/v1/movies?page_size=2&page=3"
```
