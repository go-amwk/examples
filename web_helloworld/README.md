# Web Server with HTTP Adapter (web package) Example

This example demonstrates how to create a simple web server using the HTTP adapter provided by the `web` package in AMWK. The application will respond with "Hello, World!" to any incoming HTTP request.

## Run and Test the Application

To run the application, use the following command:

```bash
go run main.go
```

Once the application is running, you can test it by sending a request to the server. You can use `curl` or any HTTP client to do this:

```bash
curl http://localhost:8000
```

You should see the response "Hello, World!" in your terminal.
