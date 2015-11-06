//START IMPORT-OMIT
package main

import (
	"bytes"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

//END IMPORT-OMIT

//START TC-OMIT
// This is a basic struct to hold basic page data variables
type PageData struct {
	Title string
	Body  string
}

// Define a basic HTML template
const html = `<HTML>
	<head><title>{{.Title}}</title></head>
	<body>
	<h1>{{.Title}}</h1>
	{{.Body}}
	</body>
	</HTML>`

//END TC-OMIT

//START FUNCTIONS-M-OMIT
func main() {
	// We need to create a router
	rt := mux.NewRouter().StrictSlash(true)

	// Add the "index" or root path
	rt.HandleFunc("/", Index)

	// Fire up the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", rt))
}

//END FUNCTIONS-M-OMIT

//START FUNCTIONS-I-OMIT
// Index is the "index" handler
func Index(w http.ResponseWriter, r *http.Request) {
	// Fill out page data for index
	pd := PageData{
		Title: "Index Page",
		Body:  "This is the body of the page.",
	}

	// Render a template with our page data
	tmpl, err := render(pd)

	// If we got an error, write it out and exit
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// All went well, so write out the template
	w.Write([]byte(tmpl))
}

//END FUNCTIONS-I-OMIT

//START FUNCTIONS-R-OMIT
func render(pd PageData) (string, error) {
	// Parse the template
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		return "", err
	}

	// We need somewhere to write the executed template to
	var out bytes.Buffer

	// Render the template with the data we passed in
	if err := tmpl.Execute(&out, pd); err != nil {
		// If we couldn't render, return a error
		return "", err
	}

	// Return the template
	return out.String(), nil
}

//END FUNCTIONS-R-OMIT
