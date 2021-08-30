// file: api.go
// Example - demonstrates REST API server implementation tests.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func postMesage(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Name string `json:"name"`
		Age  int 	`json:"age"`
	}{Name:"testUser",Age:18}

	ok(w, data)
}

func main() {
	http.HandleFunc("/function/go-app", postMesage)
	http.ListenAndServe(":8001", nil)
}

// fail writes a json response with error msg and status header
func fail(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string `json:"error"`
	}{Error: msg}

	resp, _ := json.Marshal(data)
	w.WriteHeader(status)

	fmt.Fprintf(w, string(resp))
}

// ok writes data to response with 200 status
func ok(w http.ResponseWriter, data interface{}) {


	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fail(w, "oops something evil has happened", 500)
		return
	}

	if _,err = w.Write(resp);err!=nil {
		fmt.Printf("Error write: %s\n",err.Error())
	}
}
