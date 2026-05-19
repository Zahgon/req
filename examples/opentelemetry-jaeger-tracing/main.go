package main

import (
	"context"
	"fmt"
	"log"
	"opentelemetry-jaeger-tracing/github"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

const serviceName = "github-query"

var githubClient *github.Client

func traceProvider() (*trace.TracerProvider, error) {
	_ = "STUB: not implemented"
	// Create the Jaeger exporter
	return nil, nil
}

// Record information about this application in a Resource.

// Create the TraceProvider.

// Always be sure to batch in production.

// Record information about this application in a Resource.

// QueryUser queries information for specified GitHub user, and display a
// brief introduction which includes name, blog, and the most popular repo.
func QueryUser(username string) error { _ = "STUB: not implemented"; return nil }

func findMostPopularRepo(ctx context.Context, username string) (repo *github.Repo, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func main() {
	tp, err := traceProvider()
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)

	githubClient = github.NewClient()
	if os.Getenv("DEBUG") == "on" {
		githubClient.SetDebug(true)
	}
	if token := os.Getenv("GITHUB_TOKEN"); token != "" {
		githubClient.LoginWithToken(token)
	}
	githubClient.SetTracer(otel.Tracer("github"))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigs
		fmt.Printf("Caught %s, shutting down\n", sig)
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		var name string
		fmt.Printf("Please give a github username: ")
		_, err := fmt.Fscanf(os.Stdin, "%s\n", &name)
		if err != nil {
			panic(err)
		}
		err = QueryUser(name)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
