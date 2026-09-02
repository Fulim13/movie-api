# To Run the Code

```sh
$ curl -d '{"title": "Moana", "runtime": "107 mins"}' localhost:4000/v1/movies
{Title:Moana Year:0 Runtime:107 Genres:[]}

$ curl -d '{"title": "Moana", "runtime": 107}' localhost:4000/v1/movies
{
        "error": "invalid runtime format"
}

$ curl -d '{"title": "Moana", "runtime": "107 minutes"}' localhost:4000/v1/movies
{
        "error": "invalid runtime format"
}
```
