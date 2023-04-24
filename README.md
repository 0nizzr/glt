# glt

`glt` is a simple HTTP server written in Go that supports for now three routes: `/`, `/name`, and `/health`.

## Installation

To install `glt`, you must have Go installed on your system. Once you have Go installed, you can clone the repository from GitHub using the following command:

```
git clone https://github.com/0nizzr/glt.git
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

Sure, here's a possible rewrite of the README.md file:

# GLT

GLT is a simple Go web server that displays a "Hello, world!" message on the root path `/` and a "404 Not Found" message for any other path.

## Installation

To install GLT, you can clone the repository from GitHub using the following command:

```
git clone https://github.com/[your-github-username]/glt.git
```

## Usage

To run GLT, navigate to the project directory and execute the following command:

```
go run main.go
```

This will start the web server, which listens on port 8080 by default. You can access the server by visiting `http://localhost:8080` in your web browser.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find a bug or have a feature request.
