# matrixhook

A simple and fast Matrix webhook bot made in Golang.

## Deployment

1. Move `.env.example` to `.env` and fill in the variables
2. Run with `go run .` or build it and execute the binary
3. (Optional) Add the `WEBSERVER_PRESHARED_KEY` environment variable in `.env` to require authentication.

## Usage

Send a POST request to the webserver (`/send`) with the following JSON formatted body:

```json
{
    "content": "this is my webhook message!"
}
```

The bot will then forward the content to the room ID of choice.

If `WEBSERVER_PRESHARED_KEY` is set, add the preshared key to the `authorization` header of the request.

## Attribution
- [Mautrix Go](https://github.com/mautrix/go)
- [go-gin](https://github.com/gin-gonic/gin)