package main

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	"dagger/my-module/internal/dagger"
)

type MyModule struct{}

// Test creates an Alpine container, writes three files, and emits custom spans for each.
func (m *MyModule) Test(ctx context.Context) *dagger.Directory {
	// Files to be created
	files := map[string]string{
		"test1.txt": "These are the results of test 1",
		"test2.txt": "These are the results of test 2",
		"test3.txt": "These are the results of test 3",
	}

	// Set up the Alpine container with the directory mounted
	container := dag.Container().
		From("alpine:latest").
		WithDirectory("/results", dag.Directory()).
		WithWorkdir("/results")
	for name, content := range files {
		// Create test files
		container = container.WithNewFile(name, content)
		// Emit custom spans for each test result file created
		log := "🧪 Test Result:\n" + "test file:\n" + name + "\n" + "contents:\n" + content 
		_, span := Tracer().Start(ctx, log)
		span.AddEvent("EVENT")
		span.SetStatus(codes.Ok, "STATUS")
		span.End()
    }
	return container.Directory("/results")
}

