# matrixhook

A simple and fast Matrix webhook bot made in Golang.

## Deployment

1. Move `.env.example` to `.env` and fill in the variables
2. Run with `go run .` or build it and execute the binary

## Usage

Send a POST request to the webserver with the following JSON formatted body:

```json
{
    "content": "this is my webhook message!"
}
```

The bot will then forward the content to the room ID of choice.

## Attribution
- [Mautrix Go](https://github.com/mautrix/go)
- [go-gin](https://github.com/gin-gonic/gin)