package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Use a special bool type to convert text to bools
type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}

func serverPropertiesHttp(w http.ResponseWriter, r *http.Request) {

	// Frontend sends the server.properties config VIA querystring in a get request
	// Func converts it into a server.properties file and returns it in the get request

	var jsonProperties serverProperties

	encodedJsonString := r.URL.Query().Get("jsonString")
	jsonString, err := url.QueryUnescape(encodedJsonString)
	if err != nil {
		fmt.Fprintf(w, "Error decoding url: %v", err)
	}
	fmt.Printf("Recived jsonString: %v\n", jsonString)

	err = json.Unmarshal([]byte(jsonString), &jsonProperties)
	if err != nil {
		fmt.Fprintf(w, "Error parsing JSON: %v", err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%v", generateServerProperties(jsonProperties))
}

func main() {
	// Set the port for the server to run on
	var port uint16 = 8080

	http.HandleFunc("/api/v1/server-properties", serverPropertiesHttp)
	http.Handle("/", http.FileServer(http.Dir("frontend/")))

	fmt.Printf("Server started on port %v\n", port)
	var portString string = ":" + strconv.Itoa(int(port))
	http.ListenAndServe(portString, nil)
}
