package main

import (
	"fmt"
	"log"
	"net/http"
	// Imports the Google Cloud Datastore client package.
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

type Task struct {
	Description string
}

func main() {

	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "bipp-adhoc"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Sets the kind for the new entity.
	kind := "input"
	// Sets the name/ID for the new entity.
	name := "abc"
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "Saving input",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint("Saved %v\n", taskKey)
	fmt.Fprint(w, "Hello  world!")
}
