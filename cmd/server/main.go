package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	w := os.Stdout

	if err := run(ctx, w); err != nil {
		w.Write([]byte("Error: " + err.Error()))
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer) error {
	if _, ok := os.LookupEnv("_EXPERIMENTAL_DAGGER_RUNNER_HOST"); !ok {
		return fmt.Errorf("`_EXPERIMENTAL_DAGGER_RUNNER_HOST` not set, e.g. tcp://localhost:1234")
	}

	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		out, err := client.Container().
			From("alpine:latest").
			WithExec([]string{"apk", "add", "curl"}).
			Stdout(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: " + err.Error()))
			return
		}

		w.Write([]byte("Container output:\n" + out))
	})

	fmt.Fprintf(w, "Server started at http://localhost:8080\n")

	return http.ListenAndServe(":8080", mux)
}
