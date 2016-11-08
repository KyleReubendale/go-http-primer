package main

import (
	"fmt"
	"net/http"
	"html"
)

// This the template for our root page
const template = `<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="public/css/main.css">
		<script src="public/js/main.js"></script>
	</head>
	<body>
		<h2>Thanks for visiting my path %q</h2>
	</body>
</html>
`

// main function runs the app
func main() {
	port := ":3000"

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", root)
	fmt.Printf("Starting server on port %s\n", port)
	panic(http.ListenAndServe(port, nil))

}

// root is the handler function, it tells me what to do when a user visits root
func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, template, html.EscapeString(r.URL.Path))
	return
}
