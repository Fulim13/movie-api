| URL A                | URL B                  | Same origin? | Reason                                      |
| -------------------- | ---------------------- | ------------ | ------------------------------------------- |
| `https://foo.com/a`  | `http://foo.com/a`     | No           | Different scheme (`http` vs `https`)        |
| `http://foo.com/a`   | `http://www.foo.com/a` | No           | Different host (`foo.com` vs `www.foo.com`) |
| `http://foo.com/a`   | `http://foo.com:443/a` | No           | Different port (no port vs `443`)           |
| `http://foo.com/a`   | `http://foo.com/b`     | Yes          | Only the path is different                  |
| `http://foo.com/a`   | `http://foo.com/a?b=c` | Yes          | Only the query string is different          |
| `http://foo.com/a#b` | `http://foo.com/a#c`   | Yes          | Only the fragment is different              |

# CORS

For example, let’s say that you have a webpage at https://foo.com containing some front-end JavaScript code. If this JavaScript tries to make an HTTP request to https://bar.com/data.json (a different origin), then the request will be sent and processed by the bar.com server, but the user’s web browser will block the response so that the JavaScript code from https://foo.com cannot see it.

# To Test

```sh
go run ./cmd/api -cors-trusted-origins="http://localhost:9000 http://localhost:9001"

# CORS ERROR - because 9002 is not allowed in the trusted origins
go run ./cmd/examples/cors/simple --addr=":9002"

# NOT CORS ERRORS
go run ./cmd/examples/cors/simple --addr=":9000"
```
