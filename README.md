# Example Dagger Custom Application

This is an example Go application on how you can embed/use Dagger.io inside of an application like an API with a remote Dagger Enginer/Runner.

See these docs for more information: https://docs.dagger.io/configuration/custom-runner/

## Getting Started

1. Clone the repository:

```bash
git clone git@github.com:jasonmccallister/dagger-embedded-go-api.git

cd dagger-embedded-go-api
```

2. Run the Dagger Engine in Docker Compose:

```bash
docker compose up -d
```

3. Set the `_EXPERIMENTAL_DAGGER_RUNNER_HOST`

```bash
export _EXPERIMENTAL_DAGGER_RUNNER_HOST=tcp://localhost:1234
```

4. Run the application:

```bash
go run cmd/server/main.go
```

5. Send a request to the API (using curl or browser):

```bash
http://localhost:8080
```

You will see the respone from the Dagger Engine where is made a new container, installed a package, and sent the output in the HTTP response.
