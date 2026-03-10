# Query Parameters Example

This example demonstrates how to handle query parameters in an AMWK web application. Query parameters are a common way to pass data to a server through the URL, and AMWK provides an easy way to access them in your handlers.

The AMWK framework allows you to access query parameters using the `Query`/`QueryValues`/`Queries` methods on the context. These methods return the value of the specified query parameter, all values for a parameter, or a map of all query parameters, respectively.

## Run and Test the Application

To run the application, use the following command:

```bash
go run main.go
```

Once the application is running, you can test it by sending a request with query parameters to the server. You can use `curl` or any HTTP client to do this:

```bash
curl "http://localhost:8000"
# Expected response: "Hello, World!"
curl "http://localhost:8000?name=Alice"
# Expected response: "Hello, Alice!"
curl "http://localhost:8000?name=Bob"
# Expected response: "Hello, Bob!"
```

You should see the response "Hello, World!" when no query parameter is provided, and "Hello, Alice!" or "Hello, Bob!" when the `name` query parameter is set to "Alice" or "Bob", respectively.
