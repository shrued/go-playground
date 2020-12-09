package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>How you doin'?</h1>
<p>Please watch Euphoria.</p>`)
	fmt.Fprintf(w, "This is going to be...wait for it...%s!","legendary")
} 

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Shruti here :3")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":9000", nil)
}