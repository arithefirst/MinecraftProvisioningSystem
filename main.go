package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

// Use a special bool type to convert text to bools
type Bool bool

func (bit *Bool) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	*bit = Bool(txt == "1" || txt == "true")
	return nil
}

func main() {
	// Set the port for the server to run on
	var port uint16 = 8080

	// Endpoint for server.properties config generator
	http.HandleFunc("/api/v1/server-properties", serverPropertiesHttp)

	// Endpoints for getting jarfiles
	http.HandleFunc("/api/v1/jarfile/vanilla", getJarfileVanilla)
	http.HandleFunc("/api/v1/jarfile/fabric", getJarfileFabric)
	http.HandleFunc("/api/v1/jarfile/forge", getJarfileForge)

	// Server the files in frontend/ at /
	http.Handle("/", http.FileServer(http.Dir("frontend/")))

	fmt.Printf("Server started on port %v\n", port)
	var portString string = ":" + strconv.Itoa(int(port))
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		return
	}
}
