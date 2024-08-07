package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Use a special bool type to convert text to bools
type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}

func serverPropertiesHttp(w http.ResponseWriter, r *http.Request) {

	// Frontend sends the server.properties config VIA headers in a get request
	// Func converts it into a server.properties file and returns it in the get request

	var jsonProperties serverProperties

	jsonString := r.Header.Get("jsonString")
	fmt.Printf("Recived jsonString: %v\n", jsonString)

	err := json.Unmarshal([]byte(jsonString), &jsonProperties)
	if err != nil {
		fmt.Fprintf(w, "Error parsing JSON: %v", err)
		return
	}

	fmt.Fprintf(w, "%v", jsonProperties)
}

func main() {
	http.HandleFunc("/server-properties", serverPropertiesHttp)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
