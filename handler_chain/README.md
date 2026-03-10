# Handler Chain Example

This example demonstrates how to use the handler chain in AMWK to create a simple web application. The handler chain allows you to define a sequence of handlers that will be executed for each incoming request. This can be useful for tasks such as logging, authentication, and response formatting.

In this example, we will create a simple web application that responds with "Hello, World!" to any incoming request. We will define a handler chain that includes three middleware handlers, a handler that generates the response and abort the chain, and a handler that should never be reached. The middleware handlers will log the start and end of the request processing, as well as a message in the middle of the chain.

## Run and Test the Application

To run the application, use the following command:

```bash
go run main.go
```

Once the application is running, you can test it by sending a request to the server. You can use `curl` or any HTTP client to do this:

```bash
curl http://localhost:8000
```

You should see the response "Hello, World!" in your terminal. And the server logs will show the handlers' log messages, demonstrating the execution of the handler chain.

```
2026/02/01 12:00:00 middleware1 start
2026/02/01 12:00:00 middleware2 start
2026/02/01 12:00:00 middleware3
2026/02/01 12:00:00 handler
2026/02/01 12:00:00 middleware2 end
2026/02/01 12:00:00 middleware1 end
```
