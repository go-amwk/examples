# State Management Example

This example demonstrates how to manage state in an AMWK web application. State management is the process of maintaining and sharing data across different handlers for a single request. AMWK provides a simple way to manage state using the `Set` and `Get` methods on the context.

In this example, we will create a simple web application that demonstrates how to set and get state in a handler. We will set a value in one handler and then retrieve it in another handler.

## Run and Test the Application

To run the application, use the following command:

```bash
go run main.go
```

Once the application is running, you can test it by sending a request to the server. You can use `curl` or any HTTP client to do this:

```bash
curl http://localhost:8000
# Expected response: Count: 1
curl http://localhost:8000
# Expected response: Count: 2
curl http://localhost:8000
# Expected response: Count: 3
```

You should see the response "Count: 1" for the first request, "Count: 2" for the second request, and the count will increment with each subsequent request. This demonstrates that the state is being maintained across handlers for each request.
