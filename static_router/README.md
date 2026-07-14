# Static Router Example

This example demonstrates how to use the `router` package in AMWK to define static routes and mount them into a web application. The router matches incoming requests by HTTP method and path, dispatching them to the appropriate handler — without middleware chains.

In this example, we create two routes:

- `/ping` — registered via `Any()`, responds to **all HTTP methods** (GET, POST, PUT, DELETE, etc.) with `"pong"`.
- `/secret` — registered via `GET()`, responds only to **GET** requests with `"Secret data"`.

## Run and Test the Application

To run the application, use the following command:

```bash
go run app.go
```

Once the application is running, you can test the routes using `curl` or any HTTP client:

```bash
# /ping accepts any HTTP method
curl http://localhost:8000/ping -v
# < HTTP/1.1 200 OK
# pong

curl -X POST http://localhost:8000/ping -v
# < HTTP/1.1 200 OK
# pong

# /secret only accepts GET
curl http://localhost:8000/secret -v
# < HTTP/1.1 200 OK
# Secret data

curl -X POST http://localhost:8000/secret -v
# < HTTP/1.1 404 Not Found
# 404 Not Found

# /other does not exist
curl http://localhost:8000/other -v
# < HTTP/1.1 404 Not Found
# 404 Not Found
```
