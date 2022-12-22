package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Aarya-Patel/cyoa/internal/parser"
	"github.com/Aarya-Patel/cyoa/internal/webpage"
)

/* ------ GLOBALS ------ */
var (
	filename = flag.String("filename", "gopher.json", "The filename containing the CYOA story")
)

func main() {
	flag.Parse()

	story, err := parser.ParseStory(*filename)
	if err != nil {
		panic(err)
	}

	mux := webpage.CreateWebPagesForStory(story)
	fmt.Println("Server listening on PORT 8080...")
	http.ListenAndServe(":8080", mux)
}
