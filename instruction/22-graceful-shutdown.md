| Signal  | Description                            | Keyboard shortcut | Catchable |
| ------- | -------------------------------------- | ----------------- | --------- |
| SIGINT  | Interrupt from keyboard                | Ctrl+C            | Yes       |
| SIGQUIT | Quit from keyboard                     | Ctrl+\            | Yes       |
| SIGKILL | Kill process (terminate immediately)   | -                 | No        |
| SIGTERM | Terminate process in an orderly manner | -                 | Yes       |

```sh
go run ./cmd/api
pkill -SIGKILL api
pkill -SIGTERM api
Ctrl+\
```

# Test Graceful Shutdown

```sh
go run ./cmd/api
curl localhost:4000/v1/healthcheck & pkill -SIGTERM api
```
