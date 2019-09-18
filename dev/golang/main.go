package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const namespaceFile = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

var namespace = ""

func main() {
	fmt.Println("Starting hello-world server...")
	bytes, err := ioutil.ReadFile(namespaceFile)
	if err != nil {
		panic(fmt.Sprintf("Looks like you are not running inside a Kubernetes Pod: %s", err))
	}

	namespace = string(bytes)
	http.HandleFunc("/", helloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	msg := getMessage(namespace)
	fmt.Fprint(w, msg)
}

func getMessage(namespace string) string {
	return fmt.Sprintf("Hello World from the cluster namespace '%s'", namespace)
}
