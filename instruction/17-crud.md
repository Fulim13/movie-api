# Creating a Record

```sh
set BODY '{"title":"Moana","year":2016,"runtime":"107 mins", "genres":["animation","adventure"]}'
curl -i -d "$BODY" localhost:4000/v1/movies
set BODY '{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["action","adventure"]}'
curl -d "$BODY" localhost:4000/v1/movies
set BODY '{"title":"Deadpool","year":2016, "runtime":"108 mins","genres":["action","comedy"]}'
curl -d "$BODY" localhost:4000/v1/movies
set BODY '{"title":"The Breakfast Club","year":1986, "runtime":"96 mins","genres":["drama"]}'
curl -d "$BODY" localhost:4000/v1/movies
```

# Getting a Record

```sh
curl -i localhost:4000/v1/movies/5
curl -i localhost:4000/v1/movies/42
```

# Update a Record

```sh
set BODY '{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
curl -X PUT -d "$BODY" localhost:4000/v1/movies/6
```

# Delete a Record

```sh
curl -X DELETE localhost:4000/v1/movies/7
curl -X DELETE localhost:4000/v1/movies/7
```
