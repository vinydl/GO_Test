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

type Entity struct {
	Value string
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
	k := datastore.NameKey("Entity", "stringID", nil)
	e := Entity{input}
	if _, err := client.Put(ctx, k, &e); err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	fmt.Printf("Saved %q\n", e.Value)

}
func retrievehandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	// Set your Google Cloud Platform project ID.
	projectID := "bipp-adhoc"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	k := datastore.NameKey("Entity", "stringID", nil)
	e := new(Entity)
	if err := client.Get(ctx, k, e); err != nil {
		// Handle error.
	}
	//fmt.Printf()
	fmt.Printf("Saved %v: \n", e.Value)
	fmt.Fprint(w, e.Value)
}
