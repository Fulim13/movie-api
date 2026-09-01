# Install httprouter

```sh
go go get github.com/julienschmidt/httprouter@v1
```

# To Run the Code

```sh
curl -i localhost:4000/v1/healthcheck
curl -i -X POST localhost:4000/v1/movies
curl -i localhost:4000/v1/movies/0
curl -i localhost:4000/v1/movies/a
curl -i localhost:4000/v1/movies/1
```
