# glt

`glt` is a simple HTTP server written in Go that supports for now three routes: `/`, `/name`, and `/health`.

## Installation

To install `glt`, you must have Go installed on your system. Once you have Go installed, you can download and install `glt` using the following command:

```
go get github.com/0nizzr/glt
```

## Usage

To start the `glt` server, navigate to the `glt` directory and run the following command:

```
go run main.go
```

This will start the server on port 8080.

### Routes

`glt` supports the following routes:

- `/`: This is the default route and will respond with the message "Hello, World!".

- `/name`: This route expects a `name` parameter in the query string and will respond with a personalized greeting that includes the provided name.

- `/health`: This route will respond with an HTTP 200 status code and the message "OK" to indicate that the server is healthy.

- Any other route: Visiting a route that is not `/`, `/name` or `/health` will return an HTTP 404 status code and the message "404 page not found".
