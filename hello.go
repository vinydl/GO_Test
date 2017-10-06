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

	http.HandleFunc("/", handler)
	http.HandleFunc("/save", savehandler)
	http.HandleFunc("/retrieve", retrievehandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprint("Saved %v\n", taskKey)
	fmt.Fprint(w, "Hello  world!")
}

func savehandler(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")
	ctx := context.Background()

	// Set your Google Cloud Platform project ID.
	projectID := "bipp-adhoc"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Fprint(w, input)
	// Sets the kind for the new entity.
	kind := "input"
	// Sets the name/ID for the new entity.
	name := input
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "Datastore input",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)

}
func retrievehandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	var task Task
	// Set your Google Cloud Platform project ID.
	projectID := "bipp-adhoc"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	taskKey := client.NewKey(ctx, "Entity", "stringID", 0, nil)

	q := datastore.NewQuery("input")
	taskKey, err = client.Get(ctx, q, &task)

	//kind := "input"
	// Sets the name/ID for the new entity.
	//name := input
	// Creates a Key instance.
	//taskKey := datastore..NewKey(tx, "kind", "stringID", 0, nil)

	fmt.Printf(task[taskKey].Key)

}
