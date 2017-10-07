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

// Sets the kind for the new entity.
//var k1 = datastore.NameKey("Entity", "stringID", nil)

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
	k := datastore.NameKey("Entity", "stringID", nil)
	//	k := datastore.NewIncompleteKey(ctx, "Entity", nil)
	e := Entity{input}
	if _, err = client.Put(ctx, k, &e); err != nil {

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
	k := datastore.NewQuery("__key__")
	var e []*Entity
	//	e := new(Entity)
	keys, err := client.GetAll(ctx, k, &e)
	//fmt.Printf()

	for i, key := range keys {
		fmt.Println(key)
		fmt.Println(e[i])
	}
	//	for i := range e {
	//	fmt.Printf("Saved %v: \n", e[i].Value)
	//fmt.Fprint(w, e[i].Value)
	//}
}
