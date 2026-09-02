# To Run the Code

```sh
printf '%s\n' (seq 8) | xargs -I % -P 8 curl -X PATCH -d '{"runtime": "97 mins"}' "localhost:4000/v1/movies/8"
 {
    "movie": {
        "id": 4,
        "title": "Breakfast Club",
        "year": 1985,
        "runtime": "97 mins",
        "genres": [
            "drama"
        ],
        "version": 4
    }
}
{
    "error": "unable to update the record due to an edit conflict, please try again"
}
{
    "error": "unable to update the record due to an edit conflict, please try again"
}
{
    "error": "unable to update the record due to an edit conflict, please try again"
}
{
    "error": "unable to update the record due to an edit conflict, please try again"
}
{
    "error": "unable to update the record due to an edit conflict, please try again"
}
{
    "error": "unable to update the record due to an edit conflict, please try again"
}
{
    "movie": {
        "id": 4,
        "title": "Breakfast Club",
        "year": 1985,
        "runtime": "97 mins",
        "genres": [
            "drama"
        ],
        "version": 5
    }
}
```
